// ESPN live-draft in-page agent (experimental).
//
// Proven in a full 17-round practice draft on 2026-08-30. The MCP
// screenshot->decide->click loop is far too slow for a pick clock (each round
// trip is 2-10s; ESPN auto teams pick every ~4s and the human clock is 30s),
// so this script runs *inside* the draft page and reacts in <1s. Claude's job
// shifts from clicking to strategy: it polls the [PICK]/[AGENT] console logs
// every 20-60s, recomputes the VOR board with ffdraft, and updates the ranked
// target list. ESPN's native pick queue + autopick remain the failsafe of last
// resort.
//
// Usage (from Claude via javascript_tool, or DevTools console):
//   1. Open the ESPN draft room, ideally before the draft starts.
//   2. Paste/inject this whole file, then configure the hard rails:
//        __draftAgent.myTeamName = 'Exact Team Name From Picks Panel';
//        __draftAgent.positionCaps = {QB: 2, TE: 2, K: 1, 'D/ST': 1};
//      Caps are enforced at every click point (targets, visible-target
//      shortcut, fallback, queue sync) and count from the live pick feed, so
//      a stale list or fallback can never overfill a position. Caps are
//      inert until myTeamName is set.
//   3. Update strategy any time:  __draftAgent.setTargets(['Name', ...], ['K','D/ST'])
//      (ranked, best first; exact ESPN display names — "Broncos D/ST", not
//      "Denver Broncos". setTargets also re-mirrors the ESPN queue; the
//      second argument optionally replaces avoidPositions.)
//   4. Read state:               __draftAgent.state()
//      Sync picks from state().picksSeen (stable buffer) — console clear
//      proved unreliable. [AGENT]/[PICK] console tags remain for humans.
//   5. Stop:                     __draftAgent.stop()
//
// Hard-won details encoded below:
//   - The player list is a virtualized FixedDataTable: only ~15 rows exist in
//     the DOM, and during re-renders a row's name anchor and its DRAFT button
//     can momentarily belong to different logical rows (one practice pick
//     landed on the wrong D/ST that way). So we never click a button from the
//     unfiltered list when targeting a specific player - we type the name into
//     the search box first, wait for the filter, and verify the surviving
//     row's anchor text before clicking.
//   - ESPN silently re-enables autopick whenever your clock expires; the loop
//     turns it back off every tick.
//   - The lm-api mDraftDetail endpoint does NOT reflect practice/mock drafts
//     (players stay pid:-1), so DOM is the only reliable state source there.
//   - Hook WebSocket before the draft starts and every live message is
//     captured in __draftAgent.wsLog (nice-to-have; DOM polling suffices).

(() => {
  if (window.__draftAgent) window.__draftAgent.stop();

  // A re-injection (bug fix mid-draft, page reload) inherits the previous
  // instance's plan so there is no unprotected window with empty targets.
  const prev = window.__draftAgentCfg || {};
  const A = (window.__draftAgent = {
    targets: prev.targets || [],   // ranked player display names, best first
    avoidPositions: prev.avoid || [], // e.g. ['K','D/ST','QB'] - skipped by the fallback pick
    // Hard rail: never draft past these per-position counts, no matter what a
    // stale list, fallback, or queue entry says. A 16-team rehearsal ended
    // with 3 QBs via three different paths; judgment composes the list, caps
    // make it impossible to blow through. Requires myTeamName to count picks.
    positionCaps: { QB: 2, TE: 2, K: 1, 'D/ST': 1 },
    myTeamName: '',       // exact fantasy team name as shown in the Picks panel
    log: [],
    wsLog: [],
    picking: false,       // guard so the async pick sequence runs once per turn
    intervals: [],
  });
  const say = (m) => { A.log.push({ t: new Date().toISOString(), m }); console.log('[AGENT] ' + m); };
  const sleep = (ms) => new Promise((r) => setTimeout(r, ms));

  // --- WebSocket capture (only sockets opened after injection) -------------
  if (!window.__wsHooked) {
    window.__wsHooked = true;
    const Orig = window.WebSocket;
    window.WebSocket = function (url, protocols) {
      const ws = protocols ? new Orig(url, protocols) : new Orig(url);
      A.wsLog.push({ t: Date.now(), type: 'open', url });
      ws.addEventListener('message', (e) => {
        A.wsLog.push({ t: Date.now(), type: 'msg', data: String(e.data).slice(0, 4000) });
        if (A.wsLog.length > 1000) A.wsLog.splice(0, 200);
      });
      return ws;
    };
    window.WebSocket.prototype = Orig.prototype;
    Object.assign(window.WebSocket, Orig);
  }

  // --- DOM helpers ----------------------------------------------------------
  const buttons = () => [...document.querySelectorAll('button')];
  const draftButtons = () =>
    [...document.querySelectorAll('button.action-btn')].filter(
      (b) => b.textContent.trim().toUpperCase() === 'DRAFT'
    );
  const rowName = (btn) =>
    btn.closest('.fixedDataTableCellGroupLayout_cellGroup')?.querySelector('a')?.textContent.trim() || '';
  // Position comes from the dedicated playerinfo element. The old
  // \b-anchored regex over the row's textContent silently returned '' in the
  // 2026-08-30 practice room: FixedDataTable concatenates cell text with no
  // whitespace ("Ja'Marr ChaseQCINWRFourth and Long Shot..."), so \bWR\b
  // never matches, every atCap() check no-ops, and the fallback ladder
  // degrades to btns[0] — that room ended with FOUR QBs under a QB:1 cap.
  const rowPos = (btn) => {
    const g = btn.closest('.fixedDataTableCellGroupLayout_cellGroup');
    const el = g?.querySelector('.playerinfo__playerpos, [class*="playerpos"]');
    // Two-way players render as "WR, CB"/"WRCB" — keep the leading token.
    if (el) return (el.textContent.trim().match(/^(QB|RB|WR|TE|K|D\/ST)/) || [])[1] || '';
    const txt = g?.textContent || '';
    return (txt.match(/\b(QB|RB|WR|TE|K|D\/ST)\b/) || [])[1] || '';
  };

  // Strict single-pick entry: "Name / TEAM POS R<n>, P<n> - Team Name".
  // The Picks panel also renders container elements whose text concatenates
  // several picks; those double-counted positions in live fire (RB:5 after
  // two picks) until entries were parsed strictly and deduped by pick slot.
  const PICK_RE = /^([^\/]+) \/ (\S+) (QB|RB|WR|TE|K|D\/ST)R(\d+), P(\d+) - (.+)$/;

  // My per-position counts, derived from the pick feed.
  const myPosCounts = () => {
    const counts = {}, slots = new Set();
    if (!A.myTeamName) return counts;
    for (const t of window.__seenPicksSet || []) {
      const m = t.match(PICK_RE);
      if (!m || m[6] !== A.myTeamName) continue;
      const slot = m[4] + '-' + m[5];
      if (slots.has(slot)) continue;
      slots.add(slot);
      counts[m[3]] = (counts[m[3]] || 0) + 1;
    }
    return counts;
  };
  const atCap = (pos) => {
    const cap = A.positionCaps[pos];
    return cap !== undefined && (myPosCounts()[pos] || 0) >= cap;
  };

  // Roster-legality guard (the caps' complement). Caps only bound maximums;
  // late in a draft the binding constraint flips to minimums: with N picks
  // left and N unfilled required slots, ONLY those positions are legal, and
  // ESPN's UI enforces it. Run 2 of the 2026-08-30 mocks walked WR names at
  // pick 159 that ESPN would never allow (only TE and K slots remained).
  // Defaults match fnf; override per league via __draftAgent.requiredSlots /
  // totalRounds.
  A.requiredSlots = prev.requiredSlots || { QB: 1, RB: 2, WR: 2, TE: 1, K: 1, 'D/ST': 1 };
  A.totalRounds = prev.totalRounds || 17;
  const mustFillPositions = () => {
    if (!A.myTeamName) return null; // counts unknowable - guard inert
    const c = myPosCounts();
    const mineSoFar = Object.values(c).reduce((a, b) => a + b, 0);
    const remaining = A.totalRounds - mineSoFar;
    const need = [];
    let deficit = 0;
    for (const pos of Object.keys(A.requiredSlots)) {
      const d = Math.max(0, A.requiredSlots[pos] - (c[pos] || 0));
      if (d > 0) { need.push(pos); deficit += d; }
    }
    return deficit >= remaining ? need : null; // null = no restriction yet
  };
  // Single legality test used at every click point: under cap AND, once the
  // must-fill window closes in, a position we still owe a slot. Unknown
  // positions ('') stay allowed - better a pick than a lost clock.
  const posAllowed = (pos) => {
    if (!pos) return true;
    if (atCap(pos)) return false;
    const need = mustFillPositions();
    return !need || need.includes(pos);
  };

  const searchBox = () =>
    document.querySelector('input[placeholder*="Player" i]') ||
    document.querySelector('input[type="search"]');

  // React controlled input: must use the native setter + input event.
  const setSearch = (value) => {
    const box = searchBox();
    if (!box) return false;
    const set = Object.getOwnPropertyDescriptor(window.HTMLInputElement.prototype, 'value').set;
    set.call(box, value);
    box.dispatchEvent(new Event('input', { bubbles: true }));
    return true;
  };

  // --- Picking --------------------------------------------------------------
  const tryDraftExact = async (name) => {
    if (!setSearch(name)) return false;
    await sleep(350); // let the filter re-render
    // If something else (a concurrent syncQueue) retyped the box while we
    // slept, the rows are filtered to the wrong name. This race walked past
    // NINE available targets (incl. the live #1) in one pick of the second
    // fnf mock. Re-assert our search once before trusting the rows.
    const box = searchBox();
    if (box && box.value !== name) {
      setSearch(name);
      await sleep(350);
    }
    const hit = draftButtons().find((b) => rowName(b) === name);
    if (hit && !posAllowed(rowPos(hit))) {
      say('SKIPPED ' + name + ': ' + rowPos(hit) + ' capped or not a must-fill slot');
      setSearch('');
      return false;
    }
    if (hit) {
      hit.click();
      say('DRAFT clicked: ' + name + ' (target)');
    }
    setSearch('');
    return !!hit;
  };

  // Best visible target: the button whose row name has the lowest targets
  // index. Instant (no search round trip), but limited to the ~15 rendered
  // rows of the virtualized list.
  const bestVisibleTarget = () => {
    let best = null, bestIx = Infinity;
    for (const b of draftButtons()) {
      if (!posAllowed(rowPos(b))) continue;
      const ix = A.targets.indexOf(rowName(b));
      if (ix >= 0 && ix < bestIx) { bestIx = ix; best = b; }
    }
    return { best, bestIx };
  };

  // The 2026-08-30 rehearsal lost one clock race walking dead names at 500ms
  // each while autodrafting bots picked every 1-2s. So: click a visible
  // top-3 target instantly; otherwise search-walk the live list under a time
  // budget, and past the budget take the best visible target or fallback.
  const pickSequence = async () => {
    A.picking = true;
    // 15s of a 30s clock: live fire showed 8s expiring inside a long dead-name
    // walk before the list reached the positions the pick actually needed.
    const deadline = Date.now() + 15000;
    try {
      // Canary: DRAFT buttons whose position we cannot read mean every cap
      // and the must-fill guard are silently inert (the first fnf mock drafted
      // 4 QBs exactly this way). Scream once so the operator sees it in the
      // [AGENT] log instead of discovering it from the roster.
      const btnsNow = draftButtons();
      if (btnsNow.length && btnsNow.every((b) => !rowPos(b)) && !A._warnedBlindPos) {
        A._warnedBlindPos = true;
        say('WARNING: rowPos blind on ALL visible rows - position caps and must-fill guard are INERT. DOM likely changed; check selfTest().');
      }
      let { best, bestIx } = bestVisibleTarget();
      if (best && bestIx < 3) {
        const name = rowName(best);
        best.click();
        say('DRAFT clicked: ' + name + ' (target, visible)');
        return;
      }
      // Walk the LIVE targets array (restart if replaced mid-walk) until the
      // budget runs out or we reach a target we already know is visible.
      let list = A.targets, i = 0;
      while (i < list.length && i < (bestIx === Infinity ? list.length : bestIx) && Date.now() < deadline) {
        if (list !== A.targets) { list = A.targets; i = 0; say('targets replaced mid-pick, restarting walk'); continue; }
        const name = list[i++];
        if (await tryDraftExact(name)) return;
        if (!draftButtons().length) return; // pick already went through / clock lost
      }
      ({ best, bestIx } = bestVisibleTarget());
      if (best) {
        const name = rowName(best);
        best.click();
        say('DRAFT clicked: ' + name + ' (target, visible)');
        return;
      }
      await sleep(400);
      const btns = draftButtons();
      if (btns.length) {
        // Fallback ladder: best visible player not avoided or capped; then
        // merely not capped (caps outrank avoid); only then anything at all —
        // a live-fire pick blew through the RB cap because the last resort
        // ignored it while every visible row happened to be capped/avoided.
        const pick =
          btns.find((b) => !A.avoidPositions.includes(rowPos(b)) && posAllowed(rowPos(b))) ||
          btns.find((b) => posAllowed(rowPos(b))) ||
          btns[0];
        const name = rowName(pick);
        pick.click();
        say('DRAFT clicked: ' + name + ' (best-available fallback)');
      }
    } catch (e) {
      say('ERR pick: ' + e.message);
    } finally {
      await sleep(2000); // let ESPN register the pick before re-arming
      A.picking = false;
    }
  };

  A.intervals.push(
    setInterval(() => {
      try {
        // ESPN flips autopick back on when your clock expires - keep it off.
        const dis = buttons().find((b) => /disable autopick/i.test(b.textContent));
        if (dis) { dis.click(); say('disabled autopick'); }
        if (!A.picking && draftButtons().length) pickSequence();
      } catch (e) {
        say('ERR loop: ' + e.message);
      }
    }, 400)
  );

  // --- State tracking -------------------------------------------------------
  // The seen-set lives on window so a re-injection doesn't replay old picks.
  // Read picks via state().picksSeen (a stable buffer) rather than console
  // logs: read_console_messages' clear option proved unreliable in rehearsal.
  const seen = (window.__seenPicksSet = window.__seenPicksSet || new Set());
  A.intervals.push(
    setInterval(() => {
      try {
        // Completed picks render in the right-hand Picks panel as
        // "Player Name / TEAM POS  R<n>, P<n> - Team Name".
        [...document.querySelectorAll('div,li')]
          .filter((e) => e.children.length <= 3 && /R\d+, P\d+ - /.test(e.textContent) && e.textContent.length < 120)
          .forEach((el) => {
            const txt = el.textContent.replace(/\s+/g, ' ').trim();
            if (!PICK_RE.test(txt)) return; // strict entries only - no panel concats
            if (!seen.has(txt)) {
              seen.add(txt);
              console.log('[PICK] ' + txt);
              // Self-prune: a drafted player can never be a target again, and
              // dead names slow the pick walk (that lost a clock race once).
              const nm = txt.split(' / ')[0].trim();
              const ix = A.targets.indexOf(nm);
              if (ix >= 0) { A.targets.splice(ix, 1); say('pruned drafted target: ' + nm); }
            }
          });
      } catch (e) {}
    }, 1000)
  );

  A.state = () => ({
    round: document.querySelector('[class*=clock]')?.textContent || '',
    onTheClock: [...document.querySelectorAll('div')].find((e) => /^On the Clock/i.test(e.textContent.trim()) && e.children.length === 0)?.textContent || '',
    myTurn: draftButtons().length > 0,
    roster: [...document.querySelectorAll('aside a, [class*=roster] a')].map((a) => a.textContent.trim()).filter(Boolean),
    myPosCounts: myPosCounts(),
    mustFill: mustFillPositions(),
    picksSeen: [...seen].filter((t) => PICK_RE.test(t)),
    targets: A.targets,
    recentLog: A.log.slice(-10),
  });

  // Preferred way to push a new plan: assigns targets and re-mirrors the ESPN
  // queue in the background. In rehearsal, ESPN autopick swallowed two turns
  // whose targets were only in A.targets — an up-to-date queue would have
  // drafted our list anyway.
  A.setTargets = (list, avoid) => {
    A.targets = list.slice();
    if (avoid) A.avoidPositions = avoid;
    window.__draftAgentCfg = { targets: A.targets, avoid: A.avoidPositions, requiredSlots: A.requiredSlots, totalRounds: A.totalRounds };
    if (!A.myTeamName) say('WARNING: myTeamName unset - position caps are inert');
    say('targets set (' + A.targets.length + ')');
    setTimeout(() => { A.syncQueue().catch((e) => say('ERR syncQueue: ' + e.message)); }, 500);
    return A.targets.length;
  };

  // Mirror the top targets into ESPN's native pick queue. The queue is the
  // failsafe of last resort: if the page (or the whole machine) dies and
  // ESPN's autopick takes over, it drafts from the queue before its own
  // ranks. Call after each targets update; aborts instantly if we go on the
  // clock so it never fights pickSequence for the search box.
  A.syncQueue = async (n = 8) => {
    // Empty the queue first: stale entries outrank the new plan, and ESPN's
    // autopick drafts the queue top-down.
    for (const rm of buttons().filter((b) => b.textContent.trim().toUpperCase() === 'REMOVE')) {
      rm.click();
      await sleep(150);
    }
    for (const name of A.targets.slice(0, n)) {
      if (A.picking || draftButtons().length) { setSearch(''); return say('syncQueue aborted: on the clock'); }
      if (!setSearch(name)) return say('syncQueue: no search box');
      await sleep(500);
      // Re-check after the sleep too: going on the clock mid-await is exactly
      // the race that let this loop retype the box under pickSequence's feet.
      if (A.picking || draftButtons().length) { setSearch(''); return say('syncQueue aborted: on the clock'); }
      const q = [...document.querySelectorAll('button.action-btn')]
        .filter((b) => b.textContent.trim().toUpperCase() === 'QUEUE')
        .find((b) => rowName(b) === name);
      if (q && !posAllowed(rowPos(q))) { say('not queueing ' + name + ': ' + rowPos(q) + ' capped or not a must-fill slot'); continue; }
      if (q) { q.click(); say('QUEUED: ' + name); await sleep(300); }
    }
    setSearch('');
    say('syncQueue done');
  };

  // Selector health probe. Run it right after injection AND again the moment
  // the draft UI mounts (the pre-draft room and the live room may render
  // different DOMs; ESPN can remount everything when picking begins). Healthy
  // live room: searchBox true, posEls > 0, and every actionBtnSample entry
  // has a non-empty pos. Pre-clock some of these being 0/false is expected -
  // re-run at first pick before trusting the guards.
  A.selfTest = () => {
    const btns = [...document.querySelectorAll('button.action-btn')];
    return {
      searchBox: !!searchBox(),
      actionBtns: btns.length,
      posEls: document.querySelectorAll('.playerinfo__playerpos').length,
      actionBtnSample: btns.slice(0, 5).map((b) => ({ name: rowName(b), pos: rowPos(b) })),
      picksParsed: [...(window.__seenPicksSet || [])].filter((t) => PICK_RE.test(t)).length,
      myTeamName: A.myTeamName,
      myPosCounts: myPosCounts(),
    };
  };

  A.stop = () => { A.intervals.forEach(clearInterval); A.intervals = []; say('agent stopped'); };

  say('draft agent installed (' + A.targets.length + ' targets)');
})();

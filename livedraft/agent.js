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
//   2. Paste/inject this whole file.
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

  const A = (window.__draftAgent = {
    targets: [],          // ranked player display names, best first
    avoidPositions: [],   // e.g. ['K','D/ST','QB'] - skipped by the fallback pick
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
  const rowPos = (btn) => {
    const txt = btn.closest('.fixedDataTableCellGroupLayout_cellGroup')?.textContent || '';
    return (txt.match(/\b(QB|RB|WR|TE|K|D\/ST)\b/) || [])[1] || '';
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
    await sleep(500); // let the filter re-render
    const hit = draftButtons().find((b) => rowName(b) === name);
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
    const deadline = Date.now() + 8000;
    try {
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
        // Fallback: best listed player whose position we aren't avoiding.
        const pick = btns.find((b) => !A.avoidPositions.includes(rowPos(b))) || btns[0];
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
    picksSeen: [...seen],
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
      const q = [...document.querySelectorAll('button.action-btn')]
        .filter((b) => b.textContent.trim().toUpperCase() === 'QUEUE')
        .find((b) => rowName(b) === name);
      if (q) { q.click(); say('QUEUED: ' + name); await sleep(300); }
    }
    setSearch('');
    say('syncQueue done');
  };

  A.stop = () => { A.intervals.forEach(clearInterval); A.intervals = []; say('agent stopped'); };

  say('draft agent installed (' + A.targets.length + ' targets)');
})();

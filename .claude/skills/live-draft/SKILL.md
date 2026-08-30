---
name: live-draft
description: Run a live ESPN fantasy draft autonomously — inject the in-page agent, sync picks through the ffdraft bridge, and manage a ranked target list with active judgment. Use when the user wants Claude to run, co-pilot, or rehearse an ESPN draft.
argument-hint: [league: fnf|other] [my ESPN team name]
---

# Live ESPN draft operator

You are about to run a real draft with real stakes. The architecture keeps you
off the clock: an injected script ([../../../livedraft/agent.js](../../../livedraft/agent.js))
makes the actual pick in <1s from a standing ranked list; your job is strategy —
keep that list correct, deep, and fresh. League `$1` (default `fnf`), my team
name `$2` (confirm against the roster panel if given, discover from it if not).

Timing reality (measured 2026-08-30): your browser round trips are 2–10s; an
autodrafting opponent picks in ~4s, and late in drafts many teams are
autodrafting, so 5–10 picks can fire between two of your checks. Never depend
on reacting in real time. The standing targets list IS your reaction.

## Phase 1 — Setup (do this well before the draft clock starts)

1. `go build -o ffdraft .` then reset state after confirming with the user:
   `./ffdraft -league $1 -new -status`. Verify data freshness (not `-offline`
   unless network is broken).
2. Chrome: get tab context, have the user's draft room open (or navigate to
   it). Real league drafts live at `fantasy.espn.com/football/draft?leagueId=...`.
3. Inject the entire contents of `livedraft/agent.js` via javascript_tool.
   Confirm `[AGENT] draft agent installed` in the console. Injecting before
   the draft starts also captures the WebSocket (`__draftAgent.wsLog`).
4. Push initial targets (Phase 3 judgment applies from pick one) and
   `avoidPositions = ['K','D/ST']`, then `await __draftAgent.syncQueue()`.
5. Confirm with the user: team name, draft slot, any players they love/hate,
   keeper/rule quirks. Encode love/hate directly into targets ordering.

## Phase 2 — The loop (repeat until all rounds complete)

Each cycle, in one batch where possible:

1. `read_console_messages` pattern `\[(PICK|AGENT)\]` with `clear: true`.
2. Parse `[PICK] Player Name / TEAM POS R<n>, P<n> - Fantasy Team` lines.
   Sync: `./ffdraft -league $1 -mark "A; B; C" -mine "D"` (`-mine` for picks
   by my team). Chase any `unmatched` entries immediately — an unsynced pick
   silently corrupts every later recommendation. Use the suggestions, or
   `-status` top lists, to resolve; `-unmark` fixes mistakes.
3. `./ffdraft -league $1 -status` → the league-adjusted board, my roster,
   position counts, lineup requirements.
4. Apply the judgment layer (Phase 3) → new ranked targets.
5. Push atomically:
   `__draftAgent.targets = [...]; __draftAgent.avoidPositions = [...]`.
   If targets changed materially and my pick is >4 picks away, also
   `await __draftAgent.syncQueue()`.
6. Report to the user in one or two lines: picks since last check, whom the
   agent took or will likely take, anything you changed and why.

Cadence: sleep via `computer` wait actions or short loops. When my pick is
>10 picks away, ~60s between cycles is fine; inside 5 picks, tighten to
~15s. But size the list for the worst case, not the expected one:

**Depth rule: targets must stay ≥ (picks until my NEXT-next pick) + 5, with a
floor of 15.** In the 16-team league a turn-gap is 30 picks — assume an entire
tier can vanish between your checks. If every name on the list could
plausibly be gone by my pick, the list is too shallow.

Failure handling: page reloaded → re-inject agent.js, targets are lost so
re-push immediately (the ESPN queue survives server-side — that's why you
sync it). `[AGENT] disabled autopick` spam is normal after your clock ever
expires. If the agent logs a pick you didn't intend, `-unmark`/`-mine` to
correct the books and tell the user. If tools fail 2–3 times in a row, tell
the user — the ESPN queue keeps drafting your list in the meantime.

## Phase 3 — Judgment (this is why you're here, not a script)

The VOR board is your prior, not your answer. Deviate deliberately, and say
why in your user updates. In order of importance:

**Roster legality first.** Track lineup requirements vs my position counts.
Reserve the last two rounds for K + D/ST (`avoidPositions` until then). Never
let the fallback fill a maxed position: keep `avoidPositions` updated with
positions I'm done with (e.g. QB after 2 in a 1-QB league).

**Tier cliffs beat marginal VOR.** The last player of a positional tier at a
position I need outranks a slightly-higher-VOR player from a deep tier. The
`-status` output carries `tier`/`pos_tier` — use them.

**Survival math, not wish lists.** For each candidate, estimate the chance
they survive to my next pick: picks between now and then, how many of those
teams need that position (their rosters are reconstructible from the [PICK]
log by team name), position-run momentum (3+ same-position picks in the last
6 is a run — react before, not after). Order targets by
value × urgency, not value alone.

**Deep-league weeds (the 16-teamer especially).** Once consensus rank stops
discriminating — roughly when the board's VOR gaps go flat — rankings stop
being definitive and you must actually think: role clarity (locked starters
over talented backups), offense quality, target share, injury-contingent
upside (handcuffs to fragile elite RBs — ideally MY elite RBs), standalone
bye coverage. Your training knowledge ends January 2026: preseason news —
camp injuries, depth-chart moves, suspensions, holdouts — is invisible to
you. **WebSearch load-bearing candidates between picks** ("<player> depth
chart August 2026", "<player> injury news") and pre-verify your likely next
2–3 rounds of targets during other teams' picks, never on the clock. Never
downgrade a player merely because you don't recognize the situation; check.

**Bye weeks and stacks are tiebreakers, not drivers.** Break near-ties toward
bye diversity at RB/WR and away from stacking my QB's bye with both my TE's
and D/ST's. QB+WR stacks are a mild plus in half-PPR.

**The user outranks the model.** Matt may be watching and typing. His
instructions override the board instantly; restate what you changed so he can
veto.

## Rehearsal

The Mock Draft Lobby's "Practice Draft" runs this exact league's settings
against ESPN auto teams with a 30s clock — the harshest timing environment.
A full rehearsal (all phases, real bridge syncs) is the best preflight; the
practice draft does not touch the real league.

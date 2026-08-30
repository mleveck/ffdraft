# ffdraft

A terminal draft board for fantasy football. Boris Chen's expert-consensus
tiers, re-interleaved for your league's exact scoring and lineup, with live
ESPN ADP — in a keyboard-driven TUI built for the 60-second pick clock.

## TL;DR — running it

```sh
go build -o ffdraft .          # or use the committed binary

./ffdraft                      # friends & family league (standard scoring, 10 teams)
./ffdraft -league other        # the-other-ff (half PPR, 16 teams)

./ffdraft -new                 # start over (clears saved draft state)
./ffdraft -classic             # plain Boris Chen board, no league adjustment
./ffdraft -offline             # prefer cached data (draft-day network hiccups)
./ffdraft -dump                # print the board to stdout and exit
```

Data is fetched fresh on every start (Boris Chen tiers + ESPN projections/ADP)
and cached under `~/.ffdraft/cache/`, so a dead network falls back to the last
good copy. Draft state auto-saves per league to `~/.ffdraft_state_<league>.json`.

### Controls

| Key | Action |
|---|---|
| type anything | fuzzy-search players (always focused) |
| `Ctrl+P` / `Ctrl+N` | move cursor up / down |
| `Enter` | mark player drafted (or undo if drafted) |
| `Tab` / `Shift+Tab`, `1`–`7` | cycle / jump position filter (ALL, RB, WR, TE, QB, DST, K) |
| `qq` | clear search |
| `Ctrl+D` twice | quit |

### Columns

- **Tier** — global board tier in the ALL view; Boris's finer per-position
  tier when a position filter is active.
- **VOR** — league-adjusted value over replacement (season points). Watch for
  cliffs: a 20-point drop to the next player at a position is a reason to
  reach; a flat stretch is a reason to wait.
- **ADP** — ESPN's live average draft position: the best predictor of when a
  player will actually go in an ESPN draft room.

---

# Design

## The problem

Boris Chen (borischen.co) publishes the best freely available draft tiers:
FantasyPros expert-consensus ranks, clustered into tiers with a Gaussian
mixture model (`mclust`, per his [fftiers](https://github.com/borisachen/fftiers)
repo). Two things make his combined board wrong for a specific league:

1. **Scoring**: his pipeline never sees points. Ranks in, clusters out. The
   scoring assumption lives upstream in the experts' heads, and it's a
   generic league — no 100-yard-game bonuses, no league-specific quirks.
2. **Lineup structure**: expert overall ranks assume a typical ~12-team,
   1-flex league. A 10-team league with 2 FLEX slots (or a 16-team league)
   has materially different positional scarcity — e.g. 2-FLEX standard
   scoring starts ~46 RB/WR/TE league-wide instead of ~36, pushing RB/WR
   replacement level far deeper and making generic boards overrank QBs.

A previous iteration of this tool also hand-transcribed player data into Go
source, which produced transcription errors. Rule one of this design:
**no hand-entered player data — everything is fetched and joined at runtime.**

## Data sources (all unauthenticated)

| Source | What it provides |
|---|---|
| `s3-us-west-1.amazonaws.com/fftiers/out/weekly-<POS>[-HALF].csv` | Boris's per-position tiers and ordering (half-PPR variants for RB/WR/TE; QB/K/DST are scoring-invariant) |
| `.../weekly-ALL[-HALF-PPR].csv` | Boris's combined board (classic mode) |
| ESPN `lm-api-reads.fantasy.espn.com/.../leaguedefaults/3?view=kona_player_info` | Live ADP + full-season stat-line projections per player |
| ESPN `.../seasons/<yr>?view=proTeamSchedules_wl` | Team abbreviations and bye weeks |

Boris's files carry only rank statistics (no ADP, team, or bye); ESPN fills
those. Names are joined by normalization (case/punctuation folding, Jr/Sr/III
suffix stripping, DST nickname matching, unique-lastname-within-position
fallback). A failed join blanks the display fields — it can never affect
ordering.

## The board (VOR mode, default)

The core idea: **Boris is trusted for player-vs-player order within a
position; the league's own economics decide how positions interleave.**

1. **Within-position order and tiers** come from Boris's per-position files,
   verbatim. These are finer-grained than his combined chart and are the
   part of his output least sensitive to league settings.

2. **League-exact points**: ESPN's season stat-line projections are scored
   with the league's actual rules (yardage, TDs, receptions, turnovers).
   Per-game yardage bonuses are priced as expected value:

   ```
   E[bonus] = games × [ B100·P(100 ≤ Y < 200) + B200·P(Y ≥ 200) ]
   ```

   where single-game yardage Y is modeled as Normal(μ, cv·μ) around the
   projected per-game mean μ. Dispersion constants come from historical game
   logs: rush cv ≈ 0.55, receiving ≈ 0.65, passing ≈ 0.28. This correctly
   captures the convexity that favors bell-cow volume players.

3. **Value over replacement**: replacement level per position is derived
   from the league's actual lineup — mandatory slots filled by the best
   players at each position, then flex slots allocated greedily to the best
   remaining RB/WR/TE, across all teams. VOR = points − replacement. This is
   where 2-FLEX/10-team vs 1-FLEX/16-team stops being a vibe and becomes
   arithmetic.

4. **Isotonic assignment**: within each position, the sorted VOR values are
   assigned to Boris's order. Projections shape the *value curve* for the
   position; Boris alone decides *which player* sits at each point on it. An
   outlier ESPN projection can never reorder his ranks. Players deeper than
   ESPN's projection pool extend the curve downward.

5. **Interleave and tier**: all positions merge, sorted by assigned VOR.
   Global display tiers come from optimal 1-D k-means (exact dynamic
   programming over the sorted values) — the same clustering idea Boris
   applies to expert ranks, applied to league-adjusted value. The tier count
   is fixed at ~n/8, mirroring his combined charts; BIC-style automatic
   selection was tried and rejected because it over-splits tightly grouped
   values.

6. **K and D/ST** keep their own Boris tier files and sit below the skill
   board (their projections are unreliable and nobody should draft them
   early anyway).

7. **Deep pool**: ESPN-projected players Boris doesn't rank are appended
   with VOR capped strictly below their position's Boris tail, and every
   unranked DST/K follows the ranked ones by ADP. A 16-team draft makes
   ~240 picks; the board carries ~400 players so every pick stays
   trackable, and the cap guarantees pool players never outrank a Boris
   opinion.

`-classic` skips steps 2–5 and reproduces Boris's combined board directly.
The VOR path also degrades to classic automatically if ESPN data is
unavailable and uncached.

## League profiles

Leagues are declared in `data/loader.go` with their scoring, lineup, and
Boris file set:

| | `fnf` (default) | `other` |
|---|---|---|
| Scoring | standard + per-game yardage bonuses | half PPR |
| Teams | 10 | 16 |
| Starters | QB, 2RB, 2WR, TE, 2 FLEX, D/ST, K | QB, 2RB, 2WR, TE, 1 FLEX, D/ST, K |

Adding a league = adding one `League` literal. The season is a constant in
the same file — bump it yearly.

## Architecture

```
main.go            flags, league selection, -dump
data/loader.go     league profiles, fetching + caching, name joining, board assembly
data/projections.go  scoring rules, bonus EV model, replacement levels
data/tiers1d.go    optimal 1-D k-means tiering
model/             Player, Draft, persisted state
ui/                Bubble Tea TUI (search, filters, table)
```

Reliability posture (it has to work during a live draft): network-first with
cache fallback on every fetch, VOR→classic degradation, per-league state
files, and the committed binary as a last resort. Core math
(`LeaguePoints`, `ReplacementLevels`, `ClusterDesc`) is unit-tested; the
data joins are validated empirically via `-dump`.

## Live draft agent (experimental)

`livedraft/agent.js` is an in-page script for having Claude run the draft
autonomously via browser automation. Injected into the ESPN draft room, it
keeps autopick off, logs every completed pick to the console (`[PICK]` tags),
and when you're on the clock drafts the first available player from a ranked
`__draftAgent.targets` list — search-filtered and name-verified before
clicking, reacting in under a second. Claude then works at a comfortable
cadence: read the pick log, recompute the board with ffdraft, update the
target list. Validated end-to-end in a full 17-round ESPN practice draft
(2026-08-30); see the header comment in the file for usage and the DOM
gotchas it works around.

The bridge back into this tool is a set of CLI flags (see `bridge.go`):
`-mark`/`-mine`/`-unmark` sync observed picks into the same per-league state
file the TUI uses (fuzzy name matching, D/ST nicknames, suggestions for
ambiguous inputs), and `-status`/`-targets N` emit the current
league-adjusted board as JSON. `~/.ffdraft_live_<league>.json` additionally
tracks which picks are mine. Don't run the bridge and the TUI at the same
time; both rewrite the state file.

The whole operation — setup, sync loop, judgment guidance, failure handling —
is encoded as a project skill: `/live-draft [league] [my team name]`
(`.claude/skills/live-draft/SKILL.md`).

## Known limitations

- ESPN projections are a single source; projection error flows into the
  interleave (though never into within-position order).
- The bonus model prices season-level EV, not week-to-week variance.
- VOR uses full-season points and ignores weekly lineup dynamics, bench
  value, and streaming — the standard simplification.
- D/ST and K are not value-modeled, just appended.

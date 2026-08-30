package main

// Live-draft bridge: a non-interactive CLI surface over the same board and
// persisted state the TUI uses, so an agent (see livedraft/agent.js) can sync
// picks observed in the ESPN draft room and read back the current
// league-adjusted board as JSON. Don't run it while the TUI is open — both
// rewrite the whole state file on save.
//
//	./ffdraft -league other -mark "Bijan Robinson; Broncos D/ST"
//	./ffdraft -league other -mine "Puka Nacua"      # my pick (implies -mark)
//	./ffdraft -league other -unmark "Bijan Robinson"
//	./ffdraft -league other -status                 # roster, needs, top available
//	./ffdraft -league other -targets 20             # top-N available, board order
//
// Names are matched against board names via the loader's normalization, with
// a unique-last-name fallback and D/ST nickname handling ("Broncos D/ST",
// "Denver Broncos", and "Broncos" all match). Unmatched names come back in
// the JSON with suggestions instead of failing the whole sync.

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"ffdraft/data"
	"ffdraft/model"
)

type bridgeOpts struct {
	mark    string
	unmark  string
	mine    string
	status  bool
	targets int
}

func (o bridgeOpts) active() bool {
	return o.mark != "" || o.unmark != "" || o.mine != "" || o.status || o.targets > 0
}

type unmatchedName struct {
	Input       string   `json:"input"`
	Suggestions []string `json:"suggestions,omitempty"`
}

type playerJSON struct {
	Tier    int     `json:"tier"`
	PosTier int     `json:"pos_tier"`
	Rank    int     `json:"rank"`
	Name    string  `json:"name"`
	Team    string  `json:"team"`
	Pos     string  `json:"pos"`
	Bye     int     `json:"bye"`
	ADP     float64 `json:"adp"`
	VOR     float64 `json:"vor"`
}

type bridgeOutput struct {
	League       string                  `json:"league"`
	Teams        int                     `json:"teams"`
	DraftedTotal int                     `json:"drafted_total"`
	Marked       []string                `json:"marked,omitempty"`
	Unmarked     []string                `json:"unmarked,omitempty"`
	Unmatched    []unmatchedName         `json:"unmatched,omitempty"`
	Lineup       *data.Lineup            `json:"lineup,omitempty"`
	MyTeam       []playerJSON            `json:"my_team,omitempty"`
	MyPosCounts  map[string]int          `json:"my_pos_counts,omitempty"`
	TopAvailable []playerJSON            `json:"top_available,omitempty"`
	ByPosition   map[string][]playerJSON `json:"by_position,omitempty"`
	Targets      []string                `json:"targets,omitempty"`
}

func toJSON(p model.Player) playerJSON {
	return playerJSON{
		Tier: p.Tier, PosTier: p.PosTier, Rank: p.Rank, Name: p.Name,
		Team: p.Team, Pos: string(p.Position), Bye: p.ByeWeek, ADP: p.ADP, VOR: p.VOR,
	}
}

// nameIndex matches externally-sourced player names to board indices.
type nameIndex struct {
	players   []model.Player
	byNorm    map[string]int
	byLast    map[string][]int
	dstByNick map[string]int
}

func buildIndex(players []model.Player) *nameIndex {
	ix := &nameIndex{
		players:   players,
		byNorm:    map[string]int{},
		byLast:    map[string][]int{},
		dstByNick: map[string]int{},
	}
	for i, p := range players {
		norm := data.NormalizeName(p.Name)
		ix.byNorm[norm] = i
		if p.Position == model.DST {
			for _, w := range dstNickWords(norm) {
				ix.dstByNick[w] = i
			}
			continue
		}
		words := strings.Fields(norm)
		if len(words) > 1 {
			last := words[len(words)-1]
			ix.byLast[last] = append(ix.byLast[last], i)
		}
	}
	return ix
}

// dstNickWords strips D/ST markers and returns candidate nickname words,
// e.g. "denver broncos dst" -> ["denver", "broncos"].
func dstNickWords(norm string) []string {
	var out []string
	for _, w := range strings.Fields(norm) {
		if w != "dst" && w != "d" && w != "st" && w != "defense" {
			out = append(out, w)
		}
	}
	return out
}

func looksLikeDST(norm string) bool {
	for _, w := range strings.Fields(norm) {
		if w == "dst" {
			return true
		}
	}
	return false
}

// match resolves one input name to a board index, or returns suggestions.
func (ix *nameIndex) match(input string) (int, []string) {
	norm := data.NormalizeName(input)
	if norm == "" {
		return -1, nil
	}
	if i, ok := ix.byNorm[norm]; ok {
		return i, nil
	}
	if looksLikeDST(norm) {
		words := dstNickWords(norm)
		for j := len(words) - 1; j >= 0; j-- {
			if i, ok := ix.dstByNick[words[j]]; ok {
				return i, nil
			}
		}
		return -1, nil
	}
	words := strings.Fields(norm)
	last := words[len(words)-1]
	if cands := ix.byLast[last]; len(cands) == 1 {
		return cands[0], nil
	} else if len(cands) > 1 {
		var sugg []string
		for _, c := range cands {
			sugg = append(sugg, ix.players[c].Name)
		}
		return -1, sugg
	}
	return -1, nil
}

// liveState is bridge-only persistence (kept apart from the TUI's state file
// so a TUI save can't clobber it): which drafted players are mine.
type liveState struct {
	MyPlayers []string `json:"my_players"`
}

func liveStatePath(leagueKey string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".ffdraft_live_"+leagueKey+".json"), nil
}

func loadLiveState(leagueKey string) liveState {
	var ls liveState
	path, err := liveStatePath(leagueKey)
	if err != nil {
		return ls
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return ls
	}
	json.Unmarshal(raw, &ls)
	return ls
}

func saveLiveState(leagueKey string, ls liveState) error {
	path, err := liveStatePath(leagueKey)
	if err != nil {
		return err
	}
	raw, err := json.MarshalIndent(ls, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0644)
}

func splitNames(arg string) []string {
	var out []string
	for _, s := range strings.Split(arg, ";") {
		if s = strings.TrimSpace(s); s != "" {
			out = append(out, s)
		}
	}
	return out
}

func runBridge(league data.League, offline, classic, newDraft bool, opts bridgeOpts) error {
	model.StateLeagueKey = league.Key
	if newDraft {
		if err := model.ClearDraftState(); err != nil {
			return err
		}
		if path, err := liveStatePath(league.Key); err == nil {
			os.Remove(path)
		}
	}

	players, err := data.LoadPlayers(league, offline, classic)
	if err != nil {
		return err
	}
	state, err := model.LoadDraftState()
	if err != nil {
		return err
	}
	live := loadLiveState(league.Key)
	ix := buildIndex(players)

	out := bridgeOutput{League: league.Key, Teams: league.Teams}
	mine := map[string]bool{}
	for _, n := range live.MyPlayers {
		mine[n] = true
	}

	resolve := func(inputs []string, markMine bool) {
		for _, in := range inputs {
			i, sugg := ix.match(in)
			if i < 0 {
				out.Unmatched = append(out.Unmatched, unmatchedName{Input: in, Suggestions: sugg})
				continue
			}
			name := players[i].Name
			state.DraftedPlayers[name] = true
			out.Marked = append(out.Marked, name)
			if markMine && !mine[name] {
				mine[name] = true
				live.MyPlayers = append(live.MyPlayers, name)
			}
		}
	}
	resolve(splitNames(opts.mark), false)
	resolve(splitNames(opts.mine), true)

	for _, in := range splitNames(opts.unmark) {
		i, sugg := ix.match(in)
		if i < 0 {
			out.Unmatched = append(out.Unmatched, unmatchedName{Input: in, Suggestions: sugg})
			continue
		}
		name := players[i].Name
		delete(state.DraftedPlayers, name)
		out.Unmarked = append(out.Unmarked, name)
		if mine[name] {
			delete(mine, name)
			for j, n := range live.MyPlayers {
				if n == name {
					live.MyPlayers = append(live.MyPlayers[:j], live.MyPlayers[j+1:]...)
					break
				}
			}
		}
	}

	// Keep the counters the TUI shows roughly honest: one pick per marked player.
	state.TotalTeams = league.Teams
	drafted := len(state.DraftedPlayers)
	state.CurrentRound = drafted/league.Teams + 1
	state.CurrentPick = drafted%league.Teams + 1

	if opts.mark != "" || opts.mine != "" || opts.unmark != "" || newDraft {
		if err := model.SaveDraftState(state); err != nil {
			return err
		}
		if err := saveLiveState(league.Key, live); err != nil {
			return err
		}
	}

	out.DraftedTotal = drafted

	if opts.status {
		lineup := league.Lineup
		out.Lineup = &lineup
		out.MyPosCounts = map[string]int{}
		for i := range players {
			if mine[players[i].Name] {
				out.MyTeam = append(out.MyTeam, toJSON(players[i]))
				out.MyPosCounts[string(players[i].Position)]++
			}
		}
		perPos := map[model.Position]int{
			model.QB: 6, model.RB: 12, model.WR: 12, model.TE: 6, model.DST: 4, model.K: 4,
		}
		out.ByPosition = map[string][]playerJSON{}
		for _, p := range players {
			if state.DraftedPlayers[p.Name] {
				continue
			}
			if len(out.TopAvailable) < 30 {
				out.TopAvailable = append(out.TopAvailable, toJSON(p))
			}
			if n := perPos[p.Position]; len(out.ByPosition[string(p.Position)]) < n {
				out.ByPosition[string(p.Position)] = append(out.ByPosition[string(p.Position)], toJSON(p))
			}
		}
	}

	if opts.targets > 0 {
		for _, p := range players {
			if state.DraftedPlayers[p.Name] {
				continue
			}
			out.Targets = append(out.Targets, p.Name)
			if len(out.Targets) == opts.targets {
				break
			}
		}
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(out); err != nil {
		return fmt.Errorf("encoding output: %w", err)
	}
	return nil
}

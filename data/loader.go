package data

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"ffdraft/model"
)

const borisBase = "https://s3-us-west-1.amazonaws.com/fftiers/out/"

type League struct {
	Key         string
	DisplayName string
	BorisALL    string            // combined board, used by the classic path
	BorisPos    map[string]string // per-position tier files for the VOR path
	BorisDST    string
	BorisK      string
	Teams       int
	Lineup      Lineup
	Scoring     Scoring
}

var Leagues = map[string]League{
	"fnf": {
		Key:         "fnf",
		DisplayName: "Friends, Fiends n Family (Standard, 10tm)",
		BorisALL:    borisBase + "weekly-ALL.csv",
		BorisPos: map[string]string{
			"QB": borisBase + "weekly-QB.csv",
			"RB": borisBase + "weekly-RB.csv",
			"WR": borisBase + "weekly-WR.csv",
			"TE": borisBase + "weekly-TE.csv",
		},
		BorisDST: borisBase + "weekly-DST.csv",
		BorisK:   borisBase + "weekly-K.csv",
		Teams:    10,
		Lineup:   Lineup{QB: 1, RB: 2, WR: 2, TE: 1, Flex: 2},
		Scoring: Scoring{
			PassYd: 0.04, PassTD: 4, Int: -2,
			RushYd: 0.1, RushTD: 6,
			Rec: 0, RecYd: 0.1, RecTD: 6,
			FumLost:  -2,
			Bonus100: 1, Bonus200: 1, Pass400: 1,
		},
	},
	"other": {
		Key:         "other",
		DisplayName: "the-other-ff (Half PPR, 16tm)",
		BorisALL:    borisBase + "weekly-ALL-HALF-PPR.csv",
		BorisPos: map[string]string{
			"QB": borisBase + "weekly-QB.csv",
			"RB": borisBase + "weekly-RB-HALF.csv",
			"WR": borisBase + "weekly-WR-HALF.csv",
			"TE": borisBase + "weekly-TE-HALF.csv",
		},
		BorisDST: borisBase + "weekly-DST.csv",
		BorisK:   borisBase + "weekly-K.csv",
		Teams:    16,
		Lineup:   Lineup{QB: 1, RB: 2, WR: 2, TE: 1, Flex: 1},
		Scoring: Scoring{
			PassYd: 0.04, PassTD: 4, Int: -2,
			RushYd: 0.1, RushTD: 6,
			Rec: 0.5, RecYd: 0.1, RecTD: 6,
			FumLost: -2,
		},
	},
}

func cacheDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".ffdraft", "cache")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}

var httpClient = &http.Client{Timeout: 15 * time.Second}

// fetchWithCache GETs url (with optional headers), falling back to the cached
// copy on any network failure. Successful fetches refresh the cache. With
// offline set, the cache is preferred and the network is only a fallback.
func fetchWithCache(url, cacheName string, headers map[string]string, offline bool) ([]byte, error) {
	dir, err := cacheDir()
	if err != nil {
		return nil, err
	}
	cachePath := filepath.Join(dir, cacheName)

	if offline {
		if data, err := os.ReadFile(cachePath); err == nil {
			return data, nil
		}
	}

	data, fetchErr := func() ([]byte, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("GET %s: HTTP %d", url, resp.StatusCode)
		}
		return io.ReadAll(resp.Body)
	}()

	if fetchErr != nil {
		if cached, err := os.ReadFile(cachePath); err == nil {
			fmt.Fprintf(os.Stderr, "warning: fetch failed (%v), using cached %s\n", fetchErr, cacheName)
			return cached, nil
		}
		return nil, fetchErr
	}

	_ = os.WriteFile(cachePath, data, 0644)
	return data, nil
}

func fetchBoris(fileURL string, offline bool) ([]borisRow, error) {
	u, _ := url.Parse(fileURL)
	data, err := fetchWithCache(fileURL, path.Base(u.Path), nil, offline)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path.Base(u.Path), err)
	}
	rows, err := parseBorisCSV(data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path.Base(u.Path), err)
	}
	return rows, nil
}

type borisRow struct {
	Rank     int
	Name     string
	Tier     int
	Position string
	AvgRank  float64
}

// parseBorisCSV maps columns by header name since the per-position files use
// a different column order than the ALL file.
func parseBorisCSV(data []byte) ([]borisRow, error) {
	r := csv.NewReader(bytes.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) < 2 {
		return nil, fmt.Errorf("boris csv: no data rows")
	}
	col := map[string]int{}
	for i, h := range records[0] {
		col[h] = i
	}
	for _, want := range []string{"Rank", "Player.Name", "Tier", "Position"} {
		if _, ok := col[want]; !ok {
			return nil, fmt.Errorf("boris csv: missing column %q", want)
		}
	}
	var rows []borisRow
	for _, rec := range records[1:] {
		rank, _ := strconv.Atoi(rec[col["Rank"]])
		tier, _ := strconv.Atoi(rec[col["Tier"]])
		avg := 0.0
		if i, ok := col["Avg.Rank"]; ok {
			avg, _ = strconv.ParseFloat(rec[i], 64)
		}
		rows = append(rows, borisRow{
			Rank:     rank,
			Name:     rec[col["Player.Name"]],
			Tier:     tier,
			Position: rec[col["Position"]],
			AvgRank:  avg,
		})
	}
	return rows, nil
}

type espnInfo struct {
	Name string
	Team string
	Bye  int
	ADP  float64
	Proj *ProjStats
}

type espnPlayerResp struct {
	Players []struct {
		Player struct {
			FullName          string `json:"fullName"`
			DefaultPositionID int    `json:"defaultPositionId"`
			ProTeamID         int    `json:"proTeamId"`
			Ownership         struct {
				AverageDraftPosition float64 `json:"averageDraftPosition"`
			} `json:"ownership"`
			Stats []struct {
				ID    string             `json:"id"`
				Stats map[string]float64 `json:"stats"`
			} `json:"stats"`
		} `json:"player"`
	} `json:"players"`
}

type espnScheduleResp struct {
	Settings struct {
		ProTeams []struct {
			ID      int    `json:"id"`
			Abbrev  string `json:"abbrev"`
			ByeWeek int    `json:"byeWeek"`
		} `json:"proTeams"`
	} `json:"settings"`
}

var espnPositions = map[int]string{1: "QB", 2: "RB", 3: "WR", 4: "TE", 5: "K", 16: "DST"}

const season = "2026"

// ESPN season-projection stat ids (verified against the 2026 feed).
func projFromStats(m map[string]float64) *ProjStats {
	return &ProjStats{
		PassYds:    m["3"],
		PassTDs:    m["4"],
		Ints:       m["20"],
		RushYds:    m["24"],
		RushTDs:    m["25"],
		RecYds:     m["42"],
		RecTDs:     m["43"],
		Receptions: m["53"],
		FumLost:    m["72"],
		Games:      m["210"],
	}
}

// loadESPN returns a lookup map from normalized player name to ESPN
// team/bye/ADP/projection info, plus a last-name+position fallback index of
// keys into that map.
func loadESPN(offline bool) (byName map[string]espnInfo, byLast map[string][]string, err error) {
	schedData, err := fetchWithCache(
		"https://lm-api-reads.fantasy.espn.com/apis/v3/games/ffl/seasons/"+season+"?view=proTeamSchedules_wl",
		"espn-teams.json", nil, offline)
	if err != nil {
		return nil, nil, err
	}
	var sched espnScheduleResp
	if err := json.Unmarshal(schedData, &sched); err != nil {
		return nil, nil, err
	}
	teamAbbrev := map[int]string{}
	teamBye := map[int]int{}
	for _, t := range sched.Settings.ProTeams {
		teamAbbrev[t.ID] = strings.ToUpper(t.Abbrev)
		teamBye[t.ID] = t.ByeWeek
	}

	// Deep enough to cover every Boris-ranked player; the ADP sort pushes
	// never-drafted players (deep kickers etc.) past a smaller window.
	filter := `{"players":{"limit":1500,"sortAdp":{"sortAsc":true,"sortPriority":1}}}`
	playerData, err := fetchWithCache(
		"https://lm-api-reads.fantasy.espn.com/apis/v3/games/ffl/seasons/"+season+"/segments/0/leaguedefaults/3?view=kona_player_info",
		"espn-players.json", map[string]string{"x-fantasy-filter": filter}, offline)
	if err != nil {
		return nil, nil, err
	}
	var resp espnPlayerResp
	if err := json.Unmarshal(playerData, &resp); err != nil {
		return nil, nil, err
	}

	byName = map[string]espnInfo{}
	byLast = map[string][]string{}
	for _, p := range resp.Players {
		pos, ok := espnPositions[p.Player.DefaultPositionID]
		if !ok {
			continue
		}
		info := espnInfo{
			Name: p.Player.FullName,
			Team: teamAbbrev[p.Player.ProTeamID],
			Bye:  teamBye[p.Player.ProTeamID],
			ADP:  p.Player.Ownership.AverageDraftPosition,
		}
		for _, s := range p.Player.Stats {
			if s.ID == "10"+season && len(s.Stats) > 0 {
				info.Proj = projFromStats(s.Stats)
				break
			}
		}
		if pos == "DST" {
			// ESPN name is like "Texans D/ST"; key off the nickname.
			nick := strings.ToLower(strings.TrimSpace(strings.Replace(p.Player.FullName, "D/ST", "", 1)))
			byName["dst:"+nick] = info
			continue
		}
		key := pos + ":" + normalizeName(p.Player.FullName)
		byName[key] = info
		words := strings.Fields(normalizeName(p.Player.FullName))
		if len(words) > 1 {
			last := words[len(words)-1]
			byLast[pos+":"+last] = append(byLast[pos+":"+last], key)
		}
	}
	return byName, byLast, nil
}

var suffixes = map[string]bool{"jr": true, "sr": true, "ii": true, "iii": true, "iv": true, "v": true}

func normalizeName(name string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(name) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == ' ' {
			b.WriteRune(r)
		}
	}
	words := strings.Fields(b.String())
	for len(words) > 1 && suffixes[words[len(words)-1]] {
		words = words[:len(words)-1]
	}
	return strings.Join(words, " ")
}

// lookupESPN resolves a Boris player name to its ESPN entry, returning the
// byName key of the match so callers can track which entries were consumed.
func lookupESPN(byName map[string]espnInfo, byLast map[string][]string, name, pos string) (espnInfo, string, bool) {
	if pos == "DST" {
		words := strings.Fields(strings.ToLower(name))
		if len(words) > 0 {
			key := "dst:" + words[len(words)-1]
			if info, ok := byName[key]; ok {
				return info, key, true
			}
		}
		return espnInfo{}, "", false
	}
	key := pos + ":" + normalizeName(name)
	if info, ok := byName[key]; ok {
		return info, key, true
	}
	// Fallback: unique last name within position (catches Chig/Chigoziem etc).
	words := strings.Fields(normalizeName(name))
	if len(words) > 1 {
		if keys := byLast[pos+":"+words[len(words)-1]]; len(keys) == 1 {
			return byName[keys[0]], keys[0], true
		}
	}
	return espnInfo{}, "", false
}

// LoadPlayers builds the draft board. The default (VOR) path keeps Boris
// Chen's per-position tiers and ordering as-is, but builds the cross-position
// interleave from league-exact projected value over replacement. The classic
// path reproduces Boris's combined board directly. If ESPN projections are
// unavailable the VOR path falls back to classic.
func LoadPlayers(league League, offline, classic bool) ([]model.Player, error) {
	byName, byLast, espnErr := loadESPN(offline)
	if espnErr != nil {
		byName, byLast = map[string]espnInfo{}, map[string][]string{}
	}

	if !classic && espnErr == nil {
		players, err := loadVORBoard(league, offline, byName, byLast)
		if err == nil {
			return players, nil
		}
		fmt.Fprintf(os.Stderr, "warning: VOR board failed (%v); falling back to classic board\n", err)
	} else if !classic {
		fmt.Fprintf(os.Stderr, "warning: ESPN data unavailable (%v); falling back to classic board\n", espnErr)
	}

	return loadClassicBoard(league, offline, byName, byLast)
}

// loadClassicBoard is Boris's combined ALL board plus his kicker file.
func loadClassicBoard(league League, offline bool, byName map[string]espnInfo, byLast map[string][]string) ([]model.Player, error) {
	allRows, err := fetchBoris(league.BorisALL, offline)
	if err != nil {
		return nil, fmt.Errorf("boris tiers: %w", err)
	}
	// Kickers come exclusively from the dedicated K file below; the ALL file
	// sometimes includes a partial set, which would duplicate them.
	filtered := allRows[:0]
	for _, r := range allRows {
		if r.Position != "K" {
			filtered = append(filtered, r)
		}
	}
	allRows = filtered

	kRows, err := fetchBoris(league.BorisK, offline)
	if err != nil {
		return nil, fmt.Errorf("boris kicker tiers: %w", err)
	}

	maxTier, maxRank := 0, 0
	for _, r := range allRows {
		if r.Tier > maxTier {
			maxTier = r.Tier
		}
		if r.Rank > maxRank {
			maxRank = r.Rank
		}
	}
	maxKTier := 0
	for _, r := range kRows {
		if r.Tier > maxKTier {
			maxKTier = r.Tier
		}
	}
	kTierShift := maxTier - maxKTier
	if kTierShift < 0 {
		kTierShift = 0
	}
	for i := range kRows {
		kRows[i].Tier += kTierShift
		kRows[i].Rank += maxRank
	}

	var players []model.Player
	for _, r := range append(allRows, kRows...) {
		info, _, _ := lookupESPN(byName, byLast, r.Name, r.Position)
		players = append(players, model.Player{
			Tier:     r.Tier,
			PosTier:  r.Tier,
			Rank:     r.Rank,
			Name:     r.Name,
			Team:     info.Team,
			Position: model.Position(r.Position),
			ByeWeek:  info.Bye,
			ADP:      info.ADP,
		})
	}
	return players, nil
}

var skillPositions = []string{"QB", "RB", "WR", "TE"}

// loadVORBoard assembles the league-adjusted board:
//  1. Boris per-position files fix each position's internal order and tiers.
//  2. ESPN projections scored with the league's exact rules give points;
//     VOR = points - replacement level from the league's actual lineup.
//  3. Within each position the sorted VOR values are assigned to Boris's
//     order (projections set the value curve, Boris sets the order).
//  4. Positions interleave by assigned VOR; display tiers come from optimal
//     1-D clustering of those values. DST and K sit below the skill board.
//  5. ESPN-projected players Boris doesn't rank form a deep pool slotted
//     strictly below their position's Boris tail, so a big league's late
//     rounds stay trackable without ever outranking a Boris opinion.
func loadVORBoard(league League, offline bool, byName map[string]espnInfo, byLast map[string][]string) ([]model.Player, error) {
	type boardPlayer struct {
		name    string
		pos     string
		posTier int
		info    espnInfo
		ok      bool
		vor     float64
	}
	matched := map[string]bool{}

	posPlayers := map[string][]boardPlayer{}
	pointsByPos := map[string][]float64{}
	for _, pos := range skillPositions {
		rows, err := fetchBoris(league.BorisPos[pos], offline)
		if err != nil {
			return nil, err
		}
		sort.Slice(rows, func(i, j int) bool { return rows[i].Rank < rows[j].Rank })
		for _, r := range rows {
			info, key, ok := lookupESPN(byName, byLast, r.Name, pos)
			if ok {
				matched[key] = true
			}
			posPlayers[pos] = append(posPlayers[pos], boardPlayer{name: r.Name, pos: pos, posTier: r.Tier, info: info, ok: ok})
		}
	}

	// League-exact points for every projected ESPN player at the position
	// (not just Boris-ranked ones) so replacement levels are well estimated.
	for key, info := range byName {
		if info.Proj == nil {
			continue
		}
		pos := strings.SplitN(key, ":", 2)[0]
		for _, sp := range skillPositions {
			if pos == sp {
				pointsByPos[pos] = append(pointsByPos[pos], LeaguePoints(*info.Proj, league.Scoring))
			}
		}
	}
	repl := ReplacementLevels(pointsByPos, league.Lineup, league.Teams)

	// Isotonic assignment: sorted VOR values meet Boris's order.
	for pos, players := range posPlayers {
		var vors []float64
		for _, p := range players {
			if p.ok && p.info.Proj != nil {
				vors = append(vors, LeaguePoints(*p.info.Proj, league.Scoring)-repl[pos])
			}
		}
		if len(vors) == 0 {
			return nil, fmt.Errorf("no projections matched for %s", pos)
		}
		sort.Sort(sort.Reverse(sort.Float64Slice(vors)))
		// Players beyond the matched pool extend the curve downward.
		for len(vors) < len(players) {
			vors = append(vors, vors[len(vors)-1]-0.5)
		}
		for i := range players {
			players[i].vor = vors[i]
		}
		posPlayers[pos] = players
	}

	// Deep pool: ESPN-projected players Boris doesn't rank. Their VOR is
	// capped just under the position's Boris tail so they can never leapfrog
	// a ranked player (also containing any duplicate from a failed join).
	const deepPoolMinPoints = 20
	minVOR := map[string]float64{}
	maxPosTier := map[string]int{}
	for pos, players := range posPlayers {
		minVOR[pos] = players[len(players)-1].vor
		for _, p := range players {
			if p.posTier > maxPosTier[pos] {
				maxPosTier[pos] = p.posTier
			}
		}
	}
	var deep []boardPlayer
	for key, info := range byName {
		if matched[key] || info.Proj == nil {
			continue
		}
		pos := strings.SplitN(key, ":", 2)[0]
		if _, ok := posPlayers[pos]; !ok {
			continue
		}
		pts := LeaguePoints(*info.Proj, league.Scoring)
		if pts < deepPoolMinPoints {
			continue
		}
		vor := pts - repl[pos]
		if cap := minVOR[pos] - 0.001; vor > cap {
			vor = cap
		}
		deep = append(deep, boardPlayer{name: info.Name, pos: pos, posTier: maxPosTier[pos] + 1, info: info, ok: true, vor: vor})
	}
	sort.Slice(deep, func(i, j int) bool {
		if deep[i].vor != deep[j].vor {
			return deep[i].vor > deep[j].vor
		}
		return deep[i].name < deep[j].name
	})

	var board []boardPlayer
	for _, pos := range skillPositions {
		board = append(board, posPlayers[pos]...)
	}
	board = append(board, deep...)
	sort.SliceStable(board, func(i, j int) bool { return board[i].vor > board[j].vor })

	vals := make([]float64, len(board))
	for i, p := range board {
		vals[i] = p.vor
	}
	tiers := ClusterDesc(vals, BoardTiers(len(vals)))

	var players []model.Player
	for i, p := range board {
		players = append(players, model.Player{
			Tier:     tiers[i],
			PosTier:  p.posTier,
			Rank:     i + 1,
			Name:     p.name,
			Team:     p.info.Team,
			Position: model.Position(p.pos),
			ByeWeek:  p.info.Bye,
			ADP:      p.info.ADP,
			VOR:      p.vor,
		})
	}

	// DST then K below the skill board, keeping their own Boris tier shapes,
	// each followed by the ESPN-only units Boris doesn't rank (a 16-team
	// draft needs all 32 DSTs on the board), ordered by ADP.
	maxTier, rank := 0, len(players)
	for _, p := range players {
		if p.Tier > maxTier {
			maxTier = p.Tier
		}
	}
	for _, unit := range []struct {
		url     string
		pos     string
		keyPref string
	}{
		{league.BorisDST, "DST", "dst:"},
		{league.BorisK, "K", "K:"},
	} {
		rows, err := fetchBoris(unit.url, offline)
		if err != nil {
			return nil, err
		}
		sort.Slice(rows, func(i, j int) bool { return rows[i].Rank < rows[j].Rank })
		unitMaxTier, borisMaxPosTier := maxTier, 0
		for _, r := range rows {
			rank++
			info, key, ok := lookupESPN(byName, byLast, r.Name, r.Position)
			if ok {
				matched[key] = true
			}
			players = append(players, model.Player{
				Tier:     maxTier + r.Tier,
				PosTier:  r.Tier,
				Rank:     rank,
				Name:     r.Name,
				Team:     info.Team,
				Position: model.Position(r.Position),
				ByeWeek:  info.Bye,
				ADP:      info.ADP,
			})
			if maxTier+r.Tier > unitMaxTier {
				unitMaxTier = maxTier + r.Tier
			}
			if r.Tier > borisMaxPosTier {
				borisMaxPosTier = r.Tier
			}
		}
		var extras []espnInfo
		for key, info := range byName {
			if !matched[key] && strings.HasPrefix(key, unit.keyPref) {
				extras = append(extras, info)
			}
		}
		sort.Slice(extras, func(i, j int) bool {
			ai, aj := extras[i].ADP, extras[j].ADP
			if ai <= 0 {
				ai = 9999
			}
			if aj <= 0 {
				aj = 9999
			}
			if ai != aj {
				return ai < aj
			}
			return extras[i].Name < extras[j].Name
		})
		if len(extras) > 0 {
			unitMaxTier++
		}
		for _, info := range extras {
			rank++
			players = append(players, model.Player{
				Tier:     unitMaxTier,
				PosTier:  borisMaxPosTier + 1,
				Rank:     rank,
				Name:     info.Name,
				Team:     info.Team,
				Position: model.Position(unit.pos),
				ByeWeek:  info.Bye,
				ADP:      info.ADP,
			})
		}
		maxTier = unitMaxTier
	}
	return players, nil
}

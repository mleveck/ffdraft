package data

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"ffdraft/model"
)

const borisBase = "https://s3-us-west-1.amazonaws.com/fftiers/out/"

type League struct {
	Key         string
	DisplayName string
	BorisALL    string
	BorisK      string
	Teams       int
}

var Leagues = map[string]League{
	"fnf": {
		Key:         "fnf",
		DisplayName: "Friends, Fiends n Family (Standard, 10tm)",
		BorisALL:    borisBase + "weekly-ALL.csv",
		BorisK:      borisBase + "weekly-K.csv",
		Teams:       10,
	},
	"other": {
		Key:         "other",
		DisplayName: "the-other-ff (Half PPR, 16tm)",
		BorisALL:    borisBase + "weekly-ALL-HALF-PPR.csv",
		BorisK:      borisBase + "weekly-K.csv",
		Teams:       16,
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

type borisRow struct {
	Rank     int
	Name     string
	Tier     int
	Position string
	AvgRank  float64
}

// parseBorisCSV maps columns by header name since the K file uses a
// different column order than the ALL file.
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
	Team string
	Bye  int
	ADP  float64
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

// loadESPN returns lookup maps from normalized player name (and a
// last-name+position fallback key) to ESPN team/bye/ADP info.
func loadESPN(offline bool) (byName map[string]espnInfo, byLast map[string][]espnInfo, err error) {
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

	filter := `{"players":{"limit":600,"sortAdp":{"sortAsc":true,"sortPriority":1}}}`
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
	byLast = map[string][]espnInfo{}
	for _, p := range resp.Players {
		pos, ok := espnPositions[p.Player.DefaultPositionID]
		if !ok {
			continue
		}
		info := espnInfo{
			Team: teamAbbrev[p.Player.ProTeamID],
			Bye:  teamBye[p.Player.ProTeamID],
			ADP:  p.Player.Ownership.AverageDraftPosition,
		}
		if pos == "DST" {
			// ESPN name is like "Texans D/ST"; key off the nickname.
			nick := strings.ToLower(strings.TrimSpace(strings.Replace(p.Player.FullName, "D/ST", "", 1)))
			byName["dst:"+nick] = info
			continue
		}
		byName[pos+":"+normalizeName(p.Player.FullName)] = info
		words := strings.Fields(normalizeName(p.Player.FullName))
		if len(words) > 1 {
			last := words[len(words)-1]
			byLast[pos+":"+last] = append(byLast[pos+":"+last], info)
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

func lookupESPN(byName map[string]espnInfo, byLast map[string][]espnInfo, name, pos string) (espnInfo, bool) {
	if pos == "DST" {
		words := strings.Fields(strings.ToLower(name))
		if len(words) > 0 {
			if info, ok := byName["dst:"+words[len(words)-1]]; ok {
				return info, true
			}
		}
		return espnInfo{}, false
	}
	norm := normalizeName(name)
	if info, ok := byName[pos+":"+norm]; ok {
		return info, true
	}
	// Fallback: unique last name within position (catches Chig/Chigoziem etc).
	words := strings.Fields(norm)
	if len(words) > 1 {
		if matches := byLast[pos+":"+words[len(words)-1]]; len(matches) == 1 {
			return matches[0], true
		}
	}
	return espnInfo{}, false
}

// LoadPlayers builds the draft board for a league: Boris Chen tiers merged
// with ESPN team/bye/ADP. Kickers come from a separate Boris file with their
// own 1..N tiers, so they are shifted to align with the bottom of the
// overall board and ranked after everyone else.
func LoadPlayers(league League, offline bool) ([]model.Player, error) {
	allData, err := fetchWithCache(league.BorisALL, league.Key+"-tiers.csv", nil, offline)
	if err != nil {
		return nil, fmt.Errorf("boris tiers: %w", err)
	}
	allRows, err := parseBorisCSV(allData)
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

	kData, err := fetchWithCache(league.BorisK, "k-tiers.csv", nil, offline)
	if err != nil {
		return nil, fmt.Errorf("boris kicker tiers: %w", err)
	}
	kRows, err := parseBorisCSV(kData)
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

	byName, byLast, err := loadESPN(offline)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: ESPN data unavailable (%v); ADP/team/bye will be blank\n", err)
		byName, byLast = map[string]espnInfo{}, map[string][]espnInfo{}
	}

	var players []model.Player
	for _, r := range append(allRows, kRows...) {
		info, _ := lookupESPN(byName, byLast, r.Name, r.Position)
		players = append(players, model.Player{
			Tier:     r.Tier,
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

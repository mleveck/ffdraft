package data

import (
	"math"
	"sort"
)

// Scoring holds the league's point values for the projected stat categories
// that differ meaningfully across players. Bonus values are per game:
// Bonus100/Bonus200 for 100-199 and 200+ yard rush/rec games, Pass400 for
// 400+ yard passing games.
type Scoring struct {
	PassYd, PassTD, Int float64
	RushYd, RushTD      float64
	Rec, RecYd, RecTD   float64
	FumLost             float64
	Bonus100, Bonus200  float64
	Pass400             float64
}

// Lineup is the number of starters per team at each slot.
type Lineup struct {
	QB, RB, WR, TE, Flex int
}

// ProjStats is an ESPN full-season stat-line projection.
type ProjStats struct {
	PassYds, PassTDs, Ints     float64
	RushYds, RushTDs           float64
	Receptions, RecYds, RecTDs float64
	FumLost                    float64
	Games                      float64
}

// Per-game yardage dispersion (std dev / mean) estimated from historical
// game logs: rushing is steadier than receiving; passing is steadiest.
const (
	rushCV = 0.55
	recCV  = 0.65
	passCV = 0.28
)

// pOver is P(single-game yards >= threshold) for a player whose projected
// per-game mean is mu, modeling game yardage as Normal(mu, cv*mu).
func pOver(mu, threshold, cv float64) float64 {
	if mu <= 0 {
		return 0
	}
	z := (threshold - mu) / (cv * mu)
	return 0.5 * math.Erfc(z/math.Sqrt2)
}

// LeaguePoints scores a season projection with the league's exact rules.
// The per-game yardage bonuses are priced as expected value:
// games * [ Bonus100*P(100<=Y<200) + Bonus200*P(Y>=200) ] per category.
func LeaguePoints(p ProjStats, s Scoring) float64 {
	pts := p.PassYds*s.PassYd + p.PassTDs*s.PassTD + p.Ints*s.Int +
		p.RushYds*s.RushYd + p.RushTDs*s.RushTD +
		p.Receptions*s.Rec + p.RecYds*s.RecYd + p.RecTDs*s.RecTD +
		p.FumLost*s.FumLost

	if s.Bonus100 != 0 || s.Bonus200 != 0 || s.Pass400 != 0 {
		g := p.Games
		if g <= 0 {
			g = 17
		}
		bonusEV := func(mu, cv float64) float64 {
			p200 := pOver(mu, 200, cv)
			p100 := pOver(mu, 100, cv)
			return s.Bonus100*(p100-p200) + s.Bonus200*p200
		}
		perGame := bonusEV(p.RushYds/g, rushCV) +
			bonusEV(p.RecYds/g, recCV) +
			s.Pass400*pOver(p.PassYds/g, 400, passCV)
		pts += g * perGame
	}
	return pts
}

// ReplacementLevels computes the replacement-level points for each skill
// position given the league's lineup: mandatory slots are filled by the
// best players at each position, then flex slots greedily by the best
// remaining RB/WR/TE. Replacement is the best player left unstarted.
func ReplacementLevels(pointsByPos map[string][]float64, lu Lineup, teams int) map[string]float64 {
	sorted := map[string][]float64{}
	for pos, pts := range pointsByPos {
		cp := append([]float64(nil), pts...)
		sort.Sort(sort.Reverse(sort.Float64Slice(cp)))
		sorted[pos] = cp
	}

	idx := map[string]int{
		"QB": teams * lu.QB,
		"RB": teams * lu.RB,
		"WR": teams * lu.WR,
		"TE": teams * lu.TE,
	}
	for i := 0; i < teams*lu.Flex; i++ {
		best, bestVal := "", math.Inf(-1)
		for _, pos := range []string{"RB", "WR", "TE"} {
			if idx[pos] < len(sorted[pos]) && sorted[pos][idx[pos]] > bestVal {
				best, bestVal = pos, sorted[pos][idx[pos]]
			}
		}
		if best == "" {
			break
		}
		idx[best]++
	}

	repl := map[string]float64{}
	for pos, i := range idx {
		vals := sorted[pos]
		switch {
		case len(vals) == 0:
			repl[pos] = 0
		case i < len(vals):
			repl[pos] = vals[i]
		default:
			repl[pos] = vals[len(vals)-1]
		}
	}
	return repl
}

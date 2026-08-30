package data

import (
	"math"
	"testing"
)

func TestPOver(t *testing.T) {
	if p := pOver(100, 100, 0.5); math.Abs(p-0.5) > 1e-9 {
		t.Errorf("pOver at mean = threshold should be 0.5, got %f", p)
	}
	if p := pOver(300, 100, 0.5); p < 0.9 {
		t.Errorf("pOver far above threshold should approach 1, got %f", p)
	}
	if p := pOver(10, 100, 0.5); p > 0.01 {
		t.Errorf("pOver far below threshold should approach 0, got %f", p)
	}
	if p := pOver(0, 100, 0.5); p != 0 {
		t.Errorf("pOver with zero mean should be 0, got %f", p)
	}
}

func TestLeaguePointsNoBonus(t *testing.T) {
	s := Scoring{PassYd: 0.04, PassTD: 4, Int: -2, RushYd: 0.1, RushTD: 6, Rec: 0.5, RecYd: 0.1, RecTD: 6, FumLost: -2}
	p := ProjStats{PassYds: 4000, PassTDs: 30, Ints: 10, RushYds: 300, RushTDs: 3, Receptions: 0, RecYds: 0, RecTDs: 0, FumLost: 2, Games: 17}
	want := 4000*0.04 + 30*4 + 10*-2 + 300*0.1 + 3*6 + 2*-2
	if got := LeaguePoints(p, s); math.Abs(got-want) > 1e-9 {
		t.Errorf("LeaguePoints = %f, want %f", got, want)
	}
}

func TestLeaguePointsBonusEV(t *testing.T) {
	s := Scoring{RushYd: 0.1, Bonus100: 1, Bonus200: 1}
	// 1700 rushing yards over 17 games = 100/game: P(>=100) = 0.5,
	// so expected bonus is 17 * 0.5 = 8.5 on top of 170 yardage points.
	p := ProjStats{RushYds: 1700, Games: 17}
	want := 170.0 + 8.5
	if got := LeaguePoints(p, s); math.Abs(got-want) > 1e-6 {
		t.Errorf("LeaguePoints with bonus = %f, want %f", got, want)
	}
}

func TestReplacementLevels(t *testing.T) {
	pts := map[string][]float64{
		"QB": {30, 20, 10},
		"RB": {30, 25, 20, 15, 10},
		"WR": {28, 22, 16, 12},
		"TE": {5, 4},
	}
	lu := Lineup{QB: 1, RB: 1, WR: 1, TE: 0, Flex: 1}
	repl := ReplacementLevels(pts, lu, 2)
	// Mandatory: 2 QB, 2 RB, 2 WR. Flex x2: RB 20, then WR 16.
	want := map[string]float64{"QB": 10, "RB": 15, "WR": 12, "TE": 5}
	for pos, w := range want {
		if repl[pos] != w {
			t.Errorf("replacement[%s] = %f, want %f", pos, repl[pos], w)
		}
	}
}

package model

import "sort"

type Position string

const (
	QB  Position = "QB"
	RB  Position = "RB"
	WR  Position = "WR"
	TE  Position = "TE"
	K   Position = "K"
	DST Position = "DST"
	ALL Position = "ALL"
)

type Player struct {
	Tier     int
	Rank     int
	Name     string
	Team     string
	Position Position
	ByeWeek  int
	ADP      float64
	Drafted  bool
}

func (p *Player) ToggleDrafted() {
	p.Drafted = !p.Drafted
}

func FilterByPosition(players []Player, pos Position) []Player {
	if pos == ALL {
		return players
	}

	var filtered []Player
	for _, p := range players {
		if p.Position == pos {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func SortPlayersByTierAndADP(players []Player) {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Tier != players[j].Tier {
			return players[i].Tier < players[j].Tier
		}
		if players[i].Drafted != players[j].Drafted {
			return !players[i].Drafted
		}
		return players[i].ADP < players[j].ADP
	})
}
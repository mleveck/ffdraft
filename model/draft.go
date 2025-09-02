package model

type Draft struct {
	Players        []Player
	CurrentPick    int
	CurrentRound   int
	TotalTeams     int
	MyTeam         []Player
	CurrentFilter  Position
	SearchQuery    string
	FilteredPlayers []Player
}

func NewDraft(players []Player, teams int) *Draft {
	return &Draft{
		Players:         players,
		CurrentPick:     1,
		CurrentRound:    1,
		TotalTeams:      teams,
		MyTeam:          []Player{},
		CurrentFilter:   ALL,
		FilteredPlayers: players,
	}
}

func (d *Draft) DraftPlayer(playerIndex int) {
	if playerIndex >= 0 && playerIndex < len(d.Players) {
		d.Players[playerIndex].ToggleDrafted()
		if d.Players[playerIndex].Drafted {
			d.MyTeam = append(d.MyTeam, d.Players[playerIndex])
			d.CurrentPick++
			if d.CurrentPick > d.TotalTeams {
				d.CurrentPick = 1
				d.CurrentRound++
			}
		} else {
			for i, p := range d.MyTeam {
				if p.Name == d.Players[playerIndex].Name {
					d.MyTeam = append(d.MyTeam[:i], d.MyTeam[i+1:]...)
					break
				}
			}
			d.CurrentPick--
			if d.CurrentPick < 1 {
				d.CurrentPick = d.TotalTeams
				d.CurrentRound--
			}
		}
	}
}

func (d *Draft) UpdateFilter(pos Position) {
	d.CurrentFilter = pos
	d.FilteredPlayers = FilterByPosition(d.Players, pos)
}

func (d *Draft) SaveState() DraftState {
	draftedPlayers := make(map[string]bool)
	for _, player := range d.Players {
		if player.Drafted {
			draftedPlayers[player.Name] = true
		}
	}

	return DraftState{
		DraftedPlayers: draftedPlayers,
		CurrentPick:    d.CurrentPick,
		CurrentRound:   d.CurrentRound,
		TotalTeams:     d.TotalTeams,
	}
}

func (d *Draft) LoadState(state DraftState) {
	d.CurrentPick = state.CurrentPick
	d.CurrentRound = state.CurrentRound
	d.TotalTeams = state.TotalTeams

	// Update drafted status for all players
	d.MyTeam = []Player{} // Clear existing team
	for i := range d.Players {
		playerName := d.Players[i].Name
		if state.DraftedPlayers[playerName] {
			d.Players[i].Drafted = true
			d.MyTeam = append(d.MyTeam, d.Players[i])
		} else {
			d.Players[i].Drafted = false
		}
	}
}
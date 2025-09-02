package data

import "ffdraft/model"

// UpdatePlayerData allows easy updating of player data with correct tiers and ADP
// Instructions:
// 1. Go to https://www.borischen.co/p/half-ppr-draft-tiers.html for tier data
// 2. Go to https://www.fantasypros.com/nfl/adp/half-point-ppr-overall.php for ADP data  
// 3. Update the playerData slice below with the correct information

type PlayerData struct {
	Name     string
	Team     string
	Position model.Position
	Tier     int
	ADP      float64
	ByeWeek  int
}

func GetUpdatedPlayers() []model.Player {
	// TODO: Update this data with accurate information from the websites
	playerData := []PlayerData{
		// QBs - Update tiers from Boris Chen and ADP from FantasyPros
		{"Josh Allen", "BUF", model.QB, 1, 3.2, 12},
		{"Lamar Jackson", "BAL", model.QB, 1, 4.1, 14},
		{"Jalen Hurts", "PHI", model.QB, 1, 5.3, 7},
		{"Patrick Mahomes", "KC", model.QB, 2, 6.8, 10},
		{"Dak Prescott", "DAL", model.QB, 2, 8.5, 9},
		{"Tua Tagovailoa", "MIA", model.QB, 3, 9.2, 6},
		{"Joe Burrow", "CIN", model.QB, 3, 10.1, 12},
		{"Aaron Rodgers", "NYJ", model.QB, 3, 10.8, 12},
		
		// RBs - Update tiers from Boris Chen and ADP from FantasyPros  
		{"Christian McCaffrey", "SF", model.RB, 1, 1.2, 9},
		{"Austin Ekeler", "LAC", model.RB, 2, 2.1, 5},
		{"Jonathan Taylor", "IND", model.RB, 2, 2.8, 14},
		{"Derrick Henry", "BAL", model.RB, 2, 3.5, 14},
		{"Saquon Barkley", "PHI", model.RB, 2, 4.2, 7},
		{"Josh Jacobs", "GB", model.RB, 3, 5.1, 10},
		{"Bijan Robinson", "ATL", model.RB, 3, 5.8, 12},
		{"Tony Pollard", "TEN", model.RB, 4, 6.5, 5},
		{"Najee Harris", "PIT", model.RB, 4, 7.2, 9},
		
		// WRs - Update tiers from Boris Chen and ADP from FantasyPros
		{"Tyreek Hill", "MIA", model.WR, 1, 1.5, 6},
		{"Cooper Kupp", "LAR", model.WR, 1, 2.3, 6},
		{"Davante Adams", "NYJ", model.WR, 1, 2.9, 12},
		{"Stefon Diggs", "HOU", model.WR, 2, 3.4, 14},
		{"A.J. Brown", "PHI", model.WR, 2, 4.0, 7},
		{"Ja'Marr Chase", "CIN", model.WR, 2, 4.5, 12},
		{"CeeDee Lamb", "DAL", model.WR, 2, 5.2, 9},
		{"DK Metcalf", "SEA", model.WR, 3, 6.1, 10},
		{"Mike Evans", "TB", model.WR, 3, 6.8, 11},
		{"Keenan Allen", "CHI", model.WR, 3, 7.5, 7},
		{"Amari Cooper", "CLE", model.WR, 4, 8.2, 10},
		
		// TEs - Update tiers from Boris Chen and ADP from FantasyPros
		{"Travis Kelce", "KC", model.TE, 1, 7.2, 10},
		{"Mark Andrews", "BAL", model.TE, 2, 8.1, 14},
		{"T.J. Hockenson", "MIN", model.TE, 2, 9.5, 6},
		{"George Kittle", "SF", model.TE, 2, 10.2, 9},
		{"Darren Waller", "NYG", model.TE, 3, 11.8, 11},
		{"Kyle Pitts", "ATL", model.TE, 3, 12.5, 12},
		{"Evan Engram", "JAX", model.TE, 3, 13.1, 12},
		
		// Kickers - Update with current data
		{"Justin Tucker", "BAL", model.K, 1, 12.1, 14},
		{"Harrison Butker", "KC", model.K, 1, 12.5, 10},
		{"Tyler Bass", "BUF", model.K, 2, 13.2, 12},
		{"Daniel Carlson", "LV", model.K, 2, 13.8, 10},
		{"Younghoe Koo", "ATL", model.K, 2, 14.1, 12},
		
		// DST - Update with current data
		{"San Francisco", "SF", model.DST, 1, 11.5, 9},
		{"Buffalo", "BUF", model.DST, 1, 12.2, 12},
		{"Philadelphia", "PHI", model.DST, 2, 12.8, 7},
		{"Dallas", "DAL", model.DST, 2, 13.3, 9},
		{"Pittsburgh", "PIT", model.DST, 2, 13.9, 9},
	}
	
	var players []model.Player
	for i, data := range playerData {
		players = append(players, model.Player{
			Tier:     data.Tier,
			Rank:     i + 1,
			Name:     data.Name,
			Team:     data.Team,
			Position: data.Position,
			ByeWeek:  data.ByeWeek,
			ADP:      data.ADP,
			Drafted:  false,
		})
	}
	
	// Sort by tier first, then ADP
	model.SortPlayersByTierAndADP(players)
	
	return players
}
package data

import "ffdraft/model"

func GetUpdatedPlayersFromADP() []model.Player {
	players := []model.Player{
		// Top WRs - Tier 1 (Elite WR1s)
		{Tier: 1, Rank: 1, Name: "Ja'Marr Chase", Team: "CIN", Position: model.WR, ByeWeek: 10, ADP: 1.0},
		{Tier: 1, Rank: 5, Name: "CeeDee Lamb", Team: "DAL", Position: model.WR, ByeWeek: 10, ADP: 5.3},
		{Tier: 1, Rank: 6, Name: "Justin Jefferson", Team: "MIN", Position: model.WR, ByeWeek: 6, ADP: 5.7},
		
		// Top RBs - Tier 1 (Elite RB1s)
		{Tier: 1, Rank: 2, Name: "Bijan Robinson", Team: "ATL", Position: model.RB, ByeWeek: 5, ADP: 2.3},
		{Tier: 1, Rank: 3, Name: "Saquon Barkley", Team: "PHI", Position: model.RB, ByeWeek: 9, ADP: 3.0},
		{Tier: 1, Rank: 4, Name: "Jahmyr Gibbs", Team: "DET", Position: model.RB, ByeWeek: 8, ADP: 3.7},
		{Tier: 1, Rank: 7, Name: "Christian McCaffrey", Team: "SF", Position: model.RB, ByeWeek: 14, ADP: 7.7},
		
		// Tier 2 RBs
		{Tier: 2, Rank: 8, Name: "Derrick Henry", Team: "BAL", Position: model.RB, ByeWeek: 7, ADP: 9.0},
		{Tier: 2, Rank: 10, Name: "Ashton Jeanty", Team: "LV", Position: model.RB, ByeWeek: 8, ADP: 9.7},
		{Tier: 2, Rank: 14, Name: "Josh Jacobs", Team: "GB", Position: model.RB, ByeWeek: 5, ADP: 14.7},
		{Tier: 2, Rank: 16, Name: "De'Von Achane", Team: "MIA", Position: model.RB, ByeWeek: 12, ADP: 16.3},
		{Tier: 2, Rank: 18, Name: "Jonathan Taylor", Team: "IND", Position: model.RB, ByeWeek: 11, ADP: 18.0},
		{Tier: 2, Rank: 19, Name: "Bucky Irving", Team: "TB", Position: model.RB, ByeWeek: 9, ADP: 19.3},
		{Tier: 2, Rank: 21, Name: "Chase Brown", Team: "CIN", Position: model.RB, ByeWeek: 10, ADP: 21.3},
		{Tier: 2, Rank: 25, Name: "Kyren Williams", Team: "LAR", Position: model.RB, ByeWeek: 8, ADP: 24.0},
		
		// Tier 2 WRs
		{Tier: 2, Rank: 9, Name: "Amon-Ra St. Brown", Team: "DET", Position: model.WR, ByeWeek: 8, ADP: 9.3},
		{Tier: 2, Rank: 11, Name: "Nico Collins", Team: "HOU", Position: model.WR, ByeWeek: 6, ADP: 11.0},
		{Tier: 2, Rank: 12, Name: "Malik Nabers", Team: "NYG", Position: model.WR, ByeWeek: 14, ADP: 11.0},
		{Tier: 2, Rank: 13, Name: "Puka Nacua", Team: "LAR", Position: model.WR, ByeWeek: 8, ADP: 13.7},
		{Tier: 2, Rank: 15, Name: "Brian Thomas Jr.", Team: "JAX", Position: model.WR, ByeWeek: 8, ADP: 15.7},
		{Tier: 2, Rank: 17, Name: "Drake London", Team: "ATL", Position: model.WR, ByeWeek: 5, ADP: 17.7},
		{Tier: 2, Rank: 23, Name: "A.J. Brown", Team: "PHI", Position: model.WR, ByeWeek: 9, ADP: 22.0},
		{Tier: 2, Rank: 26, Name: "Ladd McConkey", Team: "LAC", Position: model.WR, ByeWeek: 12, ADP: 26.3},
		
		// Tier 1 QBs (Elite QB1s)
		{Tier: 1, Rank: 22, Name: "Josh Allen", Team: "BUF", Position: model.QB, ByeWeek: 7, ADP: 21.3},
		{Tier: 1, Rank: 24, Name: "Lamar Jackson", Team: "BAL", Position: model.QB, ByeWeek: 7, ADP: 22.3},
		{Tier: 1, Rank: 29, Name: "Jayden Daniels", Team: "WAS", Position: model.QB, ByeWeek: 12, ADP: 30.7},
		{Tier: 1, Rank: 35, Name: "Jalen Hurts", Team: "PHI", Position: model.QB, ByeWeek: 9, ADP: 36.7},
		
		// Tier 2 QBs
		{Tier: 2, Rank: 38, Name: "Joe Burrow", Team: "CIN", Position: model.QB, ByeWeek: 10, ADP: 38.7},
		{Tier: 2, Rank: 50, Name: "Patrick Mahomes II", Team: "KC", Position: model.QB, ByeWeek: 10, ADP: 52.3},
		
		// Tier 1 TEs (Elite TE1)
		{Tier: 1, Rank: 20, Name: "Brock Bowers", Team: "LV", Position: model.TE, ByeWeek: 8, ADP: 20.7},
		{Tier: 1, Rank: 27, Name: "Trey McBride", Team: "ARI", Position: model.TE, ByeWeek: 8, ADP: 27.7},
		{Tier: 1, Rank: 34, Name: "George Kittle", Team: "SF", Position: model.TE, ByeWeek: 14, ADP: 35.7},
		
		// Tier 2 TEs
		{Tier: 2, Rank: 53, Name: "Sam LaPorta", Team: "DET", Position: model.TE, ByeWeek: 8, ADP: 54.0},
		{Tier: 2, Rank: 63, Name: "Travis Kelce", Team: "KC", Position: model.TE, ByeWeek: 10, ADP: 62.3},
		{Tier: 2, Rank: 65, Name: "T.J. Hockenson", Team: "MIN", Position: model.TE, ByeWeek: 6, ADP: 64.0},
		
		// Additional Tier 3 players
		{Tier: 3, Rank: 28, Name: "James Cook", Team: "BUF", Position: model.RB, ByeWeek: 7, ADP: 28.7},
		{Tier: 3, Rank: 30, Name: "Omarion Hampton", Team: "LAC", Position: model.RB, ByeWeek: 12, ADP: 31.0},
		{Tier: 3, Rank: 31, Name: "Tyreek Hill", Team: "MIA", Position: model.WR, ByeWeek: 12, ADP: 31.3},
		{Tier: 3, Rank: 32, Name: "Tee Higgins", Team: "CIN", Position: model.WR, ByeWeek: 10, ADP: 31.7},
		{Tier: 3, Rank: 33, Name: "Jaxon Smith-Njigba", Team: "SEA", Position: model.WR, ByeWeek: 8, ADP: 32.3},
		{Tier: 3, Rank: 36, Name: "Breece Hall", Team: "NYJ", Position: model.RB, ByeWeek: 9, ADP: 37.7},
		{Tier: 3, Rank: 37, Name: "Alvin Kamara", Team: "NO", Position: model.RB, ByeWeek: 11, ADP: 38.7},
		{Tier: 3, Rank: 39, Name: "Marvin Harrison Jr.", Team: "ARI", Position: model.WR, ByeWeek: 8, ADP: 39.3},
		{Tier: 3, Rank: 40, Name: "Kenneth Walker III", Team: "SEA", Position: model.RB, ByeWeek: 8, ADP: 39.3},
		{Tier: 3, Rank: 41, Name: "Mike Evans", Team: "TB", Position: model.WR, ByeWeek: 9, ADP: 39.7},
		{Tier: 3, Rank: 42, Name: "Garrett Wilson", Team: "NYJ", Position: model.WR, ByeWeek: 9, ADP: 40.3},
		{Tier: 3, Rank: 43, Name: "Terry McLaurin", Team: "WAS", Position: model.WR, ByeWeek: 12, ADP: 41.3},
		{Tier: 3, Rank: 44, Name: "Chuba Hubbard", Team: "CAR", Position: model.RB, ByeWeek: 14, ADP: 42.7},
		{Tier: 3, Rank: 45, Name: "Davante Adams", Team: "LAR", Position: model.WR, ByeWeek: 8, ADP: 43.7},
		
		// Kickers - Tier 1
		{Tier: 1, Rank: 117, Name: "Brandon Aubrey", Team: "DAL", Position: model.K, ByeWeek: 10, ADP: 119.7},
		{Tier: 1, Rank: 181, Name: "Ka'imi Fairbairn", Team: "HOU", Position: model.K, ByeWeek: 6, ADP: 181.0},
		{Tier: 1, Rank: 184, Name: "Harrison Butker", Team: "KC", Position: model.K, ByeWeek: 10, ADP: 185.7},
		{Tier: 1, Rank: 185, Name: "Chase McLaughlin", Team: "TB", Position: model.K, ByeWeek: 9, ADP: 188.0},
		
		// Tier 2 Kickers
		{Tier: 2, Rank: 201, Name: "Tyler Bass", Team: "BUF", Position: model.K, ByeWeek: 7, ADP: 209.0},
		{Tier: 2, Rank: 205, Name: "Evan McPherson", Team: "CIN", Position: model.K, ByeWeek: 10, ADP: 213.7},
		{Tier: 2, Rank: 206, Name: "Jake Elliott", Team: "PHI", Position: model.K, ByeWeek: 9, ADP: 215.0},
		{Tier: 2, Rank: 208, Name: "Younghoe Koo", Team: "ATL", Position: model.K, ByeWeek: 5, ADP: 217.0},
		
		// DST - Tier 1 
		{Tier: 1, Rank: 122, Name: "Denver Broncos", Team: "DEN", Position: model.DST, ByeWeek: 12, ADP: 123.3},
		{Tier: 1, Rank: 127, Name: "Philadelphia Eagles", Team: "PHI", Position: model.DST, ByeWeek: 9, ADP: 128.3},
		{Tier: 1, Rank: 137, Name: "Pittsburgh Steelers", Team: "PIT", Position: model.DST, ByeWeek: 5, ADP: 144.3},
		
		// Tier 2 DST
		{Tier: 2, Rank: 147, Name: "Baltimore Ravens", Team: "BAL", Position: model.DST, ByeWeek: 7, ADP: 155.0},
		{Tier: 2, Rank: 158, Name: "Minnesota Vikings", Team: "MIN", Position: model.DST, ByeWeek: 6, ADP: 161.0},
		{Tier: 2, Rank: 171, Name: "Houston Texans", Team: "HOU", Position: model.DST, ByeWeek: 6, ADP: 175.3},
		{Tier: 2, Rank: 178, Name: "Detroit Lions", Team: "DET", Position: model.DST, ByeWeek: 8, ADP: 180.7},
		{Tier: 2, Rank: 186, Name: "Buffalo Bills", Team: "BUF", Position: model.DST, ByeWeek: 7, ADP: 189.3},
	}
	
	// Sort by tier first, then ADP
	model.SortPlayersByTierAndADP(players)
	
	return players
}
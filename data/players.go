package data

import "ffdraft/model"

func GetSamplePlayers() []model.Player {
	// Use corrected Boris Chen tiers with exact assignments from tiers.txt
	// Brandon Aubrey is now correctly tier 22 (rank 197), Mason Taylor tier 23 (rank 194)
	// Non-Boris Chen players start at tier 24+
	return GetCorrectedFinalPlayers()
}

func GetSamplePlayersOld() []model.Player {
	players := []model.Player{
		// QBs - Tier 1 (Elite)
		{Tier: 1, Rank: 1, Name: "Josh Allen", Team: "BUF", Position: model.QB, ByeWeek: 12, ADP: 3.2},
		{Tier: 1, Rank: 2, Name: "Lamar Jackson", Team: "BAL", Position: model.QB, ByeWeek: 14, ADP: 4.1},
		{Tier: 1, Rank: 3, Name: "Jalen Hurts", Team: "PHI", Position: model.QB, ByeWeek: 7, ADP: 5.3},
		// QBs - Tier 2 (High-End QB1)
		{Tier: 2, Rank: 4, Name: "Patrick Mahomes", Team: "KC", Position: model.QB, ByeWeek: 10, ADP: 6.8},
		{Tier: 2, Rank: 5, Name: "Dak Prescott", Team: "DAL", Position: model.QB, ByeWeek: 9, ADP: 8.5},

		// RBs - Tier 1 (Elite, CMC tier)
		{Tier: 1, Rank: 6, Name: "Christian McCaffrey", Team: "SF", Position: model.RB, ByeWeek: 9, ADP: 1.2},
		// RBs - Tier 2 (High-End RB1)
		{Tier: 2, Rank: 7, Name: "Austin Ekeler", Team: "LAC", Position: model.RB, ByeWeek: 5, ADP: 2.1},
		{Tier: 2, Rank: 8, Name: "Jonathan Taylor", Team: "IND", Position: model.RB, ByeWeek: 14, ADP: 2.8},
		{Tier: 2, Rank: 9, Name: "Derrick Henry", Team: "BAL", Position: model.RB, ByeWeek: 14, ADP: 3.5},
		{Tier: 2, Rank: 10, Name: "Saquon Barkley", Team: "PHI", Position: model.RB, ByeWeek: 7, ADP: 4.2},
		// RBs - Tier 3 (Mid RB1)
		{Tier: 3, Rank: 11, Name: "Josh Jacobs", Team: "GB", Position: model.RB, ByeWeek: 10, ADP: 5.1},
		{Tier: 3, Rank: 12, Name: "Bijan Robinson", Team: "ATL", Position: model.RB, ByeWeek: 12, ADP: 5.8},

		// WRs - Tier 1 (Elite WR1)
		{Tier: 1, Rank: 13, Name: "Tyreek Hill", Team: "MIA", Position: model.WR, ByeWeek: 6, ADP: 1.5},
		{Tier: 1, Rank: 14, Name: "Cooper Kupp", Team: "LAR", Position: model.WR, ByeWeek: 6, ADP: 2.3},
		{Tier: 1, Rank: 15, Name: "Davante Adams", Team: "NYJ", Position: model.WR, ByeWeek: 12, ADP: 2.9},
		// WRs - Tier 2 (High-End WR1)
		{Tier: 2, Rank: 16, Name: "Stefon Diggs", Team: "HOU", Position: model.WR, ByeWeek: 14, ADP: 3.4},
		{Tier: 2, Rank: 17, Name: "A.J. Brown", Team: "PHI", Position: model.WR, ByeWeek: 7, ADP: 4.0},
		{Tier: 2, Rank: 18, Name: "Ja'Marr Chase", Team: "CIN", Position: model.WR, ByeWeek: 12, ADP: 4.5},
		{Tier: 2, Rank: 19, Name: "CeeDee Lamb", Team: "DAL", Position: model.WR, ByeWeek: 9, ADP: 5.2},
		// WRs - Tier 3 (Mid WR1)
		{Tier: 3, Rank: 20, Name: "DK Metcalf", Team: "SEA", Position: model.WR, ByeWeek: 10, ADP: 6.1},

		// TEs - Tier 1 (Elite TE1)
		{Tier: 1, Rank: 21, Name: "Travis Kelce", Team: "KC", Position: model.TE, ByeWeek: 10, ADP: 7.2},
		// TEs - Tier 2 (High-End TE1)
		{Tier: 2, Rank: 22, Name: "Mark Andrews", Team: "BAL", Position: model.TE, ByeWeek: 14, ADP: 8.1},
		{Tier: 2, Rank: 23, Name: "T.J. Hockenson", Team: "MIN", Position: model.TE, ByeWeek: 6, ADP: 9.5},
		{Tier: 2, Rank: 24, Name: "George Kittle", Team: "SF", Position: model.TE, ByeWeek: 9, ADP: 10.2},
		// TEs - Tier 3 (Mid TE1)
		{Tier: 3, Rank: 25, Name: "Darren Waller", Team: "NYG", Position: model.TE, ByeWeek: 11, ADP: 11.8},

		// Kickers - Tier 1 (Elite K)
		{Tier: 1, Rank: 26, Name: "Justin Tucker", Team: "BAL", Position: model.K, ByeWeek: 14, ADP: 12.1},
		{Tier: 1, Rank: 27, Name: "Harrison Butker", Team: "KC", Position: model.K, ByeWeek: 10, ADP: 12.5},
		// Kickers - Tier 2 (Good K)
		{Tier: 2, Rank: 28, Name: "Tyler Bass", Team: "BUF", Position: model.K, ByeWeek: 12, ADP: 13.2},
		{Tier: 2, Rank: 29, Name: "Daniel Carlson", Team: "LV", Position: model.K, ByeWeek: 10, ADP: 13.8},
		{Tier: 2, Rank: 30, Name: "Younghoe Koo", Team: "ATL", Position: model.K, ByeWeek: 12, ADP: 14.1},

		// DST - Tier 1 (Elite DST)
		{Tier: 1, Rank: 31, Name: "San Francisco", Team: "SF", Position: model.DST, ByeWeek: 9, ADP: 11.5},
		{Tier: 1, Rank: 32, Name: "Buffalo", Team: "BUF", Position: model.DST, ByeWeek: 12, ADP: 12.2},
		// DST - Tier 2 (Good DST)
		{Tier: 2, Rank: 33, Name: "Philadelphia", Team: "PHI", Position: model.DST, ByeWeek: 7, ADP: 12.8},
		{Tier: 2, Rank: 34, Name: "Dallas", Team: "DAL", Position: model.DST, ByeWeek: 9, ADP: 13.3},
		{Tier: 2, Rank: 35, Name: "Pittsburgh", Team: "PIT", Position: model.DST, ByeWeek: 9, ADP: 13.9},

		// More players for depth
		// QBs - Tier 3 (Mid QB1)
		{Tier: 3, Rank: 36, Name: "Tua Tagovailoa", Team: "MIA", Position: model.QB, ByeWeek: 6, ADP: 9.2},
		{Tier: 3, Rank: 37, Name: "Joe Burrow", Team: "CIN", Position: model.QB, ByeWeek: 12, ADP: 10.1},
		{Tier: 3, Rank: 38, Name: "Aaron Rodgers", Team: "NYJ", Position: model.QB, ByeWeek: 12, ADP: 10.8},
		// RBs - Tier 4 (Low RB1/High RB2)
		{Tier: 4, Rank: 39, Name: "Tony Pollard", Team: "TEN", Position: model.RB, ByeWeek: 5, ADP: 6.5},
		{Tier: 4, Rank: 40, Name: "Najee Harris", Team: "PIT", Position: model.RB, ByeWeek: 9, ADP: 7.2},
		// WRs - Tier 3 (continued)
		{Tier: 3, Rank: 41, Name: "Mike Evans", Team: "TB", Position: model.WR, ByeWeek: 11, ADP: 6.8},
		{Tier: 3, Rank: 42, Name: "Keenan Allen", Team: "CHI", Position: model.WR, ByeWeek: 7, ADP: 7.5},
		// WRs - Tier 4 (Low WR1/High WR2)
		{Tier: 4, Rank: 43, Name: "Amari Cooper", Team: "CLE", Position: model.WR, ByeWeek: 10, ADP: 8.2},
		// TEs - Tier 3 (continued)
		{Tier: 3, Rank: 44, Name: "Kyle Pitts", Team: "ATL", Position: model.TE, ByeWeek: 12, ADP: 12.5},
		{Tier: 3, Rank: 45, Name: "Evan Engram", Team: "JAX", Position: model.TE, ByeWeek: 12, ADP: 13.1},
	}

	// Sort players by tier first, then by ADP
	model.SortPlayersByTierAndADP(players)
	return players
}
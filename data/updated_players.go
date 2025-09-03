package data

import "ffdraft/model"

func GetUpdatedPlayers() []model.Player {
	players := []model.Player{
		// Tier 1 (6 players) - Boris Chen Rank 1-6
		{Tier: 1, Rank: 1, Name: "Ja'Marr Chase", Team: "CIN", Position: model.WR, ByeWeek: 10, ADP: 1.0},
		{Tier: 1, Rank: 2, Name: "Bijan Robinson", Team: "ATL", Position: model.RB, ByeWeek: 5, ADP: 2.3},
		{Tier: 1, Rank: 3, Name: "Saquon Barkley", Team: "PHI", Position: model.RB, ByeWeek: 9, ADP: 3.0},
		{Tier: 1, Rank: 4, Name: "CeeDee Lamb", Team: "DAL", Position: model.WR, ByeWeek: 10, ADP: 5.3},
		{Tier: 1, Rank: 5, Name: "Justin Jefferson", Team: "MIN", Position: model.WR, ByeWeek: 6, ADP: 5.7},
		{Tier: 1, Rank: 6, Name: "Jahmyr Gibbs", Team: "DET", Position: model.RB, ByeWeek: 8, ADP: 3.7},

		// Tier 2 (4 players) - Boris Chen Rank 7-10  
		{Tier: 2, Rank: 7, Name: "Derrick Henry", Team: "BAL", Position: model.RB, ByeWeek: 7, ADP: 9.0},
		{Tier: 2, Rank: 8, Name: "Nico Collins", Team: "HOU", Position: model.WR, ByeWeek: 6, ADP: 11.0},
		{Tier: 2, Rank: 9, Name: "Brian Thomas Jr.", Team: "JAC", Position: model.WR, ByeWeek: 8, ADP: 15.7},
		{Tier: 2, Rank: 10, Name: "Malik Nabers", Team: "NYG", Position: model.WR, ByeWeek: 14, ADP: 11.0},

		// Tier 2 continued (1 player) - Boris Chen Rank 11
		{Tier: 2, Rank: 11, Name: "Puka Nacua", Team: "LAR", Position: model.WR, ByeWeek: 8, ADP: 13.7},

		// Tier 3 (6 players) - Boris Chen Rank 12-17
		{Tier: 3, Rank: 12, Name: "Ashton Jeanty", Team: "LV", Position: model.RB, ByeWeek: 8, ADP: 9.7},
		{Tier: 3, Rank: 13, Name: "Christian McCaffrey", Team: "SF", Position: model.RB, ByeWeek: 14, ADP: 7.7},
		{Tier: 3, Rank: 14, Name: "Drake London", Team: "ATL", Position: model.WR, ByeWeek: 5, ADP: 17.7},
		{Tier: 3, Rank: 15, Name: "Jonathan Taylor", Team: "IND", Position: model.RB, ByeWeek: 11, ADP: 18.0},
		{Tier: 3, Rank: 16, Name: "A.J. Brown", Team: "PHI", Position: model.WR, ByeWeek: 9, ADP: 22.0},
		{Tier: 3, Rank: 17, Name: "Amon-Ra St. Brown", Team: "DET", Position: model.WR, ByeWeek: 8, ADP: 9.3},

		// Tier 4 (7 players) - Boris Chen Rank 18-24
		{Tier: 4, Rank: 18, Name: "Josh Jacobs", Team: "GB", Position: model.RB, ByeWeek: 5, ADP: 14.7},
		{Tier: 4, Rank: 19, Name: "George Kittle", Team: "SF", Position: model.TE, ByeWeek: 14, ADP: 35.7},
		{Tier: 4, Rank: 20, Name: "Brock Bowers", Team: "LV", Position: model.TE, ByeWeek: 8, ADP: 20.7},
		{Tier: 4, Rank: 21, Name: "Bucky Irving", Team: "TB", Position: model.RB, ByeWeek: 9, ADP: 19.3},
		{Tier: 4, Rank: 22, Name: "Ladd McConkey", Team: "LAC", Position: model.WR, ByeWeek: 12, ADP: 26.3},
		{Tier: 4, Rank: 23, Name: "Kyren Williams", Team: "LAR", Position: model.RB, ByeWeek: 8, ADP: 24.0},

		// Tier 5 (5 players) - Boris Chen Rank 24-28
		{Tier: 5, Rank: 24, Name: "De'Von Achane", Team: "MIA", Position: model.RB, ByeWeek: 12, ADP: 16.3},
		{Tier: 5, Rank: 25, Name: "Josh Allen", Team: "BUF", Position: model.QB, ByeWeek: 7, ADP: 21.3},
		{Tier: 5, Rank: 26, Name: "Lamar Jackson", Team: "BAL", Position: model.QB, ByeWeek: 7, ADP: 22.3},
		{Tier: 5, Rank: 27, Name: "Chase Brown", Team: "CIN", Position: model.RB, ByeWeek: 10, ADP: 21.3},
		{Tier: 5, Rank: 28, Name: "Tee Higgins", Team: "CIN", Position: model.WR, ByeWeek: 10, ADP: 31.7},

		// Tier 5 continued (1 player) - Boris Chen Rank 29
		{Tier: 5, Rank: 29, Name: "Mike Evans", Team: "TB", Position: model.WR, ByeWeek: 9, ADP: 39.7},

		// Tier 6 (4 players) - Boris Chen Rank 30-33
		{Tier: 6, Rank: 30, Name: "Trey McBride", Team: "ARI", Position: model.TE, ByeWeek: 8, ADP: 27.7},
		{Tier: 6, Rank: 31, Name: "Jayden Daniels", Team: "WAS", Position: model.QB, ByeWeek: 12, ADP: 30.7},
		{Tier: 6, Rank: 32, Name: "James Cook", Team: "BUF", Position: model.RB, ByeWeek: 7, ADP: 28.7},
		{Tier: 6, Rank: 33, Name: "Jaxon Smith-Njigba", Team: "SEA", Position: model.WR, ByeWeek: 8, ADP: 32.3},

		// Tier 7 (4 players) - Boris Chen Rank 34-37
		{Tier: 7, Rank: 34, Name: "Tyreek Hill", Team: "MIA", Position: model.WR, ByeWeek: 12, ADP: 31.3},
		{Tier: 7, Rank: 35, Name: "Davante Adams", Team: "LAR", Position: model.WR, ByeWeek: 8, ADP: 43.7},
		{Tier: 7, Rank: 36, Name: "Jalen Hurts", Team: "PHI", Position: model.QB, ByeWeek: 9, ADP: 36.7},
		{Tier: 7, Rank: 37, Name: "Omarion Hampton", Team: "LAC", Position: model.RB, ByeWeek: 12, ADP: 31.0},

		// Tier 8 (4 players) - Boris Chen Rank 38-41
		{Tier: 8, Rank: 38, Name: "Marvin Harrison Jr.", Team: "ARI", Position: model.WR, ByeWeek: 8, ADP: 39.3},
		{Tier: 8, Rank: 39, Name: "Kenneth Walker III", Team: "SEA", Position: model.RB, ByeWeek: 8, ADP: 39.3},
		{Tier: 8, Rank: 40, Name: "Garrett Wilson", Team: "NYJ", Position: model.WR, ByeWeek: 9, ADP: 40.3},
		{Tier: 8, Rank: 41, Name: "Terry McLaurin", Team: "WAS", Position: model.WR, ByeWeek: 12, ADP: 41.3},

		// Tier 9 (10 players) - Boris Chen Rank 42-51
		{Tier: 9, Rank: 42, Name: "DK Metcalf", Team: "SEA", Position: model.WR, ByeWeek: 8, ADP: 49.0},
		{Tier: 9, Rank: 43, Name: "Chuba Hubbard", Team: "CAR", Position: model.RB, ByeWeek: 14, ADP: 42.7},
		{Tier: 9, Rank: 44, Name: "Breece Hall", Team: "NYJ", Position: model.RB, ByeWeek: 9, ADP: 37.7},
		{Tier: 9, Rank: 45, Name: "Joe Burrow", Team: "CIN", Position: model.QB, ByeWeek: 10, ADP: 38.7},
		{Tier: 9, Rank: 46, Name: "James Conner", Team: "ARI", Position: model.RB, ByeWeek: 8, ADP: 47.0},
		{Tier: 9, Rank: 47, Name: "Courtland Sutton", Team: "DEN", Position: model.WR, ByeWeek: 12, ADP: 52.7},
		{Tier: 9, Rank: 48, Name: "TreVeyon Henderson", Team: "OSU", Position: model.RB, ByeWeek: 14, ADP: 48.0},
		{Tier: 9, Rank: 49, Name: "Jameson Williams", Team: "DET", Position: model.WR, ByeWeek: 8, ADP: 62.0},
		{Tier: 9, Rank: 50, Name: "Alvin Kamara", Team: "NO", Position: model.RB, ByeWeek: 11, ADP: 38.7},

		// Tier 10 (6 players) - Boris Chen Rank 51-56
		{Tier: 10, Rank: 51, Name: "DJ Moore", Team: "CHI", Position: model.WR, ByeWeek: 5, ADP: 51.7},
		{Tier: 10, Rank: 52, Name: "Xavier Worthy", Team: "KC", Position: model.WR, ByeWeek: 10, ADP: 57.7},
		{Tier: 10, Rank: 53, Name: "Calvin Ridley", Team: "TEN", Position: model.WR, ByeWeek: 10, ADP: 67.7},
		{Tier: 10, Rank: 54, Name: "Tetairoa McMillan", Team: "ARI", Position: model.WR, ByeWeek: 8, ADP: 61.3},
		{Tier: 10, Rank: 55, Name: "George Pickens", Team: "PIT", Position: model.WR, ByeWeek: 5, ADP: 60.3},

		// Tier 11 (13 players) - Boris Chen Rank 56-68
		{Tier: 11, Rank: 56, Name: "DeVonta Smith", Team: "PHI", Position: model.WR, ByeWeek: 9, ADP: 55.3},
		{Tier: 11, Rank: 57, Name: "David Montgomery", Team: "DET", Position: model.RB, ByeWeek: 8, ADP: 54.7},
		{Tier: 11, Rank: 58, Name: "D'Andre Swift", Team: "CHI", Position: model.RB, ByeWeek: 5, ADP: 59.7},
		{Tier: 11, Rank: 59, Name: "Zay Flowers", Team: "BAL", Position: model.WR, ByeWeek: 7, ADP: 63.0},
		{Tier: 11, Rank: 60, Name: "Tony Pollard", Team: "TEN", Position: model.RB, ByeWeek: 10, ADP: 58.3},
		{Tier: 11, Rank: 61, Name: "Sam LaPorta", Team: "DET", Position: model.TE, ByeWeek: 8, ADP: 54.0},
		{Tier: 11, Rank: 62, Name: "RJ Harvey", Team: "UCF", Position: model.RB, ByeWeek: 12, ADP: 52.7},
		{Tier: 11, Rank: 63, Name: "Isiah Pacheco", Team: "KC", Position: model.RB, ByeWeek: 10, ADP: 58.7},
		{Tier: 11, Rank: 64, Name: "Patrick Mahomes II", Team: "KC", Position: model.QB, ByeWeek: 10, ADP: 52.3},
		{Tier: 11, Rank: 65, Name: "Jaylen Waddle", Team: "MIA", Position: model.WR, ByeWeek: 12, ADP: 70.0},
		{Tier: 11, Rank: 66, Name: "Aaron Jones Sr.", Team: "MIN", Position: model.RB, ByeWeek: 6, ADP: 66.3},
		{Tier: 11, Rank: 67, Name: "Baker Mayfield", Team: "TB", Position: model.QB, ByeWeek: 9, ADP: 68.3},
		{Tier: 11, Rank: 68, Name: "Mark Andrews", Team: "BAL", Position: model.TE, ByeWeek: 7, ADP: 75.7},

		// Tier 12 (17 players) - Boris Chen Rank 69-85
		{Tier: 12, Rank: 69, Name: "Bo Nix", Team: "DEN", Position: model.QB, ByeWeek: 12, ADP: 71.7},
		{Tier: 12, Rank: 70, Name: "Kaleb Johnson", Team: "IOWA", Position: model.RB, ByeWeek: 5, ADP: 74.0},
		{Tier: 12, Rank: 71, Name: "T.J. Hockenson", Team: "MIN", Position: model.TE, ByeWeek: 6, ADP: 64.0},
		{Tier: 12, Rank: 72, Name: "Tyrone Tracy Jr.", Team: "NYG", Position: model.RB, ByeWeek: 14, ADP: 76.0},
		{Tier: 12, Rank: 73, Name: "Travis Hunter", Team: "COL", Position: model.WR, ByeWeek: 8, ADP: 72.0},
		{Tier: 12, Rank: 74, Name: "Ricky Pearsall", Team: "SF", Position: model.WR, ByeWeek: 14, ADP: 83.7},
		{Tier: 12, Rank: 75, Name: "Rome Odunze", Team: "CHI", Position: model.WR, ByeWeek: 5, ADP: 80.0},
		{Tier: 12, Rank: 76, Name: "Kyler Murray", Team: "ARI", Position: model.QB, ByeWeek: 8, ADP: 93.7},
		{Tier: 12, Rank: 77, Name: "Chris Olave", Team: "NO", Position: model.WR, ByeWeek: 11, ADP: 77.7},
		{Tier: 12, Rank: 78, Name: "Dak Prescott", Team: "DAL", Position: model.QB, ByeWeek: 10, ADP: 95.7},
		{Tier: 12, Rank: 79, Name: "Rashee Rice", Team: "KC", Position: model.WR, ByeWeek: 10, ADP: 64.7},
		{Tier: 12, Rank: 80, Name: "Jordan Mason", Team: "SF", Position: model.RB, ByeWeek: 14, ADP: 89.3},
		{Tier: 12, Rank: 81, Name: "Travis Kelce", Team: "KC", Position: model.TE, ByeWeek: 10, ADP: 62.3},
		{Tier: 12, Rank: 82, Name: "Emeka Egbuka", Team: "OSU", Position: model.WR, ByeWeek: 9, ADP: 86.7},
		{Tier: 12, Rank: 83, Name: "Jerry Jeudy", Team: "CLE", Position: model.WR, ByeWeek: 9, ADP: 77.7},
		{Tier: 12, Rank: 84, Name: "Jaylen Warren", Team: "PIT", Position: model.RB, ByeWeek: 5, ADP: 83.3},
		{Tier: 12, Rank: 85, Name: "Deebo Samuel Sr.", Team: "SF", Position: model.WR, ByeWeek: 14, ADP: 87.7},

		// Tier 13 (17 players) - Boris Chen Rank 86-102
		{Tier: 13, Rank: 86, Name: "Jordan Addison", Team: "MIN", Position: model.WR, ByeWeek: 6, ADP: 92.3},
		{Tier: 13, Rank: 87, Name: "Brock Purdy", Team: "SF", Position: model.QB, ByeWeek: 14, ADP: 97.3},
		{Tier: 13, Rank: 88, Name: "David Njoku", Team: "CLE", Position: model.TE, ByeWeek: 9, ADP: 90.0},
		{Tier: 13, Rank: 89, Name: "Tucker Kraft", Team: "GB", Position: model.TE, ByeWeek: 5, ADP: 99.7},
		{Tier: 13, Rank: 90, Name: "Justin Fields", Team: "PIT", Position: model.QB, ByeWeek: 5, ADP: 113.0},
		{Tier: 13, Rank: 91, Name: "Travis Etienne Jr.", Team: "JAC", Position: model.RB, ByeWeek: 8, ADP: 93.0},
		{Tier: 13, Rank: 92, Name: "Zach Charbonnet", Team: "SEA", Position: model.RB, ByeWeek: 8, ADP: 99.7},
		{Tier: 13, Rank: 93, Name: "Jacory Croskey-Merritt", Team: "TCU", Position: model.RB, ByeWeek: 12, ADP: 109.3},
		{Tier: 13, Rank: 94, Name: "Jayden Reed", Team: "GB", Position: model.WR, ByeWeek: 5, ADP: 111.7},
		{Tier: 13, Rank: 95, Name: "Justin Herbert", Team: "LAC", Position: model.QB, ByeWeek: 12, ADP: 110.7},
		{Tier: 13, Rank: 96, Name: "Jauan Jennings", Team: "SF", Position: model.WR, ByeWeek: 14, ADP: 103.3},
		{Tier: 13, Rank: 97, Name: "Tyler Warren", Team: "PSU", Position: model.TE, ByeWeek: 11, ADP: 98.7},
		{Tier: 13, Rank: 98, Name: "Stefon Diggs", Team: "HOU", Position: model.WR, ByeWeek: 6, ADP: 91.3},
		{Tier: 13, Rank: 99, Name: "Tank Bigsby", Team: "JAC", Position: model.RB, ByeWeek: 8, ADP: 110.3},
		{Tier: 13, Rank: 100, Name: "Matthew Golden", Team: "TEX", Position: model.WR, ByeWeek: 5, ADP: 97.0},
		{Tier: 13, Rank: 101, Name: "Drake Maye", Team: "NE", Position: model.QB, ByeWeek: 14, ADP: 120.0},
		{Tier: 13, Rank: 102, Name: "Evan Engram", Team: "JAC", Position: model.TE, ByeWeek: 8, ADP: 84.3},

		// Tier 13 continued (5 players) - Boris Chen Rank 103-107
		{Tier: 13, Rank: 103, Name: "Javonte Williams", Team: "DEN", Position: model.RB, ByeWeek: 12, ADP: 112.3},
		{Tier: 13, Rank: 104, Name: "Caleb Williams", Team: "CHI", Position: model.QB, ByeWeek: 5, ADP: 113.7},
		{Tier: 13, Rank: 105, Name: "Jakobi Meyers", Team: "LV", Position: model.WR, ByeWeek: 8, ADP: 92.3},
		{Tier: 13, Rank: 106, Name: "Jared Goff", Team: "DET", Position: model.QB, ByeWeek: 8, ADP: 102.3},
		{Tier: 13, Rank: 107, Name: "J.K. Dobbins", Team: "LAC", Position: model.RB, ByeWeek: 12, ADP: 99.0},

		// Tier 14 (25 players) - Boris Chen Rank 108-132
		{Tier: 14, Rank: 108, Name: "Rhamondre Stevenson", Team: "NE", Position: model.RB, ByeWeek: 14, ADP: 117.3},
		{Tier: 14, Rank: 109, Name: "Jordan Love", Team: "GB", Position: model.QB, ByeWeek: 5, ADP: 138.7},
		{Tier: 14, Rank: 110, Name: "Khalil Shakir", Team: "BUF", Position: model.WR, ByeWeek: 7, ADP: 110.0},
		{Tier: 14, Rank: 111, Name: "Trevor Lawrence", Team: "JAC", Position: model.QB, ByeWeek: 8, ADP: 157.3},
		{Tier: 14, Rank: 112, Name: "Keon Coleman", Team: "BUF", Position: model.WR, ByeWeek: 7, ADP: 131.3},
		{Tier: 14, Rank: 113, Name: "Cam Skattebo", Team: "ASU", Position: model.RB, ByeWeek: 14, ADP: 107.7},
		{Tier: 14, Rank: 114, Name: "C.J. Stroud", Team: "HOU", Position: model.QB, ByeWeek: 6, ADP: 128.7},
		{Tier: 14, Rank: 115, Name: "Kyle Pitts Sr.", Team: "ATL", Position: model.TE, ByeWeek: 5, ADP: 133.7},
		{Tier: 14, Rank: 116, Name: "Josh Downs", Team: "IND", Position: model.WR, ByeWeek: 11, ADP: 123.3},
		{Tier: 14, Rank: 117, Name: "J.J. McCarthy", Team: "MIN", Position: model.QB, ByeWeek: 6, ADP: 139.7},
		{Tier: 14, Rank: 118, Name: "Dallas Goedert", Team: "PHI", Position: model.TE, ByeWeek: 9, ADP: 141.7},
		{Tier: 14, Rank: 119, Name: "Chris Godwin Jr.", Team: "TB", Position: model.WR, ByeWeek: 9, ADP: 103.0},
		{Tier: 14, Rank: 120, Name: "Cooper Kupp", Team: "LAR", Position: model.WR, ByeWeek: 8, ADP: 99.7},
		{Tier: 14, Rank: 121, Name: "Dalton Kincaid", Team: "BUF", Position: model.TE, ByeWeek: 7, ADP: 126.3},
		{Tier: 14, Rank: 122, Name: "Darnell Mooney", Team: "ATL", Position: model.WR, ByeWeek: 5, ADP: 157.7},
		{Tier: 14, Rank: 123, Name: "Najee Harris", Team: "PIT", Position: model.RB, ByeWeek: 5, ADP: 124.3},
		{Tier: 14, Rank: 124, Name: "Austin Ekeler", Team: "WAS", Position: model.RB, ByeWeek: 12, ADP: 120.7},
		{Tier: 14, Rank: 125, Name: "Braelon Allen", Team: "NYJ", Position: model.RB, ByeWeek: 9, ADP: 115.0},
		{Tier: 14, Rank: 126, Name: "Michael Pittman Jr.", Team: "IND", Position: model.WR, ByeWeek: 11, ADP: 122.7},
		{Tier: 14, Rank: 127, Name: "Rashid Shaheed", Team: "NO", Position: model.WR, ByeWeek: 11, ADP: 149.7},
		{Tier: 14, Rank: 128, Name: "Trey Benson", Team: "ARI", Position: model.RB, ByeWeek: 8, ADP: 149.7},
		{Tier: 14, Rank: 129, Name: "Brian Robinson Jr.", Team: "WAS", Position: model.RB, ByeWeek: 12, ADP: 95.3},
		{Tier: 14, Rank: 130, Name: "Colston Loveland", Team: "MICH", Position: model.TE, ByeWeek: 5, ADP: 117.7},
		{Tier: 14, Rank: 131, Name: "Ray Davis", Team: "BUF", Position: model.RB, ByeWeek: 7, ADP: 159.0},
		{Tier: 14, Rank: 132, Name: "Quinshon Judkins", Team: "OSU", Position: model.RB, ByeWeek: 9, ADP: 103.0},

		// Tier 15 (10 players) - Boris Chen Rank 133-142
		{Tier: 15, Rank: 133, Name: "Joe Mixon", Team: "HOU", Position: model.RB, ByeWeek: 6, ADP: 82.7},
		{Tier: 15, Rank: 134, Name: "Nick Chubb", Team: "CLE", Position: model.RB, ByeWeek: 9, ADP: 122.3},
		{Tier: 15, Rank: 135, Name: "Brandon Aiyuk", Team: "PIT", Position: model.WR, ByeWeek: 5, ADP: 147.0},
		{Tier: 15, Rank: 136, Name: "Rachaad White", Team: "TB", Position: model.RB, ByeWeek: 9, ADP: 152.3},
		{Tier: 15, Rank: 137, Name: "Tua Tagovailoa", Team: "MIA", Position: model.QB, ByeWeek: 12, ADP: 166.0},
		{Tier: 15, Rank: 138, Name: "Jake Ferguson", Team: "DAL", Position: model.TE, ByeWeek: 10, ADP: 127.0},
		{Tier: 15, Rank: 139, Name: "Tyler Allgeier", Team: "ATL", Position: model.RB, ByeWeek: 5, ADP: 169.7},
		{Tier: 15, Rank: 140, Name: "Rashod Bateman", Team: "BAL", Position: model.WR, ByeWeek: 7, ADP: 178.0},
		{Tier: 15, Rank: 141, Name: "Christian Kirk", Team: "JAC", Position: model.WR, ByeWeek: 8, ADP: 165.7},

		// Tier 16 (11 players) - Boris Chen Rank 142-152
		{Tier: 16, Rank: 142, Name: "Jaydon Blue", Team: "TEX", Position: model.RB, ByeWeek: 10, ADP: 135.0},
		{Tier: 16, Rank: 143, Name: "Hunter Henry", Team: "NE", Position: model.TE, ByeWeek: 14, ADP: 160.0},
		{Tier: 16, Rank: 144, Name: "Bryce Young", Team: "CAR", Position: model.QB, ByeWeek: 14, ADP: 176.0},
		{Tier: 16, Rank: 145, Name: "Tyjae Spears", Team: "TEN", Position: model.RB, ByeWeek: 10, ADP: 156.7},
		{Tier: 16, Rank: 146, Name: "Marvin Mims Jr.", Team: "DEN", Position: model.WR, ByeWeek: 12, ADP: 156.0},
		{Tier: 16, Rank: 147, Name: "Michael Penix Jr.", Team: "ATL", Position: model.QB, ByeWeek: 5, ADP: 163.7},
		{Tier: 16, Rank: 148, Name: "Bhayshul Tuten", Team: "VT", Position: model.RB, ByeWeek: 8, ADP: 146.7},
		{Tier: 16, Rank: 149, Name: "Jerome Ford", Team: "CLE", Position: model.RB, ByeWeek: 9, ADP: 154.3},
		{Tier: 16, Rank: 150, Name: "Dylan Sampson", Team: "TENN", Position: model.RB, ByeWeek: 9, ADP: 160.3},
		{Tier: 16, Rank: 151, Name: "Luther Burden III", Team: "MIZZ", Position: model.WR, ByeWeek: 5, ADP: 156.7},

		// Tier 17 (10 players) - Boris Chen Rank 152-161
		{Tier: 17, Rank: 152, Name: "Cedric Tillman", Team: "CLE", Position: model.WR, ByeWeek: 9, ADP: 181.0},
		{Tier: 17, Rank: 153, Name: "Jonnu Smith", Team: "MIA", Position: model.TE, ByeWeek: 12, ADP: 140.0},
		{Tier: 17, Rank: 154, Name: "Geno Smith", Team: "SEA", Position: model.QB, ByeWeek: 8, ADP: 180.7},
		{Tier: 17, Rank: 155, Name: "Matthew Stafford", Team: "LAR", Position: model.QB, ByeWeek: 8, ADP: 175.0},
		{Tier: 17, Rank: 156, Name: "Jayden Higgins", Team: "UC", Position: model.WR, ByeWeek: 6, ADP: 148.0},
		{Tier: 17, Rank: 157, Name: "Denver Broncos", Team: "DEN", Position: model.DST, ByeWeek: 12, ADP: 123.3},
		{Tier: 17, Rank: 158, Name: "Zach Ertz", Team: "WAS", Position: model.TE, ByeWeek: 12, ADP: 161.0},
		{Tier: 17, Rank: 159, Name: "Romeo Doubs", Team: "GB", Position: model.WR, ByeWeek: 5, ADP: 208.0},
		{Tier: 17, Rank: 160, Name: "Isaiah Likely", Team: "BAL", Position: model.TE, ByeWeek: 7, ADP: 193.3},
		{Tier: 17, Rank: 161, Name: "Philadelphia Eagles", Team: "PHI", Position: model.DST, ByeWeek: 9, ADP: 128.3},

		// Tier 17 continued (6 players) - Boris Chen Rank 162-167
		{Tier: 17, Rank: 162, Name: "Rico Dowdle", Team: "DAL", Position: model.RB, ByeWeek: 10, ADP: 164.7},
		{Tier: 17, Rank: 163, Name: "Marquise Brown", Team: "KC", Position: model.WR, ByeWeek: 10, ADP: 173.3},
		{Tier: 17, Rank: 164, Name: "Keenan Allen", Team: "CHI", Position: model.WR, ByeWeek: 5, ADP: 147.7},
		{Tier: 17, Rank: 165, Name: "Pittsburgh Steelers", Team: "PIT", Position: model.DST, ByeWeek: 5, ADP: 144.3},
		{Tier: 17, Rank: 166, Name: "Tre' Harris", Team: "OLE", Position: model.WR, ByeWeek: 12, ADP: 174.0},
		{Tier: 17, Rank: 167, Name: "Isaac Guerendo", Team: "SF", Position: model.RB, ByeWeek: 14, ADP: 180.7},

		// Tier 18 (2 players) - Boris Chen Rank 168-169
		{Tier: 18, Rank: 168, Name: "Baltimore Ravens", Team: "BAL", Position: model.DST, ByeWeek: 7, ADP: 155.0},
		{Tier: 18, Rank: 169, Name: "Minnesota Vikings", Team: "MIN", Position: model.DST, ByeWeek: 6, ADP: 161.0},

		// Tier 18 continued (1 player) - Boris Chen Rank 170
		{Tier: 18, Rank: 170, Name: "Joshua Palmer", Team: "LAC", Position: model.WR, ByeWeek: 12, ADP: 172.3},

		// Tier 19 (9 players) - Boris Chen Rank 171-179
		{Tier: 19, Rank: 171, Name: "DeMario Douglas", Team: "NE", Position: model.WR, ByeWeek: 14, ADP: 191.0},
		{Tier: 19, Rank: 172, Name: "Adam Thielen", Team: "CAR", Position: model.WR, ByeWeek: 14, ADP: 175.7},
		{Tier: 19, Rank: 173, Name: "Kyle Williams", Team: "PSU", Position: model.WR, ByeWeek: 14, ADP: 183.7},
		{Tier: 19, Rank: 174, Name: "Ollie Gordon II", Team: "OKST", Position: model.RB, ByeWeek: 12, ADP: 159.7},
		{Tier: 19, Rank: 175, Name: "Jaylen Wright", Team: "MIA", Position: model.RB, ByeWeek: 12, ADP: 190.7},
		{Tier: 19, Rank: 176, Name: "Quentin Johnston", Team: "LAC", Position: model.WR, ByeWeek: 12, ADP: 260.5},
		{Tier: 19, Rank: 177, Name: "Houston Texans", Team: "HOU", Position: model.DST, ByeWeek: 6, ADP: 186.86},
		{Tier: 19, Rank: 178, Name: "Roschon Johnson", Team: "CHI", Position: model.RB, ByeWeek: 5, ADP: 225.0},

		// Tier 20 (6 players) - Boris Chen Rank 179-184
		{Tier: 20, Rank: 179, Name: "Wan'Dale Robinson", Team: "NYG", Position: model.WR, ByeWeek: 14, ADP: 197.5},
		{Tier: 20, Rank: 180, Name: "Buffalo Bills", Team: "BUF", Position: model.DST, ByeWeek: 7, ADP: 188.37},
		{Tier: 20, Rank: 181, Name: "Brandon Aubrey", Team: "DAL", Position: model.K, ByeWeek: 10, ADP: 119.7},
		{Tier: 20, Rank: 182, Name: "Kansas City Chiefs", Team: "KC", Position: model.DST, ByeWeek: 10, ADP: 190.25},
		{Tier: 20, Rank: 183, Name: "Detroit Lions", Team: "DET", Position: model.DST, ByeWeek: 8, ADP: 189.3},
		{Tier: 20, Rank: 184, Name: "Blake Corum", Team: "LAR", Position: model.RB, ByeWeek: 8, ADP: 194.0},

		// Tier 21 (8 players) - Boris Chen Rank 185-192
		{Tier: 21, Rank: 185, Name: "Cam Ward", Team: "MIA", Position: model.QB, ByeWeek: 12, ADP: 159.3},
		{Tier: 21, Rank: 186, Name: "Xavier Legette", Team: "CAR", Position: model.WR, ByeWeek: 14, ADP: 189.5},
		{Tier: 21, Rank: 187, Name: "Brenton Strange", Team: "JAC", Position: model.TE, ByeWeek: 8, ADP: 189.7},
		{Tier: 21, Rank: 188, Name: "Sam Darnold", Team: "MIN", Position: model.QB, ByeWeek: 6, ADP: 178.3},
		{Tier: 21, Rank: 189, Name: "Jalen Coker", Team: "CAR", Position: model.WR, ByeWeek: 14, ADP: 246.5},

		// Tier 22 (4 players) - Boris Chen Rank 190-193
		{Tier: 22, Rank: 190, Name: "DeAndre Hopkins", Team: "KC", Position: model.WR, ByeWeek: 10, ADP: 176.7},
		{Tier: 22, Rank: 191, Name: "Cameron Dicker", Team: "LAC", Position: model.K, ByeWeek: 12, ADP: 136.3},
		{Tier: 22, Rank: 192, Name: "Jake Bates", Team: "DET", Position: model.K, ByeWeek: 8, ADP: 146.7},

		// Tier 23 (2 players) - Boris Chen Rank 194-195
		{Tier: 23, Rank: 193, Name: "Seattle Seahawks", Team: "SEA", Position: model.DST, ByeWeek: 8, ADP: 232.7},
		{Tier: 23, Rank: 194, Name: "Alec Pierce", Team: "IND", Position: model.WR, ByeWeek: 11, ADP: 207.39},

		// Tier 24 (6 players) - Boris Chen Rank 196-201
		{Tier: 24, Rank: 195, Name: "Cade Otton", Team: "TB", Position: model.TE, ByeWeek: 9, ADP: 190.0},
		{Tier: 24, Rank: 196, Name: "Wil Lutz", Team: "DEN", Position: model.K, ByeWeek: 12, ADP: 201.3},
		{Tier: 24, Rank: 197, Name: "Los Angeles Rams", Team: "LAR", Position: model.DST, ByeWeek: 8, ADP: 235.0},
		{Tier: 24, Rank: 198, Name: "Green Bay Packers", Team: "GB", Position: model.DST, ByeWeek: 5, ADP: 234.0},
		{Tier: 24, Rank: 199, Name: "Chig Okonkwo", Team: "TEN", Position: model.TE, ByeWeek: 10, ADP: 183.7},
		{Tier: 24, Rank: 200, Name: "Will Shipley", Team: "PHI", Position: model.RB, ByeWeek: 9, ADP: 190.5},

		// Now adding FantasyPros players NOT in Boris Chen data - assigned tiers 25+
		// Tier 25+ for non-Boris Chen players
		{Tier: 25, Rank: 201, Name: "Pat Freiermuth", Team: "PIT", Position: model.TE, ByeWeek: 5, ADP: 204.3},
		{Tier: 25, Rank: 202, Name: "Mike Gesicki", Team: "CIN", Position: model.TE, ByeWeek: 10, ADP: 212.3},
		{Tier: 25, Rank: 203, Name: "Tyler Bass", Team: "BUF", Position: model.K, ByeWeek: 7, ADP: 209.0},
		{Tier: 25, Rank: 204, Name: "Jake Elliott", Team: "PHI", Position: model.K, ByeWeek: 9, ADP: 215.0},
		{Tier: 25, Rank: 205, Name: "Younghoe Koo", Team: "ATL", Position: model.K, ByeWeek: 5, ADP: 217.0},
		{Tier: 25, Rank: 206, Name: "Ka'imi Fairbairn", Team: "HOU", Position: model.K, ByeWeek: 6, ADP: 181.0},
		{Tier: 25, Rank: 207, Name: "Harrison Butker", Team: "KC", Position: model.K, ByeWeek: 10, ADP: 185.7},
		{Tier: 25, Rank: 208, Name: "Tampa Bay Buccaneers", Team: "TB", Position: model.DST, ByeWeek: 9, ADP: 204.3},
		{Tier: 25, Rank: 209, Name: "Kendre Miller", Team: "NO", Position: model.RB, ByeWeek: 11, ADP: 228.0},
		{Tier: 25, Rank: 210, Name: "Anthony Richardson Sr.", Team: "IND", Position: model.QB, ByeWeek: 11, ADP: 226.5},

		// Additional Tier 26+ players from FantasyPros
		{Tier: 26, Rank: 211, Name: "Russell Wilson", Team: "PIT", Position: model.QB, ByeWeek: 5, ADP: 228.5},
		{Tier: 26, Rank: 212, Name: "Shedeur Sanders", Team: "COL", Position: model.QB, ByeWeek: 9, ADP: 195.0},
		{Tier: 26, Rank: 213, Name: "Chris Rodriguez Jr.", Team: "WAS", Position: model.RB, ByeWeek: 12, ADP: 176.0},
		{Tier: 26, Rank: 214, Name: "Kyle Monangai", Team: "RUT", Position: model.RB, ByeWeek: 5, ADP: 182.0},
		{Tier: 26, Rank: 215, Name: "Justice Hill", Team: "BAL", Position: model.RB, ByeWeek: 7, ADP: 205.0},
		{Tier: 26, Rank: 216, Name: "Keaton Mitchell", Team: "BAL", Position: model.RB, ByeWeek: 7, ADP: 203.0},
		{Tier: 26, Rank: 217, Name: "Tyler Lockett", Team: "SEA", Position: model.WR, ByeWeek: 8, ADP: 259.5},
		{Tier: 26, Rank: 218, Name: "Darius Slayton", Team: "NYG", Position: model.WR, ByeWeek: 14, ADP: 259.0},
		{Tier: 26, Rank: 219, Name: "Jaylin Noel", Team: "HOU", Position: model.WR, ByeWeek: 6, ADP: 255.0},
		{Tier: 26, Rank: 220, Name: "Tre Tucker", Team: "LV", Position: model.WR, ByeWeek: 8, ADP: 213.0},

		// Final batch - Tier 27+ for deeper picks
		{Tier: 27, Rank: 221, Name: "Aaron Rodgers", Team: "NYJ", Position: model.QB, ByeWeek: 9, ADP: 176.3},
		{Tier: 27, Rank: 222, Name: "Daniel Jones", Team: "NYG", Position: model.QB, ByeWeek: 14, ADP: 244.5},
		{Tier: 27, Rank: 223, Name: "Joe Flacco", Team: "IND", Position: model.QB, ByeWeek: 11, ADP: 245.0},
		{Tier: 27, Rank: 224, Name: "Kirk Cousins", Team: "ATL", Position: model.QB, ByeWeek: 5, ADP: 220.0},
		{Tier: 27, Rank: 225, Name: "Jaxson Dart", Team: "OLE", Position: model.QB, ByeWeek: 14, ADP: 224.0},
		{Tier: 27, Rank: 226, Name: "Washington Commanders", Team: "WAS", Position: model.DST, ByeWeek: 12, ADP: 220.0},
		{Tier: 27, Rank: 227, Name: "New York Giants", Team: "NYG", Position: model.DST, ByeWeek: 14, ADP: 223.3},
		{Tier: 27, Rank: 228, Name: "New York Jets", Team: "NYJ", Position: model.DST, ByeWeek: 9, ADP: 239.0},
		{Tier: 27, Rank: 229, Name: "Miles Sanders", Team: "CAR", Position: model.RB, ByeWeek: 14, ADP: 249.0},
		{Tier: 27, Rank: 230, Name: "Raheem Mostert", Team: "MIA", Position: model.RB, ByeWeek: 12, ADP: 256.0},

		// Additional players to reach 300+
		{Tier: 28, Rank: 231, Name: "Dameon Pierce", Team: "HOU", Position: model.RB, ByeWeek: 6, ADP: 221.0},
		{Tier: 28, Rank: 232, Name: "MarShawn Lloyd", Team: "GB", Position: model.RB, ByeWeek: 5, ADP: 272.0},
		{Tier: 28, Rank: 233, Name: "Sean Tucker", Team: "TB", Position: model.RB, ByeWeek: 9, ADP: 255.0},
		{Tier: 28, Rank: 234, Name: "Devin Singletary", Team: "NYG", Position: model.RB, ByeWeek: 14, ADP: 254.0},
		{Tier: 28, Rank: 235, Name: "Jason Sanders", Team: "MIA", Position: model.K, ByeWeek: 12, ADP: 254.0},
		{Tier: 28, Rank: 236, Name: "Joshua Karty", Team: "LAR", Position: model.K, ByeWeek: 8, ADP: 210.0},
		{Tier: 28, Rank: 237, Name: "Evan McPherson", Team: "CIN", Position: model.K, ByeWeek: 10, ADP: 213.7},
		{Tier: 28, Rank: 238, Name: "Tyler Loop", Team: "BAL", Position: model.K, ByeWeek: 7, ADP: 213.3},
		{Tier: 28, Rank: 239, Name: "Daniel Carlson", Team: "LV", Position: model.K, ByeWeek: 8, ADP: 226.7},
		{Tier: 28, Rank: 240, Name: "Matt Gay", Team: "IND", Position: model.K, ByeWeek: 11, ADP: 226.7},

		// Final players to complete the list
		{Tier: 29, Rank: 241, Name: "Will Reichard", Team: "MIN", Position: model.K, ByeWeek: 6, ADP: 238.0},
		{Tier: 29, Rank: 242, Name: "Chris Boswell", Team: "PIT", Position: model.K, ByeWeek: 5, ADP: 172.7},
		{Tier: 29, Rank: 243, Name: "Chase McLaughlin", Team: "TB", Position: model.K, ByeWeek: 9, ADP: 188.0},
		{Tier: 29, Rank: 244, Name: "Jake Moody", Team: "SF", Position: model.K, ByeWeek: 14, ADP: 230.0},
		{Tier: 29, Rank: 245, Name: "Brandon McManus", Team: "GB", Position: model.K, ByeWeek: 5, ADP: 216.5},
		{Tier: 29, Rank: 246, Name: "Tyler Conklin", Team: "NYJ", Position: model.TE, ByeWeek: 9, ADP: 277.5},
		{Tier: 29, Rank: 247, Name: "Dalton Schultz", Team: "HOU", Position: model.TE, ByeWeek: 6, ADP: 242.0},
		{Tier: 29, Rank: 248, Name: "Cole Kmet", Team: "CHI", Position: model.TE, ByeWeek: 5, ADP: 275.0},
		{Tier: 29, Rank: 249, Name: "Noah Fant", Team: "SEA", Position: model.TE, ByeWeek: 8, ADP: 342.0},
		{Tier: 29, Rank: 250, Name: "Tyler Higbee", Team: "LAR", Position: model.TE, ByeWeek: 8, ADP: 299.0},

		// Final batch of WRs and remaining players
		{Tier: 30, Rank: 251, Name: "Taysom Hill", Team: "NO", Position: model.TE, ByeWeek: 11, ADP: 307.0},
		{Tier: 30, Rank: 252, Name: "Calvin Austin III", Team: "PIT", Position: model.WR, ByeWeek: 5, ADP: 281.0},
		{Tier: 30, Rank: 253, Name: "Roman Wilson", Team: "PIT", Position: model.WR, ByeWeek: 5, ADP: 269.0},
		{Tier: 30, Rank: 254, Name: "Andrei Iosivas", Team: "CIN", Position: model.WR, ByeWeek: 10, ADP: 269.0},
		{Tier: 30, Rank: 255, Name: "Adonai Mitchell", Team: "IND", Position: model.WR, ByeWeek: 11, ADP: 270.0},
		{Tier: 30, Rank: 256, Name: "Michael Wilson", Team: "ARI", Position: model.WR, ByeWeek: 8, ADP: 272.5},
		{Tier: 30, Rank: 257, Name: "Jalen Royals", Team: "KC", Position: model.WR, ByeWeek: 10, ADP: 275.0},
		{Tier: 30, Rank: 258, Name: "Dontayvion Wicks", Team: "GB", Position: model.WR, ByeWeek: 5, ADP: 299.5},
		{Tier: 30, Rank: 259, Name: "Isaiah Bond", Team: "TEX", Position: model.WR, ByeWeek: 9, ADP: 297.0},
		{Tier: 30, Rank: 260, Name: "Christian Watson", Team: "GB", Position: model.WR, ByeWeek: 5, ADP: 330.0},

		// Complete the list with final players
		{Tier: 30, Rank: 261, Name: "Jalen Tolbert", Team: "DAL", Position: model.WR, ByeWeek: 10, ADP: 347.0},
		{Tier: 30, Rank: 262, Name: "Brandin Cooks", Team: "DAL", Position: model.WR, ByeWeek: 10, ADP: 339.0},
		{Tier: 30, Rank: 263, Name: "Jaleel McLaughlin", Team: "DEN", Position: model.RB, ByeWeek: 12, ADP: 308.0},
	}

	return players
}
package data

import "ffdraft/model"

func GetCorrectedFinalPlayers() []model.Player {
	players := []model.Player{
		// BORIS CHEN TIER 1 (Ranks 1-6)
		{Tier: 1, Rank: 1, Name: "Ja'Marr Chase", Team: "CIN", Position: model.WR, ByeWeek: 10, ADP: 1.0},
		{Tier: 1, Rank: 2, Name: "Bijan Robinson", Team: "ATL", Position: model.RB, ByeWeek: 5, ADP: 2.3},
		{Tier: 1, Rank: 3, Name: "Jahmyr Gibbs", Team: "DET", Position: model.RB, ByeWeek: 8, ADP: 3.7},
		{Tier: 1, Rank: 4, Name: "Saquon Barkley", Team: "PHI", Position: model.RB, ByeWeek: 9, ADP: 3.0},
		{Tier: 1, Rank: 5, Name: "CeeDee Lamb", Team: "DAL", Position: model.WR, ByeWeek: 10, ADP: 5.3},
		{Tier: 1, Rank: 6, Name: "Justin Jefferson", Team: "MIN", Position: model.WR, ByeWeek: 6, ADP: 5.7},

		// BORIS CHEN TIER 2 (Ranks 7-15)
		{Tier: 2, Rank: 7, Name: "Malik Nabers", Team: "NYG", Position: model.WR, ByeWeek: 14, ADP: 11.0},
		{Tier: 2, Rank: 8, Name: "Nico Collins", Team: "HOU", Position: model.WR, ByeWeek: 6, ADP: 11.0},
		{Tier: 2, Rank: 9, Name: "Christian McCaffrey", Team: "SF", Position: model.RB, ByeWeek: 14, ADP: 7.7},
		{Tier: 2, Rank: 10, Name: "Amon-Ra St. Brown", Team: "DET", Position: model.WR, ByeWeek: 8, ADP: 9.3},
		{Tier: 2, Rank: 11, Name: "Derrick Henry", Team: "BAL", Position: model.RB, ByeWeek: 7, ADP: 9.0},
		{Tier: 2, Rank: 12, Name: "Puka Nacua", Team: "LAR", Position: model.WR, ByeWeek: 8, ADP: 13.7},
		{Tier: 2, Rank: 13, Name: "Ashton Jeanty", Team: "LV", Position: model.RB, ByeWeek: 8, ADP: 9.7},
		{Tier: 2, Rank: 14, Name: "Brian Thomas Jr.", Team: "JAX", Position: model.WR, ByeWeek: 8, ADP: 15.7},
		{Tier: 2, Rank: 15, Name: "Drake London", Team: "ATL", Position: model.WR, ByeWeek: 5, ADP: 17.7},

		// BORIS CHEN TIER 3 (Ranks 16-25)
		{Tier: 3, Rank: 16, Name: "De'Von Achane", Team: "MIA", Position: model.RB, ByeWeek: 12, ADP: 16.3},
		{Tier: 3, Rank: 17, Name: "A.J. Brown", Team: "PHI", Position: model.WR, ByeWeek: 9, ADP: 22.0},
		{Tier: 3, Rank: 18, Name: "Brock Bowers", Team: "LV", Position: model.TE, ByeWeek: 8, ADP: 20.7},
		{Tier: 3, Rank: 19, Name: "Jonathan Taylor", Team: "IND", Position: model.RB, ByeWeek: 11, ADP: 18.0},
		{Tier: 3, Rank: 20, Name: "Josh Jacobs", Team: "GB", Position: model.RB, ByeWeek: 5, ADP: 14.7},
		{Tier: 3, Rank: 21, Name: "Bucky Irving", Team: "TB", Position: model.RB, ByeWeek: 9, ADP: 19.3},
		{Tier: 3, Rank: 22, Name: "Chase Brown", Team: "CIN", Position: model.RB, ByeWeek: 10, ADP: 21.3},
		{Tier: 3, Rank: 23, Name: "Ladd McConkey", Team: "LAC", Position: model.WR, ByeWeek: 12, ADP: 26.3},
		{Tier: 3, Rank: 24, Name: "Trey McBride", Team: "ARI", Position: model.TE, ByeWeek: 8, ADP: 27.7},

		// BORIS CHEN TIER 4 (Ranks 25-29)
		{Tier: 4, Rank: 25, Name: "Josh Allen", Team: "BUF", Position: model.QB, ByeWeek: 7, ADP: 21.3},
		{Tier: 4, Rank: 26, Name: "Kyren Williams", Team: "LAR", Position: model.RB, ByeWeek: 8, ADP: 24.0},
		{Tier: 4, Rank: 27, Name: "Lamar Jackson", Team: "BAL", Position: model.QB, ByeWeek: 7, ADP: 22.3},
		{Tier: 4, Rank: 28, Name: "Tee Higgins", Team: "CIN", Position: model.WR, ByeWeek: 10, ADP: 31.7},
		{Tier: 4, Rank: 29, Name: "Jaxon Smith-Njigba", Team: "SEA", Position: model.WR, ByeWeek: 8, ADP: 32.3},

		// BORIS CHEN TIER 5 (Ranks 30-38)
		{Tier: 5, Rank: 30, Name: "George Kittle", Team: "SF", Position: model.TE, ByeWeek: 14, ADP: 35.7},
		{Tier: 5, Rank: 31, Name: "Tyreek Hill", Team: "MIA", Position: model.WR, ByeWeek: 12, ADP: 31.3},
		{Tier: 5, Rank: 32, Name: "Mike Evans", Team: "TB", Position: model.WR, ByeWeek: 9, ADP: 39.7},
		{Tier: 5, Rank: 33, Name: "Jayden Daniels", Team: "WAS", Position: model.QB, ByeWeek: 12, ADP: 30.7},
		{Tier: 5, Rank: 34, Name: "Garrett Wilson", Team: "NYJ", Position: model.WR, ByeWeek: 9, ADP: 40.3},
		{Tier: 5, Rank: 35, Name: "James Cook", Team: "BUF", Position: model.RB, ByeWeek: 7, ADP: 28.7},
		{Tier: 5, Rank: 36, Name: "Davante Adams", Team: "LAR", Position: model.WR, ByeWeek: 8, ADP: 43.7},
		{Tier: 5, Rank: 37, Name: "Omarion Hampton", Team: "LAC", Position: model.RB, ByeWeek: 12, ADP: 31.0},
		{Tier: 5, Rank: 38, Name: "Jalen Hurts", Team: "PHI", Position: model.QB, ByeWeek: 9, ADP: 36.7},

		// BORIS CHEN TIER 6 (Ranks 39-42)
		{Tier: 6, Rank: 39, Name: "Kenneth Walker III", Team: "SEA", Position: model.RB, ByeWeek: 8, ADP: 39.3},
		{Tier: 6, Rank: 40, Name: "Marvin Harrison Jr.", Team: "ARI", Position: model.WR, ByeWeek: 8, ADP: 39.3},
		{Tier: 6, Rank: 41, Name: "Terry McLaurin", Team: "WAS", Position: model.WR, ByeWeek: 12, ADP: 41.3},
		{Tier: 6, Rank: 42, Name: "Alvin Kamara", Team: "NO", Position: model.RB, ByeWeek: 11, ADP: 38.7},

		// BORIS CHEN TIER 7 (Ranks 43-50)
		{Tier: 7, Rank: 43, Name: "Breece Hall", Team: "NYJ", Position: model.RB, ByeWeek: 9, ADP: 37.7},
		{Tier: 7, Rank: 44, Name: "DJ Moore", Team: "CHI", Position: model.WR, ByeWeek: 5, ADP: 49.0},
		{Tier: 7, Rank: 45, Name: "Courtland Sutton", Team: "DEN", Position: model.WR, ByeWeek: 12, ADP: 52.7},
		{Tier: 7, Rank: 46, Name: "Joe Burrow", Team: "CIN", Position: model.QB, ByeWeek: 10, ADP: 38.7},
		{Tier: 7, Rank: 47, Name: "DK Metcalf", Team: "PIT", Position: model.WR, ByeWeek: 5, ADP: 49.0},
		{Tier: 7, Rank: 48, Name: "TreVeyon Henderson", Team: "NE", Position: model.RB, ByeWeek: 14, ADP: 48.0},
		{Tier: 7, Rank: 49, Name: "Chuba Hubbard", Team: "CAR", Position: model.RB, ByeWeek: 14, ADP: 42.7},
		{Tier: 7, Rank: 50, Name: "James Conner", Team: "ARI", Position: model.RB, ByeWeek: 8, ADP: 47.0},

		// BORIS CHEN TIER 8 (Ranks 51-55)
		{Tier: 8, Rank: 51, Name: "Tetairoa McMillan", Team: "CAR", Position: model.WR, ByeWeek: 14, ADP: 61.3},
		{Tier: 8, Rank: 52, Name: "DeVonta Smith", Team: "PHI", Position: model.WR, ByeWeek: 9, ADP: 55.3},
		{Tier: 8, Rank: 53, Name: "Jameson Williams", Team: "DET", Position: model.WR, ByeWeek: 8, ADP: 62.0},
		{Tier: 8, Rank: 54, Name: "Xavier Worthy", Team: "KC", Position: model.WR, ByeWeek: 10, ADP: 57.7},
		{Tier: 8, Rank: 55, Name: "Calvin Ridley", Team: "TEN", Position: model.WR, ByeWeek: 10, ADP: 67.7},

		// BORIS CHEN TIER 9 (Ranks 56-64)
		{Tier: 9, Rank: 56, Name: "George Pickens", Team: "DAL", Position: model.WR, ByeWeek: 10, ADP: 60.3},
		{Tier: 9, Rank: 57, Name: "Tony Pollard", Team: "TEN", Position: model.RB, ByeWeek: 10, ADP: 58.3},
		{Tier: 9, Rank: 58, Name: "Zay Flowers", Team: "BAL", Position: model.WR, ByeWeek: 7, ADP: 63.0},
		{Tier: 9, Rank: 59, Name: "D'Andre Swift", Team: "CHI", Position: model.RB, ByeWeek: 5, ADP: 59.7},
		{Tier: 9, Rank: 60, Name: "Jaylen Waddle", Team: "MIA", Position: model.WR, ByeWeek: 12, ADP: 70.0},
		{Tier: 9, Rank: 61, Name: "RJ Harvey", Team: "DEN", Position: model.RB, ByeWeek: 12, ADP: 52.7},
		{Tier: 9, Rank: 62, Name: "David Montgomery", Team: "DET", Position: model.RB, ByeWeek: 8, ADP: 54.7},
		{Tier: 9, Rank: 63, Name: "Isiah Pacheco", Team: "KC", Position: model.RB, ByeWeek: 10, ADP: 58.7},
		{Tier: 9, Rank: 64, Name: "Sam LaPorta", Team: "DET", Position: model.TE, ByeWeek: 8, ADP: 54.0},

		// BORIS CHEN TIER 10 (Ranks 65-67)
		{Tier: 10, Rank: 65, Name: "Aaron Jones Sr.", Team: "MIN", Position: model.RB, ByeWeek: 6, ADP: 66.3},
		{Tier: 10, Rank: 66, Name: "Patrick Mahomes II", Team: "KC", Position: model.QB, ByeWeek: 10, ADP: 52.3},
		{Tier: 10, Rank: 67, Name: "Travis Hunter", Team: "JAC", Position: model.WR, ByeWeek: 8, ADP: 72.0},

		// BORIS CHEN TIER 11 (Ranks 68-77)
		{Tier: 11, Rank: 68, Name: "T.J. Hockenson", Team: "MIN", Position: model.TE, ByeWeek: 6, ADP: 64.0},
		{Tier: 11, Rank: 69, Name: "Chris Olave", Team: "NO", Position: model.WR, ByeWeek: 11, ADP: 77.7},
		{Tier: 11, Rank: 70, Name: "Baker Mayfield", Team: "TB", Position: model.QB, ByeWeek: 9, ADP: 68.3},
		{Tier: 11, Rank: 71, Name: "Rome Odunze", Team: "CHI", Position: model.WR, ByeWeek: 5, ADP: 80.0},
		{Tier: 11, Rank: 72, Name: "Rashee Rice", Team: "KC", Position: model.WR, ByeWeek: 10, ADP: 64.7},
		{Tier: 11, Rank: 73, Name: "Tyrone Tracy Jr.", Team: "NYG", Position: model.RB, ByeWeek: 14, ADP: 76.0},
		{Tier: 11, Rank: 74, Name: "Bo Nix", Team: "DEN", Position: model.QB, ByeWeek: 12, ADP: 71.7},
		{Tier: 11, Rank: 75, Name: "Jerry Jeudy", Team: "CLE", Position: model.WR, ByeWeek: 9, ADP: 77.7},
		{Tier: 11, Rank: 76, Name: "Ricky Pearsall", Team: "SF", Position: model.WR, ByeWeek: 14, ADP: 83.7},

		// BORIS CHEN TIER 12 (Ranks 77-83)
		{Tier: 12, Rank: 77, Name: "Emeka Egbuka", Team: "TB", Position: model.WR, ByeWeek: 9, ADP: 86.7},
		{Tier: 12, Rank: 78, Name: "Kaleb Johnson", Team: "PIT", Position: model.RB, ByeWeek: 5, ADP: 74.0},
		{Tier: 12, Rank: 79, Name: "Jaylen Warren", Team: "PIT", Position: model.RB, ByeWeek: 5, ADP: 83.3},
		{Tier: 12, Rank: 80, Name: "Travis Kelce", Team: "KC", Position: model.TE, ByeWeek: 10, ADP: 62.3},
		{Tier: 12, Rank: 81, Name: "Mark Andrews", Team: "BAL", Position: model.TE, ByeWeek: 7, ADP: 75.7},
		{Tier: 12, Rank: 82, Name: "Kyler Murray", Team: "ARI", Position: model.QB, ByeWeek: 8, ADP: 93.7},
		{Tier: 12, Rank: 83, Name: "Stefon Diggs", Team: "NE", Position: model.WR, ByeWeek: 14, ADP: 91.3},

		// BORIS CHEN TIER 13 (Ranks 84-93)
		{Tier: 13, Rank: 84, Name: "Dak Prescott", Team: "DAL", Position: model.QB, ByeWeek: 10, ADP: 95.7},
		{Tier: 13, Rank: 85, Name: "Jakobi Meyers", Team: "LV", Position: model.WR, ByeWeek: 8, ADP: 92.3},
		{Tier: 13, Rank: 86, Name: "David Njoku", Team: "CLE", Position: model.TE, ByeWeek: 9, ADP: 90.0},
		{Tier: 13, Rank: 87, Name: "Jordan Mason", Team: "MIN", Position: model.RB, ByeWeek: 6, ADP: 89.3},
		{Tier: 13, Rank: 88, Name: "Travis Etienne Jr.", Team: "JAC", Position: model.RB, ByeWeek: 8, ADP: 93.0},
		{Tier: 13, Rank: 89, Name: "Deebo Samuel Sr.", Team: "WAS", Position: model.WR, ByeWeek: 12, ADP: 87.7},
		{Tier: 13, Rank: 90, Name: "Evan Engram", Team: "DEN", Position: model.TE, ByeWeek: 12, ADP: 84.3},
		{Tier: 13, Rank: 91, Name: "Jordan Addison", Team: "MIN", Position: model.WR, ByeWeek: 6, ADP: 92.3},
		{Tier: 13, Rank: 92, Name: "Zach Charbonnet", Team: "SEA", Position: model.RB, ByeWeek: 8, ADP: 99.7},
		{Tier: 13, Rank: 93, Name: "Brock Purdy", Team: "SF", Position: model.QB, ByeWeek: 14, ADP: 97.3},

		// BORIS CHEN TIER 14 (Ranks 94-104)
		{Tier: 14, Rank: 94, Name: "Justin Fields", Team: "NYJ", Position: model.QB, ByeWeek: 9, ADP: 113.0},
		{Tier: 14, Rank: 95, Name: "Jauan Jennings", Team: "SF", Position: model.WR, ByeWeek: 14, ADP: 103.3},
		{Tier: 14, Rank: 96, Name: "Matthew Golden", Team: "GB", Position: model.WR, ByeWeek: 5, ADP: 97.0},
		{Tier: 14, Rank: 97, Name: "Tucker Kraft", Team: "GB", Position: model.TE, ByeWeek: 5, ADP: 99.7},
		{Tier: 14, Rank: 98, Name: "Javonte Williams", Team: "DAL", Position: model.RB, ByeWeek: 10, ADP: 112.3},
		{Tier: 14, Rank: 99, Name: "Justin Herbert", Team: "LAC", Position: model.QB, ByeWeek: 12, ADP: 110.7},
		{Tier: 14, Rank: 100, Name: "Tyler Warren", Team: "IND", Position: model.TE, ByeWeek: 11, ADP: 98.7},
		{Tier: 14, Rank: 101, Name: "Drake Maye", Team: "NE", Position: model.QB, ByeWeek: 14, ADP: 120.0},
		{Tier: 14, Rank: 102, Name: "Caleb Williams", Team: "CHI", Position: model.QB, ByeWeek: 5, ADP: 113.7},
		{Tier: 14, Rank: 103, Name: "Khalil Shakir", Team: "BUF", Position: model.WR, ByeWeek: 7, ADP: 110.0},
		{Tier: 14, Rank: 104, Name: "J.K. Dobbins", Team: "DEN", Position: model.RB, ByeWeek: 12, ADP: 99.0},

		// BORIS CHEN TIER 15 (Ranks 105-110)
		{Tier: 15, Rank: 105, Name: "Jayden Reed", Team: "GB", Position: model.WR, ByeWeek: 5, ADP: 111.7},
		{Tier: 15, Rank: 106, Name: "Jacory Croskey-Merritt", Team: "WAS", Position: model.RB, ByeWeek: 12, ADP: 109.3},
		{Tier: 15, Rank: 107, Name: "Josh Downs", Team: "IND", Position: model.WR, ByeWeek: 11, ADP: 123.3},
		{Tier: 15, Rank: 108, Name: "Jared Goff", Team: "DET", Position: model.QB, ByeWeek: 8, ADP: 108.5},
		{Tier: 15, Rank: 109, Name: "Michael Pittman Jr.", Team: "IND", Position: model.WR, ByeWeek: 11, ADP: 122.7},
		{Tier: 15, Rank: 110, Name: "Tank Bigsby", Team: "JAC", Position: model.RB, ByeWeek: 8, ADP: 111.7},

		// BORIS CHEN TIER 16 (Ranks 111-117)
		{Tier: 16, Rank: 111, Name: "Rhamondre Stevenson", Team: "NE", Position: model.RB, ByeWeek: 14, ADP: 117.3},
		{Tier: 16, Rank: 112, Name: "Chris Godwin Jr.", Team: "TB", Position: model.WR, ByeWeek: 9, ADP: 103.0},
		{Tier: 16, Rank: 113, Name: "Cooper Kupp", Team: "SEA", Position: model.WR, ByeWeek: 8, ADP: 99.7},
		{Tier: 16, Rank: 114, Name: "Jordan Love", Team: "GB", Position: model.QB, ByeWeek: 5, ADP: 138.7},
		{Tier: 16, Rank: 115, Name: "Austin Ekeler", Team: "WAS", Position: model.RB, ByeWeek: 12, ADP: 120.7},
		{Tier: 16, Rank: 116, Name: "Trevor Lawrence", Team: "JAC", Position: model.QB, ByeWeek: 8, ADP: 118.2},

		// BORIS CHEN TIER 17 (Ranks 117-125)
		{Tier: 17, Rank: 117, Name: "Cam Skattebo", Team: "NYG", Position: model.RB, ByeWeek: 14, ADP: 107.7},
		{Tier: 17, Rank: 118, Name: "Jake Ferguson", Team: "DAL", Position: model.TE, ByeWeek: 10, ADP: 127.0},
		{Tier: 17, Rank: 119, Name: "C.J. Stroud", Team: "HOU", Position: model.QB, ByeWeek: 6, ADP: 128.7},
		{Tier: 17, Rank: 120, Name: "J.J. McCarthy", Team: "MIN", Position: model.QB, ByeWeek: 6, ADP: 139.7},
		{Tier: 17, Rank: 121, Name: "Keon Coleman", Team: "BUF", Position: model.WR, ByeWeek: 7, ADP: 131.3},
		{Tier: 17, Rank: 122, Name: "Dalton Kincaid", Team: "BUF", Position: model.TE, ByeWeek: 7, ADP: 126.3},
		{Tier: 17, Rank: 123, Name: "Darnell Mooney", Team: "ATL", Position: model.WR, ByeWeek: 5, ADP: 131.3},
		{Tier: 17, Rank: 124, Name: "Braelon Allen", Team: "NYJ", Position: model.RB, ByeWeek: 9, ADP: 115.0},
		{Tier: 17, Rank: 125, Name: "Dallas Goedert", Team: "PHI", Position: model.TE, ByeWeek: 9, ADP: 141.7},

		// BORIS CHEN TIER 18 (Ranks 126-137)
		{Tier: 18, Rank: 126, Name: "Colston Loveland", Team: "CHI", Position: model.TE, ByeWeek: 5, ADP: 117.7},
		{Tier: 18, Rank: 127, Name: "Rashid Shaheed", Team: "NO", Position: model.WR, ByeWeek: 11, ADP: 149.7},
		{Tier: 18, Rank: 128, Name: "Rachaad White", Team: "TB", Position: model.RB, ByeWeek: 9, ADP: 152.3},
		{Tier: 18, Rank: 129, Name: "Joe Mixon", Team: "HOU", Position: model.RB, ByeWeek: 6, ADP: 82.7},
		{Tier: 18, Rank: 130, Name: "Trey Benson", Team: "ARI", Position: model.RB, ByeWeek: 8, ADP: 149.7},
		{Tier: 18, Rank: 131, Name: "Brandon Aiyuk", Team: "SF", Position: model.WR, ByeWeek: 14, ADP: 147.0},
		{Tier: 18, Rank: 132, Name: "Brian Robinson Jr.", Team: "SF", Position: model.RB, ByeWeek: 14, ADP: 95.3},
		{Tier: 18, Rank: 133, Name: "Najee Harris", Team: "LAC", Position: model.RB, ByeWeek: 12, ADP: 124.3},
		{Tier: 18, Rank: 134, Name: "Christian Kirk", Team: "HOU", Position: model.WR, ByeWeek: 6, ADP: 165.7},
		{Tier: 18, Rank: 135, Name: "Ray Davis", Team: "BUF", Position: model.RB, ByeWeek: 7, ADP: 159.0},
		{Tier: 18, Rank: 136, Name: "Quinshon Judkins", Team: "CLE", Position: model.RB, ByeWeek: 9, ADP: 103.0},

		// BORIS CHEN TIER 19 (Ranks 137-155)
		{Tier: 19, Rank: 137, Name: "Nick Chubb", Team: "HOU", Position: model.RB, ByeWeek: 6, ADP: 122.3},
		{Tier: 19, Rank: 138, Name: "Tyjae Spears", Team: "TEN", Position: model.RB, ByeWeek: 10, ADP: 156.7},
		{Tier: 19, Rank: 139, Name: "Kyle Pitts Sr.", Team: "ATL", Position: model.TE, ByeWeek: 5, ADP: 133.7},
		{Tier: 19, Rank: 140, Name: "Jaydon Blue", Team: "DAL", Position: model.RB, ByeWeek: 10, ADP: 135.0},
		{Tier: 19, Rank: 141, Name: "Rashod Bateman", Team: "BAL", Position: model.WR, ByeWeek: 7, ADP: 178.0},
		{Tier: 19, Rank: 142, Name: "Tua Tagovailoa", Team: "MIA", Position: model.QB, ByeWeek: 12, ADP: 166.0},
		{Tier: 19, Rank: 143, Name: "Tyler Allgeier", Team: "ATL", Position: model.RB, ByeWeek: 5, ADP: 169.7},
		{Tier: 19, Rank: 144, Name: "Bhayshul Tuten", Team: "JAC", Position: model.RB, ByeWeek: 8, ADP: 146.7},
		{Tier: 19, Rank: 145, Name: "Marvin Mims Jr.", Team: "DEN", Position: model.WR, ByeWeek: 12, ADP: 156.0},
		{Tier: 19, Rank: 146, Name: "Cedric Tillman", Team: "CLE", Position: model.WR, ByeWeek: 9, ADP: 181.0},
		{Tier: 19, Rank: 147, Name: "Luther Burden III", Team: "CHI", Position: model.WR, ByeWeek: 5, ADP: 156.7},
		{Tier: 19, Rank: 148, Name: "Bryce Young", Team: "CAR", Position: model.QB, ByeWeek: 14, ADP: 176.0},
		{Tier: 19, Rank: 149, Name: "Hunter Henry", Team: "NE", Position: model.TE, ByeWeek: 14, ADP: 160.0},
		{Tier: 19, Rank: 150, Name: "Dylan Sampson", Team: "CLE", Position: model.RB, ByeWeek: 9, ADP: 160.3},
		{Tier: 19, Rank: 151, Name: "Jerome Ford", Team: "CLE", Position: model.RB, ByeWeek: 9, ADP: 154.3},
		{Tier: 19, Rank: 152, Name: "Michael Penix Jr.", Team: "ATL", Position: model.QB, ByeWeek: 5, ADP: 163.7},
		{Tier: 19, Rank: 153, Name: "Zach Ertz", Team: "WAS", Position: model.TE, ByeWeek: 12, ADP: 161.0},
		{Tier: 19, Rank: 154, Name: "Jayden Higgins", Team: "HOU", Position: model.WR, ByeWeek: 6, ADP: 148.0},
		{Tier: 19, Rank: 155, Name: "Keenan Allen", Team: "LAC", Position: model.WR, ByeWeek: 12, ADP: 147.7},

		// BORIS CHEN TIER 20 (Rank 153 only)
		{Tier: 20, Rank: 153, Name: "Zach Ertz", Team: "WAS", Position: model.TE, ByeWeek: 12, ADP: 160.9},

		// BORIS CHEN TIER 21 (Ranks 156-170)
		{Tier: 21, Rank: 156, Name: "Geno Smith", Team: "LV", Position: model.QB, ByeWeek: 8, ADP: 164.7},
		{Tier: 21, Rank: 157, Name: "Matthew Stafford", Team: "LAR", Position: model.QB, ByeWeek: 8, ADP: 175.0},
		{Tier: 21, Rank: 158, Name: "Jonnu Smith", Team: "PIT", Position: model.TE, ByeWeek: 5, ADP: 140.0},
		{Tier: 21, Rank: 159, Name: "Romeo Doubs", Team: "GB", Position: model.WR, ByeWeek: 5, ADP: 164.4},
		{Tier: 21, Rank: 160, Name: "Wan'Dale Robinson", Team: "NYG", Position: model.WR, ByeWeek: 14, ADP: 168.9},
		{Tier: 21, Rank: 161, Name: "Marquise Brown", Team: "KC", Position: model.WR, ByeWeek: 10, ADP: 166.5},
		{Tier: 21, Rank: 162, Name: "Rico Dowdle", Team: "CAR", Position: model.RB, ByeWeek: 14, ADP: 163.3},
		{Tier: 21, Rank: 163, Name: "DeMario Douglas", Team: "NE", Position: model.WR, ByeWeek: 14, ADP: 171.0},
		{Tier: 21, Rank: 164, Name: "Adam Thielen", Team: "MIN", Position: model.WR, ByeWeek: 6, ADP: 172.8},
		{Tier: 21, Rank: 165, Name: "Kyle Williams", Team: "NE", Position: model.WR, ByeWeek: 14, ADP: 176.4},
		{Tier: 21, Rank: 166, Name: "Brenton Strange", Team: "JAC", Position: model.TE, ByeWeek: 8, ADP: 174.7},
		{Tier: 21, Rank: 167, Name: "Joshua Palmer", Team: "BUF", Position: model.WR, ByeWeek: 7, ADP: 176.4},
		{Tier: 21, Rank: 168, Name: "Tre' Harris", Team: "LAC", Position: model.WR, ByeWeek: 12, ADP: 175.6},

		// BORIS CHEN TIER 22 (Ranks 169-179 + DSTs)
		{Tier: 22, Rank: 169, Name: "Cam Ward", Team: "TEN", Position: model.QB, ByeWeek: 10, ADP: 187.4},
		{Tier: 22, Rank: 170, Name: "Xavier Legette", Team: "CAR", Position: model.WR, ByeWeek: 14, ADP: 183.8},
		{Tier: 22, Rank: 171, Name: "Ollie Gordon II", Team: "MIA", Position: model.RB, ByeWeek: 12, ADP: 186.5},
		{Tier: 22, Rank: 172, Name: "Sam Darnold", Team: "SEA", Position: model.QB, ByeWeek: 8, ADP: 190.1},
		{Tier: 22, Rank: 173, Name: "Isaac Guerendo", Team: "SF", Position: model.RB, ByeWeek: 14, ADP: 184.5},
		{Tier: 22, Rank: 174, Name: "Roschon Johnson", Team: "CHI", Position: model.RB, ByeWeek: 5, ADP: 187.0},
		{Tier: 22, Rank: 175, Name: "Jaylen Wright", Team: "MIA", Position: model.RB, ByeWeek: 12, ADP: 190.4},
		{Tier: 22, Rank: 176, Name: "Chig Okonkwo", Team: "TEN", Position: model.TE, ByeWeek: 10, ADP: 191.5},
		{Tier: 22, Rank: 177, Name: "Quentin Johnston", Team: "LAC", Position: model.WR, ByeWeek: 12, ADP: 186.1},
		{Tier: 22, Rank: 178, Name: "Isaiah Likely", Team: "BAL", Position: model.TE, ByeWeek: 7, ADP: 193.5},
		// DSTs in Boris Chen tiers
		{Tier: 22, Rank: 195, Name: "Houston Texans", Team: "HOU", Position: model.DST, ByeWeek: 6, ADP: 186.9},
		{Tier: 22, Rank: 196, Name: "Buffalo Bills", Team: "BUF", Position: model.DST, ByeWeek: 7, ADP: 188.4},
		{Tier: 22, Rank: 197, Name: "Brandon Aubrey", Team: "DAL", Position: model.K, ByeWeek: 10, ADP: 189.1}, // CORRECTED!
		{Tier: 22, Rank: 198, Name: "Kansas City Chiefs", Team: "KC", Position: model.DST, ByeWeek: 10, ADP: 190.3},
		{Tier: 22, Rank: 200, Name: "Detroit Lions", Team: "DET", Position: model.DST, ByeWeek: 8, ADP: 192.3},

		// BORIS CHEN TIER 23 (Ranks 179-199 + remaining)
		{Tier: 23, Rank: 179, Name: "Woody Marks", Team: "HOU", Position: model.RB, ByeWeek: 6, ADP: 206.2},
		{Tier: 23, Rank: 181, Name: "Jalen Coker", Team: "CAR", Position: model.WR, ByeWeek: 14, ADP: 189.4},
		{Tier: 23, Rank: 182, Name: "Blake Corum", Team: "LAR", Position: model.RB, ByeWeek: 8, ADP: 190.0},
		{Tier: 23, Rank: 183, Name: "Cade Otton", Team: "TB", Position: model.TE, ByeWeek: 9, ADP: 204.2},
		{Tier: 23, Rank: 185, Name: "Mike Gesicki", Team: "CIN", Position: model.TE, ByeWeek: 10, ADP: 202.2},
		{Tier: 23, Rank: 188, Name: "Pat Freiermuth", Team: "PIT", Position: model.TE, ByeWeek: 5, ADP: 209.9},
		{Tier: 23, Rank: 190, Name: "DeAndre Hopkins", Team: "BAL", Position: model.WR, ByeWeek: 7, ADP: 200.5},
		{Tier: 23, Rank: 191, Name: "Kyle Monangai", Team: "CHI", Position: model.RB, ByeWeek: 5, ADP: 212.9},
		{Tier: 23, Rank: 192, Name: "Chris Rodriguez Jr.", Team: "WAS", Position: model.RB, ByeWeek: 12, ADP: 230.3},
		{Tier: 23, Rank: 193, Name: "Will Shipley", Team: "PHI", Position: model.RB, ByeWeek: 9, ADP: 200.6},
		{Tier: 23, Rank: 194, Name: "Mason Taylor", Team: "NYJ", Position: model.TE, ByeWeek: 9, ADP: 213.4}, // CORRECTED!
		{Tier: 23, Rank: 199, Name: "Alec Pierce", Team: "IND", Position: model.WR, ByeWeek: 11, ADP: 207.4},
		// DST in Boris Chen tier 19/21
		{Tier: 19, Rank: 180, Name: "Denver Broncos", Team: "DEN", Position: model.DST, ByeWeek: 12, ADP: 156.9},
		{Tier: 21, Rank: 184, Name: "Philadelphia Eagles", Team: "PHI", Position: model.DST, ByeWeek: 9, ADP: 167.6},
		{Tier: 21, Rank: 186, Name: "Pittsburgh Steelers", Team: "PIT", Position: model.DST, ByeWeek: 5, ADP: 171.9},
		{Tier: 21, Rank: 187, Name: "Baltimore Ravens", Team: "BAL", Position: model.DST, ByeWeek: 7, ADP: 174.2},
		{Tier: 21, Rank: 189, Name: "Minnesota Vikings", Team: "MIN", Position: model.DST, ByeWeek: 6, ADP: 177.6},

		// NOW ADDITIONAL FANTASYPROS PLAYERS NOT IN BORIS CHEN (TIERS 24+)
		{Tier: 24, Rank: 201, Name: "Tyler Bass", Team: "BUF", Position: model.K, ByeWeek: 7, ADP: 209.0},
		{Tier: 24, Rank: 202, Name: "Darren Waller", Team: "MIA", Position: model.TE, ByeWeek: 12, ADP: 210.7},
		{Tier: 24, Rank: 203, Name: "Mike Gesicki", Team: "CIN", Position: model.TE, ByeWeek: 10, ADP: 212.3},
		{Tier: 24, Rank: 204, Name: "Tyler Loop", Team: "BAL", Position: model.K, ByeWeek: 7, ADP: 213.3},
		{Tier: 24, Rank: 205, Name: "Evan McPherson", Team: "CIN", Position: model.K, ByeWeek: 10, ADP: 213.7},
		{Tier: 24, Rank: 206, Name: "Jake Elliott", Team: "PHI", Position: model.K, ByeWeek: 9, ADP: 215.0},
		{Tier: 24, Rank: 207, Name: "Dont'e Thornton Jr.", Team: "LV", Position: model.WR, ByeWeek: 8, ADP: 175.0},
		{Tier: 24, Rank: 208, Name: "Younghoe Koo", Team: "ATL", Position: model.K, ByeWeek: 5, ADP: 217.0},
		{Tier: 24, Rank: 209, Name: "Woody Marks", Team: "HOU", Position: model.RB, ByeWeek: 6, ADP: 177.5},
		{Tier: 24, Rank: 210, Name: "Washington Commanders", Team: "WAS", Position: model.DST, ByeWeek: 12, ADP: 220.0},

		// Continue with more FantasyPros players (Tier 24+)
		{Tier: 25, Rank: 215, Name: "Will Shipley", Team: "PHI", Position: model.RB, ByeWeek: 9, ADP: 190.5},
		{Tier: 25, Rank: 220, Name: "Justice Hill", Team: "BAL", Position: model.RB, ByeWeek: 7, ADP: 205.0},
		{Tier: 25, Rank: 225, Name: "Isaac TeSlaa", Team: "DET", Position: model.WR, ByeWeek: 8, ADP: 205.0},
		{Tier: 25, Rank: 230, Name: "Daniel Carlson", Team: "LV", Position: model.K, ByeWeek: 8, ADP: 226.7},
		{Tier: 25, Rank: 235, Name: "Matt Gay", Team: "WAS", Position: model.K, ByeWeek: 12, ADP: 226.7},
		{Tier: 25, Rank: 240, Name: "Will Reichard", Team: "MIN", Position: model.K, ByeWeek: 6, ADP: 238.0},
		{Tier: 25, Rank: 245, Name: "Chris Boswell", Team: "PIT", Position: model.K, ByeWeek: 5, ADP: 172.7},
		{Tier: 25, Rank: 250, Name: "Chase McLaughlin", Team: "TB", Position: model.K, ByeWeek: 9, ADP: 188.0},
		{Tier: 25, Rank: 255, Name: "Jake Moody", Team: "SF", Position: model.K, ByeWeek: 14, ADP: 230.0},
		{Tier: 25, Rank: 260, Name: "Brandon McManus", Team: "GB", Position: model.K, ByeWeek: 5, ADP: 216.5},

		// Additional deeper players from FantasyPros
		{Tier: 26, Rank: 265, Name: "Jason Sanders", Team: "MIA", Position: model.K, ByeWeek: 12, ADP: 254.0},
		{Tier: 26, Rank: 270, Name: "Joshua Karty", Team: "LAR", Position: model.K, ByeWeek: 8, ADP: 210.0},
		{Tier: 26, Rank: 275, Name: "Tyler Conklin", Team: "LAC", Position: model.TE, ByeWeek: 12, ADP: 277.5},
		{Tier: 26, Rank: 280, Name: "Dalton Schultz", Team: "HOU", Position: model.TE, ByeWeek: 6, ADP: 242.0},
		{Tier: 26, Rank: 285, Name: "Cole Kmet", Team: "CHI", Position: model.TE, ByeWeek: 5, ADP: 275.0},
		{Tier: 26, Rank: 290, Name: "Noah Fant", Team: "CIN", Position: model.TE, ByeWeek: 10, ADP: 342.0},
		{Tier: 26, Rank: 295, Name: "Tyler Higbee", Team: "LAR", Position: model.TE, ByeWeek: 8, ADP: 299.0},
		{Tier: 26, Rank: 300, Name: "Taysom Hill", Team: "NO", Position: model.TE, ByeWeek: 11, ADP: 307.0},
	}

	// Sort by tier first, then ADP
	model.SortPlayersByTierAndADP(players)

	return players
}
package main

import (
	"fmt"
	"ffdraft/data"
)

// Test Case 1: Verify specific Boris Chen tier assignments from tiers.txt
func testBorisChentiers() {
	players := data.GetCorrectedPlayers()
	playerMap := make(map[string]int) // name -> tier
	
	for _, p := range players {
		playerMap[p.Name] = p.Tier
	}
	
	// Spot checks from tiers.txt (line numbers correspond to ranks)
	testCases := []struct {
		name         string
		expectedTier int
		line         string
	}{
		{"Zay Flowers", 9, "Line 59: Zay Flowers tier 9"},
		{"Brandon Aubrey", 22, "Line 198: Brandon Aubrey tier 22"},  
		{"Mason Taylor", 23, "Line 195: Mason Taylor tier 23"},
		{"Alec Pierce", 23, "Line 200: Alec Pierce tier 23"},
		{"Patrick Mahomes II", 10, "Line 67: Patrick Mahomes II tier 10"},
		{"Travis Hunter", 10, "Line 68: Travis Hunter tier 10"},
		{"Denver Broncos DST", 19, "Line 181: Denver Broncos DST tier 19"},
		{"Ja'Marr Chase", 1, "Line 2: Ja'Marr Chase tier 1"},
	}
	
	fmt.Println("=== BORIS CHEN TIER TEST RESULTS ===")
	errors := 0
	
	for _, tc := range testCases {
		actualTier, exists := playerMap[tc.name]
		if !exists {
			fmt.Printf("❌ MISSING: %s not found in dataset (%s)\n", tc.name, tc.line)
			errors++
			continue
		}
		
		if actualTier != tc.expectedTier {
			fmt.Printf("❌ WRONG TIER: %s - Expected: %d, Got: %d (%s)\n", 
				tc.name, tc.expectedTier, actualTier, tc.line)
			errors++
		} else {
			fmt.Printf("✅ CORRECT: %s - Tier %d (%s)\n", tc.name, actualTier, tc.line)
		}
	}
	
	fmt.Printf("\nBoris Chen Tier Test: %d errors found\n\n", errors)
}

// Test Case 2: Verify ADP values from FantasyPros PDF
func testFantasyProsADP() {
	players := data.GetCorrectedPlayers()
	playerMap := make(map[string]float64) // name -> ADP
	
	for _, p := range players {
		playerMap[p.Name] = p.ADP
	}
	
	// Spot checks from FantasyPros PDF (approximate values)
	testCases := []struct {
		name        string
		expectedADP float64
		tolerance   float64
		source      string
	}{
		{"Ja'Marr Chase", 1.0, 0.1, "FantasyPros Rank 1"},
		{"Bijan Robinson", 2.3, 0.1, "FantasyPros Rank 2"},
		{"Saquon Barkley", 3.0, 0.1, "FantasyPros Rank 3"},
		{"Josh Allen", 21.3, 1.0, "FantasyPros early round QB"},
		{"Patrick Mahomes II", 52.3, 2.0, "FantasyPros mid-round QB"},
		{"Brandon Aubrey", 119.7, 5.0, "FantasyPros top kicker"},
	}
	
	fmt.Println("=== FANTASY PROS ADP TEST RESULTS ===")
	errors := 0
	
	for _, tc := range testCases {
		actualADP, exists := playerMap[tc.name]
		if !exists {
			fmt.Printf("❌ MISSING: %s not found in dataset (%s)\n", tc.name, tc.source)
			errors++
			continue
		}
		
		diff := actualADP - tc.expectedADP
		if diff < 0 {
			diff = -diff
		}
		
		if diff > tc.tolerance {
			fmt.Printf("❌ WRONG ADP: %s - Expected: %.1f, Got: %.1f, Diff: %.1f (%s)\n", 
				tc.name, tc.expectedADP, actualADP, diff, tc.source)
			errors++
		} else {
			fmt.Printf("✅ CORRECT: %s - ADP %.1f (Expected: %.1f ± %.1f) (%s)\n", 
				tc.name, actualADP, tc.expectedADP, tc.tolerance, tc.source)
		}
	}
	
	fmt.Printf("\nFantasyPros ADP Test: %d errors found\n\n", errors)
}

// Test Case 3: Additional comprehensive tests for Round 3
func testAdditionalBorisChentiers() {
	players := data.GetCorrectedPlayers()
	playerMap := make(map[string]int)
	
	for _, p := range players {
		playerMap[p.Name] = p.Tier
	}
	
	// Additional Boris Chen players from different tiers to test comprehensively
	testCases := []struct {
		name         string
		expectedTier int
		source       string
	}{
		{"Nico Collins", 2, "Boris Chen tier 2"},
		{"Brock Bowers", 3, "Boris Chen tier 3"}, 
		{"George Kittle", 5, "Boris Chen tier 5"},
		{"Chris Olave", 11, "Boris Chen tier 11"},
		{"Stefon Diggs", 12, "Boris Chen tier 12"},
		{"Travis Kelce", 12, "Boris Chen tier 12"},
		{"Jakobi Meyers", 13, "Boris Chen tier 13"},
		{"Justin Herbert", 14, "Boris Chen tier 14"},
		{"Josh Downs", 15, "Boris Chen tier 15"},
		{"Cooper Kupp", 16, "Boris Chen tier 16"},
		{"C.J. Stroud", 17, "Boris Chen tier 17"},
		{"Joe Mixon", 18, "Boris Chen tier 18"},
		{"Najee Harris", 18, "Boris Chen tier 18"},
		{"Houston Texans DST", 22, "Boris Chen DST tier 22"},
		{"Buffalo Bills DST", 22, "Boris Chen DST tier 22"},
		{"Kansas City Chiefs DST", 22, "Boris Chen DST tier 22"},
	}
	
	fmt.Println("=== ADDITIONAL BORIS CHEN TIER TESTS (ROUND 3) ===")
	errors := 0
	
	for _, tc := range testCases {
		actualTier, exists := playerMap[tc.name]
		if !exists {
			fmt.Printf("❌ MISSING: %s not found in dataset (%s)\n", tc.name, tc.source)
			errors++
			continue
		}
		
		if actualTier != tc.expectedTier {
			fmt.Printf("❌ WRONG TIER: %s - Expected: %d, Got: %d (%s)\n", 
				tc.name, tc.expectedTier, actualTier, tc.source)
			errors++
		} else {
			fmt.Printf("✅ CORRECT: %s - Tier %d (%s)\n", tc.name, actualTier, tc.source)
		}
	}
	
	fmt.Printf("\nAdditional Boris Chen Tests: %d errors found\n\n", errors)
}

// Test Case 4: Test mid-tier ADP accuracy
func testMidTierADP() {
	players := data.GetCorrectedPlayers()
	playerMap := make(map[string]float64)
	
	for _, p := range players {
		playerMap[p.Name] = p.ADP
	}
	
	// Mid-tier players from FantasyPros to verify ADP accuracy
	testCases := []struct {
		name        string
		expectedADP float64
		tolerance   float64
		source      string
	}{
		{"Tyreek Hill", 31.3, 1.0, "Mid-tier WR ADP"},
		{"George Kittle", 35.7, 1.0, "Top TE ADP"},
		{"Kenneth Walker III", 39.3, 1.0, "Mid-tier RB ADP"},
		{"Calvin Ridley", 67.7, 2.0, "Later WR ADP"},
		{"Sam LaPorta", 54.0, 2.0, "TE2 ADP"},
		{"Chris Olave", 77.7, 2.0, "WR in 70s ADP"},
	}
	
	fmt.Println("=== MID-TIER ADP TEST RESULTS (ROUND 3) ===")
	errors := 0
	
	for _, tc := range testCases {
		actualADP, exists := playerMap[tc.name]
		if !exists {
			fmt.Printf("❌ MISSING: %s not found in dataset (%s)\n", tc.name, tc.source)
			errors++
			continue
		}
		
		diff := actualADP - tc.expectedADP
		if diff < 0 {
			diff = -diff
		}
		
		if diff > tc.tolerance {
			fmt.Printf("❌ WRONG ADP: %s - Expected: %.1f, Got: %.1f, Diff: %.1f (%s)\n", 
				tc.name, tc.expectedADP, actualADP, diff, tc.source)
			errors++
		} else {
			fmt.Printf("✅ CORRECT: %s - ADP %.1f (Expected: %.1f ± %.1f) (%s)\n", 
				tc.name, actualADP, tc.expectedADP, tc.tolerance, tc.source)
		}
	}
	
	fmt.Printf("\nMid-tier ADP Test: %d errors found\n\n", errors)
}

// Test Case 5: Verify non-Boris Chen players have tier 24+
func testNonBorisChentiers() {
	players := data.GetCorrectedPlayers()
	
	// These players should NOT be in Boris Chen tiers.txt 
	nonBorisPlayers := []struct {
		name        string
		minTier     int
		description string
	}{
		{"Miles Sanders", 24, "Non-Boris Chen RB should be tier 24+"},
		{"Jason Sanders", 24, "Non-Boris Chen K should be tier 24+"}, 
		{"Tyler Conklin", 24, "Non-Boris Chen TE should be tier 24+"},
	}
	
	fmt.Println("=== NON-BORIS CHEN TIER TEST RESULTS ===")
	errors := 0
	
	for _, p := range players {
		if p.Tier >= 24 {
			fmt.Printf("INFO: %s assigned tier %d (non-Boris Chen)\n", p.Name, p.Tier)
		}
	}
	
	for _, tc := range nonBorisPlayers {
		for _, p := range players {
			if p.Name == tc.name {
				if p.Tier < tc.minTier {
					fmt.Printf("❌ WRONG TIER: %s - Should be tier %d+, got %d (%s)\n", 
						tc.name, tc.minTier, p.Tier, tc.description)
					errors++
				} else {
					fmt.Printf("✅ CORRECT: %s - Tier %d (%s)\n", tc.name, p.Tier, tc.description)
				}
				break
			}
		}
	}
	
	fmt.Printf("\nNon-Boris Chen Tier Test: %d errors found\n\n", errors)
}

func main() {
	fmt.Println("FFDRAFT DATA ACCURACY TEST SUITE - ROUND 3")
	fmt.Println("==========================================")
	fmt.Println()
	
	testBorisChentiers()
	testFantasyProsADP()
	testAdditionalBorisChentiers()
	testMidTierADP()
	testNonBorisChentiers()
	
	fmt.Println("Round 3 test suite complete. Review errors above and fix dataset accordingly.")
}
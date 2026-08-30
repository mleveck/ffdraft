package data

import "testing"

func TestClusterDescFixedK(t *testing.T) {
	vals := []float64{10, 9.9, 9.8, 5, 4.9, 1, 0.9}
	got := ClusterDesc(vals, 3)
	want := []int{1, 1, 1, 2, 2, 3, 3}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("ClusterDesc = %v, want %v", got, want)
		}
	}
}

func TestClusterDescMonotonic(t *testing.T) {
	vals := []float64{100, 90, 80, 70, 40, 39, 38, 10, 9, 8}
	got := ClusterDesc(vals, 4)
	if got[0] != 1 {
		t.Errorf("best value should be tier 1, got %d", got[0])
	}
	for i := 1; i < len(got); i++ {
		if got[i] < got[i-1] {
			t.Errorf("tiers must be non-decreasing down the board: %v", got)
		}
	}
	maxTier := got[len(got)-1]
	if maxTier != 4 {
		t.Errorf("expected 4 tiers used, got %d (%v)", maxTier, got)
	}
}

func TestClusterDescOptimalBoundaries(t *testing.T) {
	// Three obvious groups must be recovered exactly at k=3.
	vals := []float64{100, 99, 98, 50, 49, 48, 2, 1, 0}
	got := ClusterDesc(vals, 3)
	want := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("ClusterDesc = %v, want %v", got, want)
		}
	}
}

func TestBoardTiers(t *testing.T) {
	if k := BoardTiers(190); k != 24 {
		t.Errorf("BoardTiers(190) = %d, want 24", k)
	}
	if k := BoardTiers(10); k != 4 {
		t.Errorf("BoardTiers(10) = %d, want floor of 4", k)
	}
	if k := BoardTiers(1000); k != 28 {
		t.Errorf("BoardTiers(1000) = %d, want cap of 28", k)
	}
}

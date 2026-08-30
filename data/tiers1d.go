package data

import "math"

// ClusterDesc groups values (sorted descending) into k tiers using optimal
// 1-D k-means (dynamic programming over the sorted values, so clusters are
// contiguous and the total within-cluster SSE is minimal). This mirrors the
// clustering borischen.co applies to expert ranks — including fixing the
// tier count rather than selecting it automatically (his combined chart
// runs ~n/8 tiers, which BoardTiers reproduces) — applied here to
// league-adjusted VOR. Returns a tier number (1 = best) per index.
func ClusterDesc(valsDesc []float64, k int) []int {
	n := len(valsDesc)
	if n == 0 {
		return nil
	}
	if k < 1 {
		k = 1
	}
	if k > n {
		k = n
	}
	kMax := k

	// Work over ascending order for the standard DP formulation.
	asc := make([]float64, n)
	for i, v := range valsDesc {
		asc[n-1-i] = v
	}

	prefix := make([]float64, n+1)
	prefixSq := make([]float64, n+1)
	for i, v := range asc {
		prefix[i+1] = prefix[i] + v
		prefixSq[i+1] = prefixSq[i] + v*v
	}
	// sse of asc[j..i] inclusive
	sse := func(j, i int) float64 {
		cnt := float64(i - j + 1)
		sum := prefix[i+1] - prefix[j]
		sumSq := prefixSq[i+1] - prefixSq[j]
		s := sumSq - sum*sum/cnt
		if s < 0 {
			s = 0
		}
		return s
	}

	const inf = math.MaxFloat64
	// cost[k][i]: minimal SSE splitting asc[0..i] into k clusters.
	cost := make([][]float64, kMax+1)
	split := make([][]int, kMax+1)
	for k := range cost {
		cost[k] = make([]float64, n)
		split[k] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		cost[1][i] = sse(0, i)
	}
	for k := 2; k <= kMax; k++ {
		for i := k - 1; i < n; i++ {
			cost[k][i] = inf
			for j := k - 1; j <= i; j++ {
				c := cost[k-1][j-1] + sse(j, i)
				if c < cost[k][i] {
					cost[k][i] = c
					split[k][i] = j
				}
			}
		}
	}

	bestK := kMax

	// Recover boundaries for bestK clusters over asc.
	tiersAsc := make([]int, n)
	i, kk := n-1, bestK
	for kk >= 1 {
		j := 0
		if kk > 1 {
			j = split[kk][i]
		}
		for m := j; m <= i; m++ {
			tiersAsc[m] = kk // ascending: last cluster holds highest values
		}
		i, kk = j-1, kk-1
	}

	// Map back to descending input order; tier 1 = highest values.
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = bestK - tiersAsc[n-1-i] + 1
	}
	return out
}

// BoardTiers picks the tier count for a board of n players: one tier per
// ~8 players (the granularity of Boris's combined charts), clamped.
func BoardTiers(n int) int {
	k := (n + 4) / 8
	if k < 4 {
		k = 4
	}
	if k > 28 {
		k = 28
	}
	return k
}

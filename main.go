package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"ffdraft/data"
	"ffdraft/ui"
)

func main() {
	var newDraft bool
	var leagueKey string
	var offline bool
	var dump bool
	flag.BoolVar(&newDraft, "new", false, "Start a new draft (ignore saved state)")
	flag.StringVar(&leagueKey, "league", "fnf", "League profile: fnf (standard) or other (half PPR)")
	flag.BoolVar(&offline, "offline", false, "Prefer cached data; skip network fetch")
	flag.BoolVar(&dump, "dump", false, "Print the loaded draft board and exit")
	flag.Parse()

	league, ok := data.Leagues[leagueKey]
	if !ok {
		keys := make([]string, 0, len(data.Leagues))
		for k := range data.Leagues {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		fmt.Printf("Unknown league %q. Available: %v\n", leagueKey, keys)
		os.Exit(1)
	}

	if dump {
		players, err := data.LoadPlayers(league, offline)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s — %d players\n", league.DisplayName, len(players))
		fmt.Printf("%-4s %-4s %-24s %-4s %-4s %-4s %-6s\n", "Tier", "Rank", "Player", "Team", "Pos", "Bye", "ADP")
		for _, p := range players {
			fmt.Printf("%-4d %-4d %-24s %-4s %-4s %-4d %-6.1f\n",
				p.Tier, p.Rank, p.Name, p.Team, string(p.Position), p.ByeWeek, p.ADP)
		}
		return
	}

	if err := ui.Run(newDraft, league, offline); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

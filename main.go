package main

import (
	"flag"
	"fmt"
	"os"

	"ffdraft/ui"
)

func main() {
	var newDraft bool
	flag.BoolVar(&newDraft, "new", false, "Start a new draft (ignore saved state)")
	flag.Parse()

	if err := ui.Run(newDraft); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
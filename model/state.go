package model

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type DraftState struct {
	DraftedPlayers map[string]bool `json:"drafted_players"`
	CurrentPick    int             `json:"current_pick"`
	CurrentRound   int             `json:"current_round"`
	TotalTeams     int             `json:"total_teams"`
}

func getStateFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".ffdraft_state.json"), nil
}

func SaveDraftState(state DraftState) error {
	stateFile, err := getStateFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(stateFile, data, 0644)
}

func LoadDraftState() (DraftState, error) {
	var state DraftState
	
	stateFile, err := getStateFilePath()
	if err != nil {
		return state, err
	}

	// Check if file exists
	if _, err := os.Stat(stateFile); os.IsNotExist(err) {
		// Return empty state if file doesn't exist
		return DraftState{
			DraftedPlayers: make(map[string]bool),
			CurrentPick:    1,
			CurrentRound:   1,
			TotalTeams:     12,
		}, nil
	}

	data, err := os.ReadFile(stateFile)
	if err != nil {
		return state, err
	}

	err = json.Unmarshal(data, &state)
	if err != nil {
		return state, err
	}

	// Ensure the map is initialized
	if state.DraftedPlayers == nil {
		state.DraftedPlayers = make(map[string]bool)
	}

	return state, nil
}

func ClearDraftState() error {
	stateFile, err := getStateFilePath()
	if err != nil {
		return err
	}

	// Remove the file if it exists
	if _, err := os.Stat(stateFile); err == nil {
		return os.Remove(stateFile)
	}

	return nil
}
package ui

import (
	"fmt"

	"ffdraft/data"
	"ffdraft/model"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sahilm/fuzzy"
)

type Model struct {
	draft        *model.Draft
	cursor       int
	searchBox    textinput.Model
	width        int
	height       int
	positions    []model.Position
	activePos    int
	lastKey      string
	ctrlDPressed bool
}

func NewModel(resetState bool) Model {
	players := data.GetCorrectedPlayers()
	draft := model.NewDraft(players, 12)

	// Load saved state unless we're resetting
	if !resetState {
		if savedState, err := model.LoadDraftState(); err == nil {
			draft.LoadState(savedState)
		}
	} else {
		// Clear the saved state file when starting new
		model.ClearDraftState()
	}

	search := textinput.New()
	search.Placeholder = "Search players..."
	search.CharLimit = 50
	search.Width = 30
	search.Focus() // Always focused for immediate typing

	positions := []model.Position{
		model.ALL, model.RB, model.WR, model.TE, model.QB, model.DST, model.K,
	}

	return Model{
		draft:        draft,
		cursor:       0,
		searchBox:    search,
		positions:    positions,
		activePos:    0,
		lastKey:      "",
		ctrlDPressed: false,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		// Handle special keys first
		keyStr := msg.String()
		
		switch keyStr {
		case "q":
			// Handle qq sequence for clearing search
			if m.lastKey == "q" {
				// Second q - clear search
				m.searchBox.SetValue("")
				m.draft.SearchQuery = ""
				m.updateFilteredPlayers()
				m.cursor = 0
				m.lastKey = ""
				return m, nil
			} else {
				// First q - just store it
				m.lastKey = "q"
				return m, nil
			}
			
		case "ctrl+d":
			// Handle ctrl+d twice to exit
			if m.ctrlDPressed {
				// Second ctrl+d - exit
				return m, tea.Quit
			} else {
				// First ctrl+d - just store it
				m.ctrlDPressed = true
				return m, nil
			}

		case "ctrl+p":
			// Navigate up
			if m.cursor > 0 {
				m.cursor--
			}
			// Reset key tracking
			m.lastKey = ""
			m.ctrlDPressed = false
			return m, nil

		case "ctrl+n":
			// Navigate down
			if m.cursor < len(m.draft.FilteredPlayers)-1 {
				m.cursor++
			}
			// Reset key tracking
			m.lastKey = ""
			m.ctrlDPressed = false
			return m, nil

		case "enter":
			// Draft/undraft player
			if len(m.draft.FilteredPlayers) > 0 && m.cursor < len(m.draft.FilteredPlayers) {
				playerName := m.draft.FilteredPlayers[m.cursor].Name
				for i, p := range m.draft.Players {
					if p.Name == playerName {
						m.draft.DraftPlayer(i)
						// Clear search after drafting/undrafting
						m.searchBox.SetValue("")
						m.draft.SearchQuery = ""
						m.updateFilteredPlayers()
						m.cursor = 0
						// Auto-save state after drafting
						if state := m.draft.SaveState(); true {
							model.SaveDraftState(state)
						}
						break
					}
				}
			}
			// Reset key tracking
			m.lastKey = ""
			m.ctrlDPressed = false
			return m, nil

		case "tab", "right":
			m.activePos = (m.activePos + 1) % len(m.positions)
			m.draft.UpdateFilter(m.positions[m.activePos])
			m.updateFilteredPlayers()
			m.cursor = 0
			// Reset key tracking
			m.lastKey = ""
			m.ctrlDPressed = false

		case "shift+tab", "left":
			m.activePos = (m.activePos - 1 + len(m.positions)) % len(m.positions)
			m.draft.UpdateFilter(m.positions[m.activePos])
			m.updateFilteredPlayers()
			m.cursor = 0
			// Reset key tracking
			m.lastKey = ""
			m.ctrlDPressed = false

		case "1":
			m.activePos = 0
			m.draft.UpdateFilter(m.positions[m.activePos])
			m.updateFilteredPlayers()
			m.cursor = 0
			// Reset key tracking
			m.lastKey = ""
			m.ctrlDPressed = false
		case "2":
			m.activePos = 1
			m.draft.UpdateFilter(m.positions[m.activePos])
			m.updateFilteredPlayers()
			m.cursor = 0
			// Reset key tracking
			m.lastKey = ""
			m.ctrlDPressed = false
		case "3":
			m.activePos = 2
			m.draft.UpdateFilter(m.positions[m.activePos])
			m.updateFilteredPlayers()
			m.cursor = 0
			// Reset key tracking
			m.lastKey = ""
			m.ctrlDPressed = false
		case "4":
			m.activePos = 3
			m.draft.UpdateFilter(m.positions[m.activePos])
			m.updateFilteredPlayers()
			m.cursor = 0
			// Reset key tracking
			m.lastKey = ""
			m.ctrlDPressed = false
		case "5":
			m.activePos = 4
			m.draft.UpdateFilter(m.positions[m.activePos])
			m.updateFilteredPlayers()
			m.cursor = 0
			// Reset key tracking
			m.lastKey = ""
			m.ctrlDPressed = false
		case "6":
			m.activePos = 5
			m.draft.UpdateFilter(m.positions[m.activePos])
			m.updateFilteredPlayers()
			m.cursor = 0
			// Reset key tracking
			m.lastKey = ""
			m.ctrlDPressed = false
		case "7":
			m.activePos = 6
			m.draft.UpdateFilter(m.positions[m.activePos])
			m.updateFilteredPlayers()
			m.cursor = 0
			// Reset key tracking
			m.lastKey = ""
			m.ctrlDPressed = false
			return m, nil
		
		default:
			// Reset key tracking for any other key
			m.lastKey = ""
			m.ctrlDPressed = false
		}

		// Always process search input (search box is always focused)
		var cmd tea.Cmd
		m.searchBox, cmd = m.searchBox.Update(msg)
		cmds = append(cmds, cmd)

		if m.searchBox.Value() != m.draft.SearchQuery {
			m.draft.SearchQuery = m.searchBox.Value()
			m.updateFilteredPlayers()
			m.cursor = 0
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) updateFilteredPlayers() {
	if m.draft.SearchQuery == "" {
		// When search is empty, hide drafted players
		positionFiltered := model.FilterByPosition(m.draft.Players, m.draft.CurrentFilter)
		var availablePlayers []model.Player
		for _, p := range positionFiltered {
			if !p.Drafted {
				availablePlayers = append(availablePlayers, p)
			}
		}
		m.draft.FilteredPlayers = availablePlayers
		// Sort the filtered results by tier and ADP
		model.SortPlayersByTierAndADP(m.draft.FilteredPlayers)
		return
	}

	var searchTargets []string
	positionFiltered := model.FilterByPosition(m.draft.Players, m.draft.CurrentFilter)

	for _, p := range positionFiltered {
		searchTargets = append(searchTargets, p.Name)
	}

	matches := fuzzy.Find(m.draft.SearchQuery, searchTargets)

	var filteredPlayers []model.Player
	for _, match := range matches {
		for _, p := range positionFiltered {
			if p.Name == match.Str {
				filteredPlayers = append(filteredPlayers, p)
				break
			}
		}
	}

	m.draft.FilteredPlayers = filteredPlayers
	// Sort the filtered results by tier and ADP
	model.SortPlayersByTierAndADP(m.draft.FilteredPlayers)
}

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Loading..."
	}

	var sections []string

	// Header
	title := fmt.Sprintf("FF Draft - Round %d, Pick %d", m.draft.CurrentRound, m.draft.CurrentPick)
	sections = append(sections, headerStyle.Render(title))

	// Position tabs
	var tabs []string
	for i, pos := range m.positions {
		style := tabStyle
		if i == m.activePos {
			style = activeTabStyle
		}
		tabs = append(tabs, style.Render(string(pos)))
	}
	sections = append(sections, lipgloss.JoinHorizontal(lipgloss.Top, tabs...))

	// Search box (always active)
	searchLabel := "🔍 "
	searchBox := lipgloss.JoinHorizontal(lipgloss.Center, searchLabel, m.searchBox.View())
	sections = append(sections, searchStyle.Render(searchBox))

	// Table
	table := m.renderTable()
	sections = append(sections, table)

	// Footer
	help := "Ctrl+P/N: navigate • Enter: draft/undraft • Tab: filter • qq: clear search • Ctrl+D×2: quit"
	footer := footerStyle.Render(help)
	sections = append(sections, footer)

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func (m Model) renderTable() string {
	if len(m.draft.FilteredPlayers) == 0 {
		return baseStyle.Render("No players found")
	}

	// Table header
	header := tableHeaderStyle.Render(
		fmt.Sprintf("%-4s %-4s %-20s %-4s %-4s %-4s %-6s %-8s",
			"Tier", "Rank", "Player", "Team", "Pos", "Bye", "ADP", "Status"),
	)

	var rows []string
	maxRows := m.height - 12 // Reserve space for other UI elements
	if maxRows < 5 {
		maxRows = 5
	}

	start := 0
	end := len(m.draft.FilteredPlayers)

	// Scroll to keep cursor visible
	if m.cursor >= maxRows {
		start = m.cursor - maxRows + 1
		end = start + maxRows
		if end > len(m.draft.FilteredPlayers) {
			end = len(m.draft.FilteredPlayers)
		}
	} else {
		if end > maxRows {
			end = maxRows
		}
	}

	for i := start; i < end; i++ {
		player := m.draft.FilteredPlayers[i]
		status := "Available"
		if player.Drafted {
			status = "DRAFTED"
		}

		row := fmt.Sprintf("%-4d %-4d %-20s %-4s %-4s %-4d %-6.1f %-8s",
			player.Tier,
			player.Rank,
			truncateString(player.Name, 20),
			player.Team,
			string(player.Position),
			player.ByeWeek,
			player.ADP,
			status,
		)

		var style lipgloss.Style
		if i == m.cursor {
			style = selectedRowStyle
		} else if player.Drafted {
			style = draftedRowStyle
		} else {
			style = tableRowStyle
		}

		rows = append(rows, style.Render(row))
	}

	table := lipgloss.JoinVertical(lipgloss.Left, header)
	if len(rows) > 0 {
		table = lipgloss.JoinVertical(lipgloss.Left, table, lipgloss.JoinVertical(lipgloss.Left, rows...))
	}

	return table
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func Run(resetState bool) error {
	m := NewModel(resetState)
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	return err
}
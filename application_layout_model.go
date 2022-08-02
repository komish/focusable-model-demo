package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// applicationLayoutModel is a model that doesn't handle Update calls, but presents
// a compound view, or layout.
type applicationLayoutModel struct {
	Models []tea.Model
}

func (m *applicationLayoutModel) Init() tea.Cmd {
	cmds := []tea.Cmd{}
	for _, mod := range m.Models {
		if c := mod.Init(); c != nil {
			cmds = append(cmds, c)
		}
	}

	return tea.Batch(cmds...)
}

// Updates for the applicationLayoutModel are delegated to the FocusController
func (m *applicationLayoutModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *applicationLayoutModel) View() string {
	rendered := []string{}
	for _, mod := range m.Models {
		rendered = append(rendered, mod.View())
	}
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		rendered...,
	)
}

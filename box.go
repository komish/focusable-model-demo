package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func PlainText(t string) *styledText {
	return styleText(t, lipgloss.NewStyle())
}

func styleText(text string, style lipgloss.Style) *styledText {
	return &styledText{
		s: style,
		t: text,
	}
}

type styledText struct {
	s lipgloss.Style
	t string
}

func (s *styledText) String() string {
	return s.s.Render(s.t)
}

type Box struct {
	raw string
	// Using a stringer here seems less useful. Might be better to use a view.
	text            fmt.Stringer
	unselectedStyle lipgloss.Style
	selectedStyle   lipgloss.Style
	isSelected      bool
}

type BoxOption func(*Box)

func IsSelected(b *Box) {
	b.isSelected = true
}

func NewBox(text string, opts ...BoxOption) *Box {
	baseStyle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		Padding(0, 4)

	selectedStyle := baseStyle.Copy()
	selectedStyle.Background(lipgloss.AdaptiveColor{
		Light: "#000",
		Dark:  "#fff",
	}).BorderBackground(lipgloss.AdaptiveColor{
		Light: "#000",
		Dark:  "#fff",
	})

	t := "BoxPlaceholderText"
	if text != "" {
		t = text
	}

	return &Box{
		raw:             t,
		text:            PlainText(t),
		isSelected:      false,
		unselectedStyle: baseStyle,
		selectedStyle:   selectedStyle,
	}
}

func (b *Box) Select() {
	b.isSelected = true
}

func (b *Box) IsSelected() *Box {
	b.Select()
	return b
}

// String returns the string representation of Box.
func (b *Box) String() string {
	s := b.unselectedStyle
	if b.isSelected {
		s = b.selectedStyle
	}

	return s.Render(b.text.String())
}

// Raw returns the raw value with no style.
func (b *Box) Raw() string {
	return b.raw
}

// Render is an alias for String
func (b *Box) Render() string { return b.String() }
func (b *Box) View() string   { return b.String() }

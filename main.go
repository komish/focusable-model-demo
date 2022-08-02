package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	displayOnlyBoxKey = "selectedValue"
)

func main() {
	// These will be reachable by the user.
	interactiveComponent1 := FocusableBoxArray("1", "2", "3", "4", layout1x4, displayOnlyBoxKey)
	interactiveComponent2 := FocusableBoxArray("A", "B", "C", "D", layout2x2, displayOnlyBoxKey)
	interactiveComponent3 := FocusableBoxArray("!", "@", "#", "$", layout4x1, displayOnlyBoxKey)

	// This will not be reachable, even though it's focusable.
	nonInteractiveComponent := FocusableBoxArray("H", "J", "K", "L", layout2x2, displayOnlyBoxKey)

	// This displays the value the user selected with Enter
	// in any of the interactive components above.
	selectedValueBox := ViperBox(displayOnlyBoxKey)

	appModel := FocusController{
		// Anything that's being rendered to the user must be in the layout.
		Layout: &applicationLayoutModel{
			Models: []tea.Model{
				selectedValueBox,
				interactiveComponent1,
				interactiveComponent2,
				interactiveComponent3,
				nonInteractiveComponent,
			},
		},
		// The items that can receive focus using the arrow keys must be here
		// in FocusableComponent
		FocusableComponent: []Focusable{
			interactiveComponent1,
			interactiveComponent2,
			interactiveComponent3,
			// leaving the nonInteractiveComponent prevents it
			// from receiving focus.
		},
	}

	p := tea.NewProgram(&appModel)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

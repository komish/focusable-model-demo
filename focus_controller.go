package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Viewable contains the View tea.Model method. It's for
// components that do need Updates or Initialization in the
// traditional sense of a tea.Model
type Viewable interface {
	View() string
}

// FocusController is a Model that wraps a model which may contain multiple resources
// that can attain focus.
//
// Components that have individual focus points must be registered as a FocusableComponent.
// Left and Right tea.Msgs are captured to advance Focus to individual items in each FocusableComponent.
// Other tea.Msgs are passed directly to the component in focus for handling.
// FocusableComponent methods must be bound to pointer receivers on implementations in order for
// the Layout to render appropriately.
//
// Layout represents the View components of a tea.Model. This view should include FocusableComponents, or other components that don't have focus elements.
// As FocusableComponents receive focus events, or tea.Msgs, theyir views are updated however they see fit.
// The Layout is then rendered when a FocusController's View method is called.
type FocusController struct {
	// Layout contains the visual aspects of what needs to be rendered.
	Layout Viewable
	// Model tea.Model
	// FocusableComponent contains components that are in the focus queue.
	FocusableComponent []Focusable
	// ComponentInFocus is the index of the currently focused component.
	ComponentInFocus int
}

// FocusMsg is a response to a FocusNext or FocusPrevious request.
type FocusMsg = string

const FocusMsgNoMoreItems = "OutOfFocusableItems"
const FocusMsgOK = "OK"

type Focusable interface {
	tea.Model
	FocusNext() FocusMsg
	FocusPrevious() FocusMsg
	ReceiveFocusFromStart()
	ReceiveFocusFromEnd()
	RemoveFocus()
}

func (m *FocusController) Init() tea.Cmd {
	// initialize the focus
	m.ComponentInFocus = 0
	m.FocusableComponent[0].ReceiveFocusFromStart()
	return nil
}

func (m *FocusController) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyRight:
			switch m.FocusableComponent[m.ComponentInFocus].FocusNext() {
			case FocusMsgNoMoreItems:
				// If we have another focusable component.
				if m.ComponentInFocus < len(m.FocusableComponent)-1 {
					m.FocusableComponent[m.ComponentInFocus].RemoveFocus()
					m.ComponentInFocus += 1
					m.FocusableComponent[m.ComponentInFocus].ReceiveFocusFromStart()
				}
				// otherwise, we can't do anything.
			}
		case tea.KeyLeft:
			switch m.FocusableComponent[m.ComponentInFocus].FocusPrevious() {
			case FocusMsgNoMoreItems:
				// If we have another focusable component.
				if m.ComponentInFocus != 0 {
					m.FocusableComponent[m.ComponentInFocus].RemoveFocus()
					m.ComponentInFocus -= 1
					m.FocusableComponent[m.ComponentInFocus].ReceiveFocusFromEnd()
				}
				// otherwise, we can't do anything.
			}
		default:
			// All other tea.Msg are passed to the component in focus.
			m.FocusableComponent[m.ComponentInFocus].Update(msg)
		}
	}

	return m, nil
}

func (m *FocusController) View() string {
	return m.Layout.View()
}

package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

type arrayLayout = string

const (
	layout2x2 = "2x2"
	layout1x4 = "1x4"
	layout4x1 = "4x1"
)

// FocusableBoxArray is a 2x2 Box Model that implements Focusable and also updates
// Viper at selectedKey when a value is chosen.
func FocusableBoxArray(t1, t2, t3, t4 string, layout arrayLayout, selectKey string) *focusableBoxArray {
	return &focusableBoxArray{
		boxes: []Box{
			*NewBox(t1),
			*NewBox(t2),
			*NewBox(t3),
			*NewBox(t4),
		},
		selected:  nil,
		selectKey: selectKey,
		layout:    layout,
	}
}

type focusableBoxArray struct {
	boxes     []Box
	selected  *int
	selectKey string
	layout    arrayLayout
}

func (b *focusableBoxArray) Init() tea.Cmd {
	return nil
}

func (b *focusableBoxArray) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return b, tea.Quit
		case tea.KeyRight:
			// if *b.selected < len(b.boxes)-1 {
			// 	*b.selected += 1
			// }
			fmt.Println("Reached Inner KeyRight")
			b.FocusNext()
		case tea.KeyLeft:
			fmt.Println("Reached Inner KeyLeft")
			b.FocusPrevious()
		case tea.KeyDown:
			if *b.selected+2 <= len(b.boxes)-1 {
				*b.selected += 2
			}
		case tea.KeyUp:
			if *b.selected-2 >= 0 {
				*b.selected -= 2
			}
		case tea.KeyEnter:
			viper.Set(b.selectKey, b.boxes[*b.selected].Raw())
		}
	}

	return b, nil
}

func (b *focusableBoxArray) View() string {
	return b.render()
}

func (b *focusableBoxArray) render() string {
	// Make a copy of the boxes so that we can avoid having to reset
	// their selected states.
	toRender := make([]Box, len(b.boxes))
	copy(toRender, b.boxes)
	if b.selected != nil {
		toRender[*b.selected].Select()
	}

	switch b.layout {
	case layout2x2:
		return b.as2x2(toRender[0], toRender[1], toRender[2], toRender[3])
	case layout1x4:
		return b.as1x4(toRender[0], toRender[1], toRender[2], toRender[3])
	case layout4x1:
		return b.as4x1(toRender[0], toRender[1], toRender[2], toRender[3])
	default:
		// default 2x2
		return lipgloss.JoinVertical(lipgloss.Top,
			lipgloss.JoinHorizontal(lipgloss.Top, toRender[0].Render(), toRender[1].Render()),
			lipgloss.JoinHorizontal(lipgloss.Top, toRender[2].Render(), toRender[3].Render()),
		)
	}
}

func (b *focusableBoxArray) as2x2(b1, b2, b3, b4 Box) string {
	return lipgloss.JoinVertical(lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Top, b1.Render(), b2.Render()),
		lipgloss.JoinHorizontal(lipgloss.Top, b3.Render(), b4.Render()),
	)
}

func (b *focusableBoxArray) as1x4(b1, b2, b3, b4 Box) string {
	return lipgloss.JoinVertical(lipgloss.Top,
		b1.Render(),
		b2.Render(),
		b3.Render(),
		b4.Render(),
	)
}

func (b *focusableBoxArray) as4x1(b1, b2, b3, b4 Box) string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		b1.Render(),
		b2.Render(),
		b3.Render(),
		b4.Render(),
	)
}

func (b *focusableBoxArray) FocusNext() FocusMsg {
	if *b.selected < len(b.boxes)-1 {
		*b.selected += 1
		return FocusMsgOK
	}

	return FocusMsgNoMoreItems
}

func (b *focusableBoxArray) FocusPrevious() FocusMsg {
	if *b.selected > 0 {
		*b.selected -= 1
		return FocusMsgOK
	}

	return FocusMsgNoMoreItems
}
func (b *focusableBoxArray) ReceiveFocusFromStart() {
	b.selected = new(int)
}
func (b *focusableBoxArray) ReceiveFocusFromEnd() {
	e := 3 // downsides of selected being an int pointer.
	b.selected = &e
}
func (b *focusableBoxArray) RemoveFocus() {
	b.selected = nil
}

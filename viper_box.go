package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

// ViperBox is a box model that gets it value from viper at the
// provided key.
func ViperBox(key string) *viperBox {
	v := viper.GetString(key)
	if len(v) == 0 {
		viper.Set(key, "?")
	}

	return &viperBox{key: key}
}

type viperBox struct {
	key string
}

func (b *viperBox) Init() tea.Cmd {
	return nil
}

func (b *viperBox) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return b, nil
}

func (b *viperBox) View() string {
	return b.render()
}

func (b *viperBox) render() string {
	return NewBox(viper.GetString(b.key)).Render()
}

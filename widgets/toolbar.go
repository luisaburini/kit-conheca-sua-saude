package widgets

import (
	"conheca/sua/saude/controllers"
	"conheca/sua/saude/resources"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const (
	Speak = iota
	ShowHome
	ClearEntry
	ShowCollection
)

type ButtonInfo struct {
	Text     string
	Icon     fyne.Resource
	OnTapped func()
}

type Toolbar struct {
	state binding.Int
	view  *fyne.Container
}

func NewToolbar() *Toolbar {
	t := &Toolbar{
		state: binding.NewInt(),
		view:  container.NewHBox(),
	}
	return t
}

func (t *Toolbar) GetView(callbacks []func(), state controllers.State) *fyne.Container {
	fmt.Println("Callbacks len " + fmt.Sprint(len(callbacks)))
	t.view.RemoveAll()
	for _, buttonInfo := range getButtonInfo(callbacks) {
		t.view.Add(widget.NewButtonWithIcon(buttonInfo.Text,
			buttonInfo.Icon,
			buttonInfo.OnTapped))
	}
	t.ChangeState(state)
	return t.view
}

func (t *Toolbar) Hide(button int) {
	t.view.Objects[button].Hide()
}

func (t *Toolbar) Show(button int) {
	t.view.Objects[button].Show()
}

func (t *Toolbar) ChangeState(s controllers.State) {
	switch s {
	case controllers.Collection:
		log.Println("Going to show collection view")
		t.Hide(Speak)
		t.Show(ShowHome)
		t.Hide(ClearEntry)
		t.Hide(ShowCollection)
	case controllers.Home:
		log.Println("Going to show home")
		t.Show(Speak)
		t.Hide(ShowHome)
		t.Show(ClearEntry)
		t.Show(ShowCollection)
	}
}

func getButtonInfo(callbacks []func()) []ButtonInfo {
	return []ButtonInfo{
		{
			Text:     "Falar",
			Icon:     resources.IconSpeak(),
			OnTapped: callbacks[Speak],
		},
		{
			Text:     "Voltar",
			Icon:     resources.IconBack(),
			OnTapped: callbacks[ShowHome],
		},
		{
			Text:     "Limpar",
			Icon:     resources.IconClear(),
			OnTapped: callbacks[ClearEntry],
		},
		{
			Text:     "Acervo",
			Icon:     resources.IconDatabase(),
			OnTapped: callbacks[ShowCollection],
		},
	}
}

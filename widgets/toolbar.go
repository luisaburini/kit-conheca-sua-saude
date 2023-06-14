package widgets

import (
	"conheca/sua/saude/resources"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const (
	Speak           = 0
	PersistWords    = 1
	ShowMySentences = 2
	ShowHome        = 3
	ClearEntry      = 4
	ShowCollection  = 5
	SaveCollection  = 6
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

func (t *Toolbar) GetView(callbacks []func(), state State) *fyne.Container {
	for _, buttonInfo := range getButtonInfo(callbacks) {
		t.view.Add(widget.NewButtonWithIcon(buttonInfo.Text,
			buttonInfo.Icon,
			buttonInfo.OnTapped))
	}
	return t.view
}

func (t *Toolbar) Hide(button int) {
	t.view.Objects[button].Hide()
}

func (t *Toolbar) Show(button int) {
	t.view.Objects[button].Show()
}

func getButtonInfo(callbacks []func()) []ButtonInfo {
	return []ButtonInfo{
		{
			Text:     "Falar",
			Icon:     resources.IconSpeak(),
			OnTapped: callbacks[Speak],
		},
		{
			Text: "Salvar",
			Icon: resources.IconSave(),
			OnTapped: func() {
				callbacks[PersistWords]()
				callbacks[ClearEntry]()
			},
		},
		{
			Text:     "Frases",
			Icon:     resources.IconMySentences(),
			OnTapped: callbacks[ShowMySentences],
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
		{
			Text: "Salvar",
			Icon: resources.IconSave(),
			OnTapped: func() {
				callbacks[SaveCollection]()
			},
		},
	}
}

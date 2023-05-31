package widgets

import (
	"conheca/sua/saude/resources"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type HToolbar struct {
}

//0 speak func(),
//1 persistWords func(),
//2 showMyWords func(),
//3 clearEntry func()

const (
	Speak        = 0
	PersistWords = 1
	ShowMyWords  = 2
	ShowHome     = 2
	ClearEntry   = 3
)

func GetVToolbar(callbacks []func(), state State) *fyne.Container {
	speakButton := widget.NewButtonWithIcon("Falar", resources.IconSpeak(), func() {
		callbacks[Speak]()
	})
	saveButton := widget.NewButtonWithIcon("Salvar", resources.IconSave(), func() {
		callbacks[PersistWords]()
		callbacks[ClearEntry]()
	})
	var button *widget.Button
	switch state {
	case Home:
		button = widget.NewButtonWithIcon("Frases", resources.IconMySentences(), func() {
			callbacks[ShowMyWords]()
		})
	case SentenceList:
		button = widget.NewButtonWithIcon("Voltar", resources.IconBack(), func() {
			callbacks[ShowHome]()
		})
	}

	clearButton := widget.NewButtonWithIcon("Limpar", resources.IconClear(), func() {
		callbacks[ClearEntry]()
	})
	return container.NewGridWithColumns(1, speakButton, saveButton, button, clearButton)
}

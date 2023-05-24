package widgets

import (
	"conheca/sua/saude/resources"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type HToolbar struct {
}

func GetVToolbar(persistWords func(), showMyWords func()) *fyne.Container {
	saveButton := NewPictogram(resources.GetIconSave(), "Salvar")
	saveButton.SetOnTapped(func() {
		persistWords()
	})
	myWordsButton := NewPictogram(resources.GetIconMinhasFrases(), "Minhas Frases")
	myWordsButton.SetOnTapped(func() {
		showMyWords()
	})
	return container.NewAdaptiveGrid(1, saveButton.GetView(), myWordsButton.GetView())
}

package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type HToolbar struct {
}

func GetVToolbar(persistWords func(), showMyWords func()) *fyne.Container {
	saveButton := NewPictogram(GetIconSave(), "Salvar")
	saveButton.SetOnTapped(func() {
		persistWords()
	})
	myWordsButton := NewPictogram(GetIconMinhasFrases(), "Minhas Frases")
	myWordsButton.SetOnTapped(func() {
		// SHOW SCREEN
		showMyWords()
	})
	return container.NewAdaptiveGrid(1, saveButton.GetView(), myWordsButton.GetView())
}

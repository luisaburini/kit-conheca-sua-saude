package widgets

import (
	"conheca/sua/saude/resources"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type HToolbar struct {
}

func GetVToolbar(persistWords func(), showMyWords func(), clearEntry func()) *fyne.Container {
	saveButton := NewPictogram(resources.IconSave(), "Salvar")
	saveButton.SetOnTapped(func() {
		persistWords()
		clearEntry()
	})
	myWordsButton := NewPictogram(resources.IconMySentences(), "Minhas Frases")
	myWordsButton.SetOnTapped(func() {
		showMyWords()
	})
	clearButton := NewPictogram(resources.IconClear(), "Limpar")
	clearButton.SetOnTapped(func() {
		clearEntry()
	})
	return container.NewVBox(saveButton.GetView(), myWordsButton.GetView(), clearButton.GetView())
}

package widgets

import (
	"conheca/sua/saude/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MyWords struct {
}

func GetMyWordsView(onTapped func(string), database *storage.Database) *fyne.Container {
	savedWords := database.GetSentences()

	return container.NewMax(widget.NewList(
		// Length
		func() int {
			return len(savedWords)
		},
		// Create
		func() fyne.CanvasObject {
			return widget.NewButton("", func() {})
		},
		// Update
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			button, ok := co.(*widget.Button)
			if !ok {
				return
			}
			button.SetText(savedWords[lii])
			button.OnTapped = func() {
				onTapped(savedWords[lii])
			}
		}))
}

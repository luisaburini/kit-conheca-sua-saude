package widgets

import (
	"conheca/sua/saude/storage"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MySentences struct {
	wordsList  *widget.List
	savedWords []string
	database   *storage.Database
}

func NewMySentences(onTapped func(string), database *storage.Database) *MySentences {
	var mySentences *MySentences
	mySentences = &MySentences{
		database: database,
		wordsList: widget.NewList(
			// Length
			func() int {
				return len(mySentences.savedWords)
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
				button.SetText(mySentences.savedWords[lii])
				button.OnTapped = func() {
					onTapped(mySentences.savedWords[lii])
				}
			}),
	}
	mySentences.savedWords = database.GetSentences()
	return mySentences
}

func (mw *MySentences) GetView() *fyne.Container {
	return container.New(layout.NewPaddedLayout(), mw.wordsList)
}

func (mw *MySentences) Refresh() {
	if mw != nil {
		mw.savedWords = mw.database.GetSentences()
		mw.wordsList.Refresh()
	} else {
		log.Println("My Words is NULL")
	}
}

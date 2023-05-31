package widgets

import (
	"conheca/sua/saude/storage"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MyWords struct {
	wordsList  *widget.List
	savedWords []string
	database   *storage.Database
}

func NewMyWords(onTapped func(string), database *storage.Database) *MyWords {
	var myWords *MyWords
	myWords = &MyWords{
		database: database,
		wordsList: widget.NewList(
			// Length
			func() int {
				return len(myWords.savedWords)
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
				button.SetText(myWords.savedWords[lii])
				button.OnTapped = func() {
					onTapped(myWords.savedWords[lii])
				}
			}),
	}
	myWords.savedWords = database.GetSentences()
	return myWords
}

func (mw *MyWords) GetView() *fyne.Container {
	return container.New(layout.NewPaddedLayout(), mw.wordsList)
}

func (mw *MyWords) Refresh() {
	if mw != nil {
		mw.savedWords = mw.database.GetSentences()
		mw.wordsList.Refresh()
	} else {
		log.Println("My Words is NULL")
	}
}

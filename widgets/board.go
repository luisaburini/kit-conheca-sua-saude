package widgets

import (
	"conheca/sua/saude/audio"
	"conheca/sua/saude/resources"
	"conheca/sua/saude/storage"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetBoardView(w fyne.Window, database *storage.Database) *fyne.Container {
	var screen *fyne.Container
	grid := container.NewAdaptiveGrid(6)
	texts := GetWords()
	speakEntry := widget.NewEntry()
	for i, resource := range resources.Words() {
		pictogram := NewPictogram(resource, texts[i])
		pictogram.SetOnTapped(func() {
			speakEntry.Text = speakEntry.Text + " " + pictogram.GetWord()
			speakEntry.Refresh()
		})
		grid.Add(pictogram.GetView())
	}
	speakButton := widget.NewButtonWithIcon("Falar", resources.GetIconFalar(), func() {
		audio.Play(speakEntry.Text, w)
	})
	clearButton := widget.NewButtonWithIcon("Limpar", resources)
	top := container.New(layout.NewFormLayout(), speakButton, speakEntry)
	var left *fyne.Container
	left = GetVToolbar(func() {
		log.Println("Persist!")
		err := database.AddSentence(speakEntry.Text)
		if err != nil {
			log.Println("Add sentence " + err.Error())
		}
	}, func() {
		log.Println("Show My Words screen")
		showMyWordsScreen(speakEntry, top, left, grid, screen, database)
	})
	screen = container.NewBorder(top, nil, left, nil, grid)
	return screen
}

func showMyWordsScreen(speakEntry *widget.Entry, top, left, grid, screen *fyne.Container, database *storage.Database) {
	screen.RemoveAll()
	screen.Add(GetMyWordsView(func(s string) {
		screen.RemoveAll()
		speakEntry.SetText(s)
		screen.Add(grid)
		screen.Add(top)
		screen.Add(left)
	}, database))
	screen.Add(top)
	screen.Add(left)
}

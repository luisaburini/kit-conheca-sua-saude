package widgets

import (
	"conheca/sua/saude/audio"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetBoardView(w fyne.Window) *fyne.Container {
	grid := container.NewAdaptiveGrid(6)
	texts := GetWords()
	speakEntry := widget.NewEntry()
	for i, resource := range GetResources() {
		pictogram := NewPictogram(resource, texts[i])
		pictogram.SetOnTapped(func() {
			speakEntry.Text = speakEntry.Text + " " + pictogram.GetWord()
			speakEntry.Refresh()
		})
		grid.Add(pictogram.GetView())
	}

	speakButton := widget.NewButtonWithIcon("Falar", GetIconFalar(), func() {
		audio.Play(speakEntry.Text, w)
	})
	top := container.New(layout.NewFormLayout(), speakButton, speakEntry)
	left := GetVToolbar(func() {
		log.Println("Persist!")
	}, func() {
		log.Println("Show My Words screen")
	})

	return container.NewBorder(top, nil, left, nil, grid)
}

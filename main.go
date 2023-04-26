package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"golang.org/x/mobile/exp/audio/al"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func main() {
	a := app.New()
	w := a.NewWindow("Kit Conheca Sua Saude")

	files, err := ioutil.ReadDir("./assets/icons/")
	if err != nil {
		log.Fatal(err)
	}
	grid := container.NewAdaptiveGrid(3)
	for _, file := range files {
		grid.Add(widget.NewCard("", fileNameWithoutExtSliceNotation(file.Name()), canvas.NewImageFromFile("./assets/icons/"+file.Name())))
	}
	w.SetContent(container.NewMax(grid))
	w.ShowAndRun()
}

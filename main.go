package main

import (
	"conheca/sua/saude/widgets"
	"path/filepath"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func main() {
	a := app.New()
	w := a.NewWindow("Kit Conheca Sua Saude")

	grid := container.NewAdaptiveGrid(4)

	texts := []string{
		"Anticoncepcional",
		"Boca",
		"Câncer de Pênis",
		"Câncer de Vulva",
		"Coração",
		"Dedo",
		"Gonorréia",
		"Gravidez",
		"Saúde Mental",
		"Saúde",
		"Sexo",
		"Vulva",
	}

	for i, resource := range widgets.GetResources() {
		grid.Add(widgets.NewPictogram(resource, texts[i]))
	}
	w.SetContent(container.NewMax(grid))
	w.ShowAndRun()
}

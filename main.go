package main

import (
	"conheca/sua/saude/widgets"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Kit Conheca Sua Saude")

	w.SetContent(widgets.GetBoardView(w))
	w.ShowAndRun()
}

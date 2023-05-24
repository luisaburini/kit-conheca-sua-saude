package main

import (
	"conheca/sua/saude/storage"
	"conheca/sua/saude/widgets"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Kit Conheca Sua Saude")
	database := storage.NewDatabase()
	w.SetContent(widgets.GetBoardView(w, database))
	w.ShowAndRun()
	database.Close()
}

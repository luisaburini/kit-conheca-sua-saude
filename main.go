package main

import (
	"conheca/sua/saude/storage"
	"conheca/sua/saude/widgets"

	"fyne.io/fyne/v2/app"
)

func main() {
	//os.Setenv("FYNE_SCALE", "1.8")
	a := app.New()
	w := a.NewWindow("Kit Conheca Sua Saude")
	database := storage.NewDatabase()
	defer database.Close()
	board := widgets.NewBoard(w)
	w.SetContent(board.GetView(database))
	w.ShowAndRun()
}

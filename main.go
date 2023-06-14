package main

import (
	"conheca/sua/saude/storage"
	"conheca/sua/saude/widgets"

	"fyne.io/fyne/v2/app"
	//_ "golang.org/x/mobile/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("Kit Conheca Sua Saude")
	database := storage.NewDatabase()
	defer database.Close()
	board := widgets.NewBoard(database, window)
	window.SetContent(board.GetView())
	window.ShowAndRun()
}

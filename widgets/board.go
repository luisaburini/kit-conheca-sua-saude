package widgets

import (
	"conheca/sua/saude/audio"
	"conheca/sua/saude/resources"
	"conheca/sua/saude/storage"
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type State uint8

const (
	Home         State = 0
	SentenceList State = 1
)

type Board struct {
	state State

	speakEntry *widget.Entry

	screen  *fyne.Container
	vscroll *fyne.Container
	window  fyne.Window
	myWords *MyWords
}

func NewBoard(window fyne.Window) *Board {
	board := &Board{
		speakEntry: widget.NewEntry(),
		window:     window,
	}
	board.vscroll = board.initVScroll()
	board.speakEntry.MultiLine = true
	board.speakEntry.Wrapping = fyne.TextWrapWord
	board.screen = container.NewGridWithColumns(1)
	return board
}

func (b *Board) GetView(database *storage.Database) *fyne.Container {
	b.myWords = NewMyWords(func(s string) {
		b.speakEntry.SetText(s)
	}, database)
	b.showHome(database)
	return container.NewMax(b.screen)
}

func (b *Board) showScreen(database *storage.Database) {
	fmt.Println(b.state)
	b.screen.RemoveAll()
	toolbar := GetVToolbar(
		[]func(){
			func() {
				fmt.Println("Clicked Speak")
				audio.Play(b.speakEntry.Text, b.window)
			},
			func() {
				log.Println("Clicked Persist")
				b.persist(database)
			}, func() {
				switch b.state {
				case Home:
					log.Println("Show my Words")
					b.showMyWordsScreen(database)
				case SentenceList:
					log.Println("Show Home")
					b.showHome(database)
				}
			}, func() {
				log.Println("Clicked Clear")
				b.clearSpeakEntry()
			},
		}, b.state)
	b.screen.Add(container.NewPadded(b.speakEntry))
	b.screen.Add(toolbar)
	switch b.state {
	case Home:
		b.screen.Add(container.NewCenter(container.NewGridWrap(fyne.NewSize(4*iconWidth, 3*iconHeight),
			b.vscroll)))
	case SentenceList:
		b.screen.Add(b.myWords.GetView())
	}
	b.screen.Refresh()
}

func (b *Board) showHome(database *storage.Database) {
	log.Println("showHome")
	b.state = Home
	b.showScreen(database)
}

func (b *Board) showMyWordsScreen(database *storage.Database) {
	log.Println("showMyWordsScreen")
	b.state = SentenceList
	b.showScreen(database)
}

func (b *Board) persist(database *storage.Database) {
	err := database.AddSentence(b.speakEntry.Text)
	if err != nil {
		log.Println("Add sentence " + err.Error())
	}
}

func (b *Board) remove(database *storage.Database) {
	log.Println("Entered remove")
	log.Println("State " + fmt.Sprint(b.state))
	if b.state == SentenceList {
		err := database.RemoveSentence(b.speakEntry.Text)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (b *Board) clearSpeakEntry() {
	log.Println("Entered clear speak entry")
	b.speakEntry.Text = ""
	b.speakEntry.Refresh()
	b.myWords.Refresh()
}

const (
	iconWidth  = 100
	iconHeight = 100
)

func (b *Board) initVScroll() *fyne.Container {
	texts := GetWords()
	textResources := resources.Words()
	// for i, resource := range resources.Words() {
	// 	var pictogram *Pictogram
	// 	pictogram = NewPictogram(texts[i], resource, func() {
	// 		log.Println("Tapped " + pictogram.word)
	// 		b.speakEntry.Text = b.speakEntry.Text + " " + pictogram.GetWord()
	// 		b.speakEntry.Refresh()
	// 	})
	// 	grid.Add(pictogram.GetView())
	// }
	var list *widget.List
	list = widget.NewList(
		func() int {
			return len(texts)
		},
		func() fyne.CanvasObject {
			label := canvas.NewText("", color.Black)
			label.TextSize = 15
			label.Alignment = fyne.TextAlignCenter
			//icon := widget.NewIcon(resources.IconClear())
			icon := canvas.NewImageFromResource(resources.IconClear())
			icon.FillMode = canvas.ImageFillContain
			icon.Resize(fyne.NewSize(iconWidth, iconHeight))
			icon.Refresh()
			return container.NewGridWithColumns(1,
				container.NewCenter(label),
				container.NewMax(container.NewCenter(container.NewGridWrap(fyne.NewSize(iconWidth, iconHeight),
					icon))))
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			label, ok := co.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*canvas.Text)
			if !ok {
				return
			}
			//icon, ok := co.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*widget.Icon)
			icon, ok := co.(*fyne.Container).Objects[1].(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*canvas.Image)
			if !ok {
				return
			}
			label.Text = texts[lii]
			label.Refresh()
			icon.Resource = textResources[lii]
			icon.Refresh()
		})
	list.OnSelected = func(id widget.ListItemID) {
		log.Println(texts[id])
		b.speakEntry.Text = b.speakEntry.Text + " " + texts[id]
		b.speakEntry.Refresh()
	}
	return container.NewMax(list)
}

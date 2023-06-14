package widgets

import (
	"conheca/sua/saude/audio"
	"conheca/sua/saude/resources"
	"conheca/sua/saude/storage"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type State uint8

const (
	Home         State = 0
	SentenceList State = 1
	Collection   State = 2
)

type Board struct {
	state State

	speakEntry *widget.Entry
	database   *storage.Database
	screen     *fyne.Container
	vscroll    *fyne.Container
	window     fyne.Window
	myWords    *MySentences
	collection *CollectionView
}

func NewBoard(database *storage.Database, window fyne.Window) *Board {
	board := &Board{
		speakEntry: widget.NewEntry(),
		window:     window,
		database:   database,
		collection: NewCollectionView(),
	}
	board.vscroll = board.initVScroll()
	board.speakEntry.MultiLine = true
	board.speakEntry.Wrapping = fyne.TextWrapWord
	board.screen = container.NewMax()
	return board
}

func (b *Board) GetView() *fyne.Container {
	b.myWords = NewMySentences(func(s string) {
		b.speakEntry.SetText(s)
	}, b.database)
	b.showCollection()
	return b.screen
}

func (b *Board) showScreen() {
	fmt.Println(b.state)
	b.screen.RemoveAll()
	toolbar := NewToolbar()
	toolbarView := toolbar.GetView(
		[]func(){
			func() { // Speak
				audio.Play(b.speakEntry.Text, b.window)
				//audio.Play()
			},
			b.persist,               // PersistWords
			b.showMySentencesScreen, // ShowMySentences
			b.showHome,              // ShowHome
			b.clearSpeakEntry,       // ClearEntry
			b.showCollection,        // ShowCollection
			func() {},
		}, b.state)

	switch b.state {
	case Collection:
		log.Println("Going to show collection view")
		toolbar.Hide(Speak)
		toolbar.Hide(PersistWords)
		toolbar.Show(ShowMySentences)
		toolbar.Show(ShowHome)
		toolbar.Hide(ClearEntry)
		toolbar.Hide(ShowCollection)
		toolbar.Show(SaveCollection)
		b.screen.Add(container.NewBorder(toolbarView, nil, nil, nil, b.collection.GetView()))
	case Home:
		log.Println("Going to show home")
		toolbar.Show(Speak)
		toolbar.Show(PersistWords)
		toolbar.Show(ShowMySentences)
		toolbar.Hide(ShowHome)
		toolbar.Show(ClearEntry)
		toolbar.Show(ShowCollection)
		toolbar.Hide(SaveCollection)
		b.screen.Add(container.NewBorder(container.NewVBox(toolbarView, container.NewPadded(b.speakEntry)),
			nil, nil, nil,
			b.initVScroll()))
	case SentenceList:
		log.Println("Going to show sentence list")
		toolbar.Show(Speak)
		toolbar.Show(PersistWords)
		toolbar.Hide(ShowMySentences)
		toolbar.Show(ShowHome)
		toolbar.Hide(ClearEntry)
		toolbar.Hide(ShowCollection)
		toolbar.Hide(SaveCollection)
		b.screen.Add(container.NewBorder(container.NewVBox(toolbarView, container.NewPadded(b.speakEntry)),
			nil, nil, nil,
			b.myWords.GetView()))
	}
	b.screen.Refresh()
}

func (b *Board) showHome() {
	log.Println("showHome")
	b.state = Home
	b.showScreen()
}

func (b *Board) showMySentencesScreen() {
	log.Println("showMySentencesScreen")
	b.state = SentenceList
	b.showScreen()
}

func (b *Board) showCollection() {
	log.Println("showCollection")
	b.state = Collection
	b.showScreen()
}

func (b *Board) persist() {
	err := b.database.AddSentence(b.speakEntry.Text)
	if err != nil {
		log.Println("Add sentence " + err.Error())
	}
}

func (b *Board) clearSpeakEntry() {
	b.speakEntry.Text = ""
	b.speakEntry.Refresh()
	b.myWords.Refresh()
}

func (b *Board) initVScroll() *fyne.Container {
	texts := b.getBoardTexts()
	collection := resources.Collection()
	screen := container.NewAdaptiveGrid(2)
	for _, text := range texts {
		fmt.Println(text)
		pictogram := NewPictogram(text, collection[text])
		pictogram.SetOnTapped(func() {
			b.speakEntry.Text = b.speakEntry.Text + " " + pictogram.Word
			b.speakEntry.Refresh()
		})
		screen.Add(pictogram.GetView())
	}
	return screen
}

func (b *Board) getBoardTexts() []string {
	var texts []string
	for _, p := range b.collection.GetSelected() {
		fmt.Println(p.Word)
		texts = append(texts, p.Word)
	}
	return texts
}

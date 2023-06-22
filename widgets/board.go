package widgets

import (
	"conheca/sua/saude/audio"
	"conheca/sua/saude/controllers"
	"conheca/sua/saude/resources"
	"conheca/sua/saude/storage"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Board struct {
	stateManager *controllers.StateManager
	toolbar      *Toolbar
	speakEntry   *widget.Entry
	database     *storage.Database
	screen       *fyne.Container
	vscroll      *fyne.Container
	window       fyne.Window
	myWords      *MySentences
	collection   *CollectionView
}

func NewBoard(database *storage.Database, window fyne.Window) *Board {
	board := &Board{
		speakEntry: widget.NewEntry(),
		window:     window,
		database:   database,
		collection: NewCollectionView(),
		toolbar:    NewToolbar(),
	}
	board.vscroll = board.initVScroll()
	board.speakEntry.MultiLine = true
	board.speakEntry.Wrapping = fyne.TextWrapWord
	board.stateManager = controllers.NewStateManager(board.onStateChange)
	board.screen = container.NewMax(board.toolbar.GetView(board.toolbarCallbacks(), board.stateManager.GetState()))
	return board
}

func (b *Board) toolbarCallbacks() []func() {
	return []func(){
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
	}
}

func (b *Board) GetView() *fyne.Container {
	b.myWords = NewMySentences(func(s string) {
		b.speakEntry.SetText(s)
	}, b.database)
	b.showCollection()
	return b.screen
}

func (b *Board) onStateChange() {
	state := b.stateManager.GetState()
	fmt.Println(state)
	b.toolbar.ChangeState(state)
	b.changeScreen(state)
}

func (b *Board) changeScreen(state controllers.State) {
	b.screen.RemoveAll()
	fmt.Println("Before Toolbar view " + fmt.Sprint(state))
	toolbarView := b.toolbar.GetView(b.toolbarCallbacks(), state)
	fmt.Println("After Toolbar view")
	switch state {
	case controllers.Collection:
		b.screen.Add(container.NewBorder(toolbarView, nil, nil, nil, b.collection.GetView()))
	case controllers.Home:
		b.screen.Add(container.NewBorder(container.NewVBox(toolbarView, container.NewPadded(b.speakEntry)),
			nil, nil, nil,
			b.initVScroll()))
	case controllers.SentenceList:
		b.screen.Add(container.NewBorder(container.NewVBox(toolbarView, container.NewPadded(b.speakEntry)),
			nil, nil, nil,
			b.myWords.GetView()))
	}
	b.screen.Refresh()
}

func (b *Board) showHome() {
	log.Println("showHome")
	b.stateManager.SetState(controllers.Home)
}

func (b *Board) showMySentencesScreen() {
	log.Println("showMySentencesScreen")
	b.stateManager.SetState(controllers.SentenceList)
}

func (b *Board) showCollection() {
	log.Println("showCollection")
	b.stateManager.SetState(controllers.Collection)
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

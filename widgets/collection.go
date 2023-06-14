package widgets

import (
	"conheca/sua/saude/resources"
	"fmt"
	"log"
	"sort"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

const maxPictograms = 6

type CollectionView struct {
	view       *fyne.Container
	pictograms map[string]*Pictogram
	selected   int
	mu         sync.Mutex
}

func NewCollectionView() *CollectionView {
	c := &CollectionView{
		view:       container.NewAdaptiveGrid(4),
		pictograms: make(map[string]*Pictogram),
	}
	collection := GetCollection()
	textResources := resources.Collection()
	if len(collection) != len(textResources) {
		c.view = container.NewMax()
	} else {
		for _, text := range collection {
			log.Println(text)
			pictogram := NewPictogram(text, textResources[text])
			pictogram.SetOnTapped(func() {
				c.onSelectPictogram(pictogram)
			})
			c.view.Add(pictogram.GetView())
		}
	}
	return c
}

func (c *CollectionView) onSelectPictogram(pictogram *Pictogram) {
	if pictogram.isSelected {
		c.selected--
		fmt.Println(pictogram.Label.Text + fmt.Sprint(c.selected))
		c.mu.Lock()
		delete(c.pictograms, pictogram.Word)
		for _, p := range c.pictograms {
			if p.isSelected {
				p.SetNumber(c.getSelectedIndex(p) + 1)
			}
		}
		pictogram.Select(c.selected)
		c.mu.Unlock()
		return
	}
	if c.selected < maxPictograms {
		c.selected++
		fmt.Println(pictogram.Label.Text + fmt.Sprint(c.selected))
		c.pictograms[pictogram.Word] = pictogram
		pictogram.Select(c.selected)
	}
}

func (c *CollectionView) getSelectedIndex(selected *Pictogram) int {
	var sortedWords []string
	for _, pictogram := range c.pictograms {
		sortedWords = append(sortedWords, pictogram.Word)
	}
	sort.Strings(sortedWords)
	for i, word := range sortedWords {
		if word == selected.Word {
			return i
		}
	}
	return -1
}

func (c *CollectionView) GetView() *fyne.Container {
	return c.view
}

func (c *CollectionView) GetSelected() map[string]*Pictogram {
	return c.pictograms
}

package widgets

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newTappableIcon(res fyne.Resource) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)
	return icon
}

type tappableIcon struct {
	widget.Icon
	onTapped func()
}

func (t *tappableIcon) Tapped(_ *fyne.PointEvent) {
	t.onTapped()
}

func (t *tappableIcon) TappedSecondary(_ *fyne.PointEvent) {
	t.onTapped()
}

const SQUARE_PERIMETER = 4

func NewPictogram(text string, resource fyne.Resource) *Pictogram {
	pictogram := &Pictogram{}
	pictogram.selectedNumber = canvas.NewText("", color.Black)
	pictogram.selectedNumber.TextSize = 15
	pictogram.selectedNumber.Hide()
	pictogram.Label = canvas.NewText(text, color.Black)
	pictogram.Label.TextSize = 15
	pictogram.Label.Alignment = fyne.TextAlignCenter
	pictogram.Icon = newTappableIcon(resource)
	pictogram.Word = text
	for i := 0; i < SQUARE_PERIMETER; i++ {
		pictogram.frame = append(pictogram.frame, canvas.NewRectangle(color.Black))
	}
	board := container.NewGridWithColumns(1, pictogram.Label, pictogram.Icon)
	pictogram.Container = container.NewBorder(nil, pictogram.frame[1],
		nil, pictogram.frame[3], pictogram.selectedNumber, board)
	pictogram.hideFrame()
	return pictogram
}

type Pictogram struct {
	Word           string
	Icon           *tappableIcon
	Label          *canvas.Text
	selectedNumber *canvas.Text
	frame          []*canvas.Rectangle
	Container      *fyne.Container
	isHidden       bool
	isSelected     bool
}

func (p *Pictogram) SetText(text string) {
	p.Label = canvas.NewText(text, color.Black)
	p.Label.TextSize = 15
	p.Label.Alignment = fyne.TextAlignCenter
	p.Word = text
}

func (p *Pictogram) SetResource(resource fyne.Resource) {
	p.Icon = newTappableIcon(resource)
}

func (p *Pictogram) SetOnTapped(onTapped func()) {
	p.Icon.onTapped = onTapped
}

func (p *Pictogram) GetView() *fyne.Container {
	return p.Container
}

func (p *Pictogram) SetNumber(n int) {
	p.selectedNumber.Text = strconv.Itoa(n)
	p.selectedNumber.Refresh()
}

func (p *Pictogram) Select(n int) {
	if p.isSelected {
		p.hideFrame()
		p.selectedNumber.Hide()
	} else {
		p.selectedNumber.Text = strconv.Itoa(n)
		p.selectedNumber.Show()
		p.showFrame()
	}
	p.isSelected = !p.isSelected
	p.Refresh()
}

func (p *Pictogram) hideFrame() {
	for i := 0; i < SQUARE_PERIMETER; i++ {
		p.frame[i].Hide()
	}
}

func (p *Pictogram) showFrame() {
	for i := 0; i < SQUARE_PERIMETER; i++ {
		p.frame[i].Show()
	}
}

func (p *Pictogram) GetWord() string {
	return p.Word
}

func (p *Pictogram) Hide() {
	p.Icon.Hide()
	p.Label.Hide()
	p.isHidden = true
}

func (p *Pictogram) Show() {
	p.Icon.Show()
	p.Label.Show()
	p.isHidden = false
}

func (p *Pictogram) Visible() bool {
	return !p.isHidden
}

func (p *Pictogram) MinSize() fyne.Size {
	return p.Container.MinSize()
}

func (p *Pictogram) Size() fyne.Size {
	return p.Container.Size()
}

func (p *Pictogram) Move(pos fyne.Position) {
	p.Container.Move(pos)
	p.Icon.Move(pos)
	p.Label.Move(pos)
}

func (p *Pictogram) Position() fyne.Position {
	return p.Container.Position()
}

func (p *Pictogram) Refresh() {
	p.Icon.Refresh()
	p.Label.Refresh()
	p.Container.Refresh()
	for i := 0; i < SQUARE_PERIMETER; i++ {
		p.frame[i].Refresh()
	}
	p.selectedNumber.Refresh()
}

func (p *Pictogram) Resize(size fyne.Size) {
	p.Container.Resize(size)
}

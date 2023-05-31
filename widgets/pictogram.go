package widgets

import (
	"image/color"

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

// func NewPictogram(text string, resource *fyne.StaticResource, onTapped func()) *Pictogram {
// 	pictogram := &Pictogram{}
// 	pictogram.label = canvas.NewText(text, color.Black)
// 	pictogram.label.TextSize = 15
// 	pictogram.label.Alignment = fyne.TextAlignCenter
// 	pictogram.icon = newTappableIcon(resource)
// 	pictogram.word = text
// 	pictogram.c = container.NewGridWithColumns(1, pictogram.label, pictogram.icon)
// 	pictogram.SetOnTapped(onTapped)
// 	return pictogram
// }

func NewPictogram() *Pictogram {
	p := &Pictogram{}
	p.Container = container.NewGridWithColumns(1, p.Label, p.Icon)
	return p
}

type Pictogram struct {
	word      string
	Icon      *tappableIcon
	Label     *canvas.Text
	Container *fyne.Container
	isHidden  bool
}

func (p *Pictogram) SetText(text string) {
	p.Label = canvas.NewText(text, color.Black)
	p.Label.TextSize = 15
	p.Label.Alignment = fyne.TextAlignCenter
	p.word = text
}

func (p *Pictogram) SetResource(resource *fyne.StaticResource) {
	p.Icon = newTappableIcon(resource)
}

func (p *Pictogram) SetOnTapped(onTapped func()) {
	p.Icon.onTapped = onTapped
}

func (p *Pictogram) GetView() *fyne.Container {
	return p.Container
}

func (p *Pictogram) GetWord() string {
	return p.word
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
}

func (p *Pictogram) Resize(size fyne.Size) {
	p.Container.Resize(size)
}

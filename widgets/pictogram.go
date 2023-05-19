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

func NewPictogram(resource *fyne.StaticResource, text string) *Pictogram {
	pictogram := &Pictogram{}
	pictogram.label = canvas.NewText(text, color.Black)
	pictogram.label.TextSize = 15
	pictogram.label.Alignment = fyne.TextAlignCenter
	pictogram.icon = newTappableIcon(resource)
	pictogram.word = text
	pictogram.c = container.NewGridWithColumns(1, pictogram.label, pictogram.icon)
	return pictogram
}

type Pictogram struct {
	word  string
	icon  *tappableIcon
	label *canvas.Text
	c     *fyne.Container
}

func (p *Pictogram) SetOnTapped(onTapped func()) {
	p.icon.onTapped = onTapped
}

func (p *Pictogram) GetView() *fyne.Container {
	return p.c
}

func (p *Pictogram) GetWord() string {
	return p.word
}

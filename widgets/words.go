package widgets

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
)

type Word struct {
	text *canvas.Text
}

func NewWord(content string, color color.Color) *Word {
	return &Word{
		text: canvas.NewText(content, color),
	}
}

func GetCollection() []string {
	return []string{
		"Anticoncepcional",
		"Boca",
		"Câncer de Pênis",
		"Câncer de Vulva",
		"Coração",
		"Dedo",
		"Gonorréia",
		"Gravidez",
		"Saúde Mental",
		"Saúde",
		"Sexo",
		"Vulva",
	}
}

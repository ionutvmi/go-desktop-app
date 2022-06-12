package activity

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type Activity interface {
	GetContent() *fyne.Container
}

type DefaultActivity struct {
	Id string
}

func (a *DefaultActivity) GetContent() *fyne.Container {
	mainContent := canvas.NewText(
		a.Id+" is not implemented yet !",
		color.White,
	)
	return container.New(layout.NewVBoxLayout(), mainContent)
}

func New(id string) Activity {
	return &DefaultActivity{id}
}

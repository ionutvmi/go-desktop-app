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

var activityCache = map[string]Activity{}

func (a *DefaultActivity) GetContent() *fyne.Container {
	mainContent := canvas.NewText(
		a.Id+" is not implemented yet !",
		color.White,
	)
	return container.New(layout.NewVBoxLayout(), mainContent)
}
func CachedOrNew(id string) Activity {
	savedActivity, ok := activityCache[id]

	if ok {
		return savedActivity
	}

	activityCache[id] = New(id)
	return activityCache[id]
}

func New(id string) Activity {
	if id == "transform-text" {
		return NewTransformTextActivity(id)
	}

	if id == "transform-text-files" {
		return NewTransformFilesActivity(id)
	}

	return &DefaultActivity{Id: id}
}

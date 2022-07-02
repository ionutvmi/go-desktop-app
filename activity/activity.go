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

func New(id string) Activity {

	savedActivity, ok := activityCache[id]

	if ok {
		return savedActivity
	}

	var newActivity Activity

	if id == "transform-text" {
		newActivity = NewTransformTextActivity(id)
	}

	if id == "transform-text-files" {
		newActivity = NewTransformFilesActivity(id)
	}

	if newActivity != nil {
		activityCache[id] = newActivity
		return newActivity
	}

	return &DefaultActivity{Id: id}
}

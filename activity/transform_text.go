package activity

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TransformTextActivity struct {
	Id string

	textArea        *widget.Entry
	transformedText *widget.Label
}

func NewTransformTextActivity(id string) *TransformTextActivity {
	activity := &TransformTextActivity{Id: id}

	textArea := widget.NewMultiLineEntry()
	transformedText := widget.NewLabel("")

	textArea.OnChanged = func(s string) {
		transformedText.Text = strings.ToUpper(s)
		transformedText.Refresh()
	}
	textArea.PlaceHolder = "type something in here"
	textArea.SetMinRowsVisible(10)

	activity.textArea = textArea
	activity.transformedText = transformedText
	return activity
}

func (activity *TransformTextActivity) GetContent() *fyne.Container {

	output := container.NewScroll(activity.transformedText)
	output.SetMinSize(fyne.NewSize(100, 200))

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Input",
				Widget: activity.textArea,
			},
			{
				Text:   "Output",
				Widget: output,
			},
		},
	}

	return container.NewPadded(form)
}

package activity

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TransformTextActivity struct {
	Id string
}

func (a *TransformTextActivity) GetContent() *fyne.Container {
	textArea := widget.NewMultiLineEntry()
	transformedText := widget.NewLabel("")

	textArea.OnChanged = func(s string) {
		transformedText.Text = strings.ToUpper(s)
		transformedText.Refresh()
	}
	textArea.PlaceHolder = "type something in here"
	textArea.SetMinRowsVisible(10)

	output := container.NewScroll(transformedText)
	output.SetMinSize(fyne.NewSize(100, 200))

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Input",
				Widget: textArea,
			},
			{
				Text:   "Output",
				Widget: output,
			},
		},
	}

	return container.NewPadded(form)
}

package activity

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	output.SetMinSize(fyne.NewSize(300, 200))

	return container.New(
		layout.NewPaddedLayout(),
		container.New(
			layout.NewVBoxLayout(),
			newRow(
				widget.NewLabel("Input text"),
				container.NewHScroll(textArea),
			),
			newRow(
				widget.NewLabel("Output text"),
				output,
			),
		),
	)
}

func newRow(left fyne.CanvasObject, right fyne.CanvasObject) *fyne.Container {
	return container.New(
		layout.NewBorderLayout(nil, nil, left, nil),
		left,
		right,
	)
}

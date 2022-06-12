package main

import (
	"fmt"
	"go-desktop-app/activity"
	"go-desktop-app/sidebar"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	// a.Settings().SetTheme(theme.DarkTheme())

	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(600, 400))

	appTitle := widget.NewLabel("Golang Utilities")

	sidePanel := sidebar.NewSidebar("Activities", []sidebar.Element{
		{
			Id:    "transform-text",
			Label: "Transform text",
		},
		{
			Id:    "transform-text-files",
			Label: "Transform text files",
		},
		{
			Id:    "transform-images",
			Label: "Transform images",
		},
		{
			Id:    "read-excel-file",
			Label: "Read excel file",
		},
	})

	mainPanel := container.NewWithoutLayout()

	sidePanel.OnSelected = func(item sidebar.Element) {
		fmt.Println("selected ", item.Id)

		newAction := activity.New(item.Id)
		mainPanel.Objects = []fyne.CanvasObject{
			newAction.GetContent(),
		}
	}

	var top fyne.CanvasObject = container.New(layout.NewCenterLayout(), appTitle)
	var left fyne.CanvasObject = sidePanel.GetContent()

	appContent := container.New(
		layout.NewBorderLayout(top, nil, left, nil),
		top, left,
		mainPanel,
	)

	w.SetContent(appContent)
	w.ShowAndRun()
}

// cta := widget.NewButton("click me", func() {
// 	menuItems = append(menuItems, "Something new")
// 	list.Refresh()
// })

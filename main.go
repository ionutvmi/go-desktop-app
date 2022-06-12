package main

import (
	"fmt"
	"go-desktop-app/sidebar"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(600, 400))

	appTitle := canvas.NewText("Golang Utilities", color.White)

	sidePanel := sidebar.NewSidebar("Actions", []sidebar.Element{
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
		mainContent := canvas.NewText(
			"Main content here"+item.Label,
			color.White,
		)

		mainPanel.Objects = []fyne.CanvasObject{
			container.New(layout.NewVBoxLayout(), mainContent),
		}
	}

	var top fyne.CanvasObject = container.New(layout.NewCenterLayout(), appTitle)
	var right fyne.CanvasObject = nil
	var bottom fyne.CanvasObject = nil
	var left fyne.CanvasObject = sidePanel.GetContent()

	appContent := container.New(
		layout.NewBorderLayout(top, bottom, left, right),
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

package sidebar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Element struct {
	Id    string
	Label string
}

type Sidebar struct {
	Title      string
	Items      []Element
	OnSelected func(item Element)
}

func NewSidebar(title string, items []Element) *Sidebar {
	return &Sidebar{title, items, nil}
}

func (sidebar *Sidebar) GetContent() *fyne.Container {
	sidebarTitle := container.New(
		layout.NewCenterLayout(),
		widget.NewLabel(sidebar.Title),
	)
	list := sidebar.generateList()

	return container.New(layout.NewBorderLayout(sidebarTitle, nil, nil, nil),
		sidebarTitle,
		list,
	)
}

func (sidebar *Sidebar) generateList() *widget.List {
	list := widget.NewList(
		func() int {
			return len(sidebar.Items)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template with some large text")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			label.SetText(sidebar.Items[i].Label)
		})

	list.OnSelected = func(id widget.ListItemID) {
		if sidebar.OnSelected != nil {
			sidebar.OnSelected(sidebar.Items[id])
		}
	}

	return list
}

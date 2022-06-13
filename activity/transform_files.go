package activity

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type TransformFilesActivity struct {
	Id string

	sourceFolder      fyne.ListableURI
	destinationFolder fyne.ListableURI

	sourceFilesLabel *widget.Label
	destinationLabel *widget.Label
	resultLabel      *widget.Label
}

type FolderType int

const (
	FOLDER_SOURCE FolderType = iota
	FOLDER_DESTINATION
)

func (activity *TransformFilesActivity) GetContent() *fyne.Container {
	activity.sourceFilesLabel = widget.NewLabel("0 files found")
	activity.destinationLabel = widget.NewLabel("")
	activity.resultLabel = widget.NewLabel("")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Source folder",
				Widget: activity.sourceFilesLabel,
			},
			{
				Text: "",
				Widget: widget.NewButton("Select source folder", func() {
					activity.openFolder(FOLDER_SOURCE)
				}),
			},
			{
				Text:   "Destination folder",
				Widget: activity.destinationLabel,
			},
			{
				Text: "",
				Widget: widget.NewButton("Select destination folder", func() {
					activity.openFolder(FOLDER_DESTINATION)
				}),
			},
			{
				Text: "",
				Widget: &widget.Button{
					Text:       "Process Files",
					Importance: widget.HighImportance,
					OnTapped:   activity.processSelectedFiles,
				},
			},
			{
				Text:   "Result log",
				Widget: activity.resultLabel,
			},
		},
	}

	return container.NewBorder(
		container.NewVBox(
			widget.NewLabel("Process text files"),
			widget.NewSeparator(),
		),
		nil,
		nil,
		nil,
		form,
	)
}

func (activity *TransformFilesActivity) openFolder(folderType FolderType) {
	window := fyne.CurrentApp().Driver().AllWindows()[0]
	dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		if list == nil {
			// cancelled
			return
		}

		if folderType == FOLDER_SOURCE {
			activity.sourceFolder = list
		} else if folderType == FOLDER_DESTINATION {
			activity.destinationFolder = list
		}

		activity.updateLabels()
	}, window)
}

func (activity *TransformFilesActivity) updateLabels() {
	if activity.sourceFolder != nil {
		sourceFilesText := activity.sourceFolder.Path() + "\n\n"

		children, err := activity.sourceFolder.List()
		if err != nil {
			activity.sourceFilesLabel.Text = err.Error()
			activity.sourceFilesLabel.Refresh()
			return
		}

		validFilesCount := 0
		for _, v := range children {
			if v.Extension() == ".txt" {
				sourceFilesText += v.Name() + "\n"
				validFilesCount++
			}
		}
		sourceFilesText += strconv.Itoa(validFilesCount) + " files found !"

		activity.sourceFilesLabel.Text = sourceFilesText
		activity.sourceFilesLabel.Refresh()
	}

	if activity.destinationFolder != nil {
		activity.destinationLabel.Text = activity.destinationFolder.Path()
		activity.destinationLabel.Refresh()
	}
}

func (activity *TransformFilesActivity) processSelectedFiles() {

}

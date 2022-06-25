package activity

import (
	"bufio"
	"os"
	"path"
	"strconv"
	"strings"

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

		sourceFiles, err := activity.getValidFileNames()

		if err != nil {
			activity.sourceFilesLabel.Text = err.Error()
			activity.sourceFilesLabel.Refresh()
			return
		}

		sourceFilesText += strings.Join(sourceFiles, "\n") + "\n"
		sourceFilesText += strconv.Itoa(len(sourceFiles)) + " files found !"

		activity.sourceFilesLabel.Text = sourceFilesText
		activity.sourceFilesLabel.Refresh()
	}

	if activity.destinationFolder != nil {
		activity.destinationLabel.Text = activity.destinationFolder.Path()
		activity.destinationLabel.Refresh()
	}
}

func (activity *TransformFilesActivity) getValidFileNames() ([]string, error) {
	result := []string{}

	children, err := activity.sourceFolder.List()
	if err != nil {
		return nil, err
	}

	for _, v := range children {
		if v.Extension() == ".txt" {
			result = append(result, v.Name())
		}
	}

	return result, nil
}

func (activity *TransformFilesActivity) processSelectedFiles() {
	sourceBasePath := activity.sourceFolder.Path()
	destinationBasePath := activity.destinationFolder.Path()

	files, err := activity.getValidFileNames()
	if err != nil {
		return
	}

	for _, name := range files {
		sourceFullPath := path.Join(sourceBasePath, name)
		sourceFile, err := os.OpenFile(sourceFullPath, os.O_RDONLY, os.ModePerm)

		if err != nil {
			println("Failed to open the file", err.Error())
			return
		}
		defer sourceFile.Close()

		destinationFullPath := path.Join(destinationBasePath, name)
		destinationFile, err := os.OpenFile(destinationFullPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)

		if err != nil {
			println("Failed to open the file", err.Error())
			return
		}

		defer destinationFile.Close()

		sc := bufio.NewScanner(sourceFile)

		for sc.Scan() {
			line := sc.Text()
			line = strings.ToUpper(line) + "\n"
			destinationFile.Write([]byte(line))
		}

		println("Processed", name, "successfully !")
	}

	return
}

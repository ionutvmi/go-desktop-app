package activity

import (
	"fmt"
	"image"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"image/jpeg"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type ImagesActivity struct {
	Id string

	sourceFolder      fyne.ListableURI
	destinationFolder fyne.ListableURI

	sourceFilesLabel *widget.Label
	destinationLabel *widget.Label
	resultLabel      *widget.Label
	imagesPreview    *fyne.Container
}

func NewImagesActivity(id string) *ImagesActivity {
	activity := &ImagesActivity{Id: id}

	activity.sourceFilesLabel = widget.NewLabel("0 files found")
	activity.destinationLabel = widget.NewLabel("")
	activity.resultLabel = widget.NewLabel("")

	return activity
}

func (activity *ImagesActivity) GetContent() *fyne.Container {
	resultLog := container.NewScroll(activity.resultLabel)
	resultLog.SetMinSize(fyne.NewSize(100, 200))
	activity.imagesPreview = container.NewGridWrap(
		fyne.NewSize(90, 90),
	)

	sourceLog := container.NewScroll(
		activity.sourceFilesLabel,
	)

	sourceLog.SetMinSize(fyne.NewSize(100, 80))

	previewScroll := container.NewScroll(
		activity.imagesPreview,
	)

	previewScroll.SetMinSize(fyne.NewSize(100, 120))

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Source folder",
				Widget: sourceLog,
			},
			{
				Text:   "",
				Widget: previewScroll,
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
				Widget: resultLog,
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

func (activity *ImagesActivity) openFolder(folderType FolderType) {
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
			activity.log("Selected source folder %s", list.Path())

			go func() {
				activity.updatePreviewImages()
			}()
		} else if folderType == FOLDER_DESTINATION {
			activity.destinationFolder = list
			activity.log("Selected destination folder %s", list.Path())
		}

		activity.updateLabels()
	}, window)
}

func (activity *ImagesActivity) updatePreviewImages() {
	sourceBasePath := activity.sourceFolder.Path()

	files, err := activity.getValidFileNames()
	if err != nil {
		return
	}

	previews := []fyne.CanvasObject{}

	for _, name := range files {
		sourceFullPath := path.Join(sourceBasePath, name)
		previews = append(previews, canvas.NewImageFromFile(sourceFullPath))
	}

	activity.imagesPreview.Objects = previews
	activity.imagesPreview.Refresh()
}

func (activity *ImagesActivity) updateLabels() {
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

func (activity *ImagesActivity) getValidFileNames() ([]string, error) {
	result := []string{}

	children, err := activity.sourceFolder.List()
	if err != nil {
		return nil, err
	}

	for _, v := range children {
		if v.Extension() == ".jpg" {
			result = append(result, v.Name())
		}
	}

	return result, nil
}

func (activity *ImagesActivity) log(msg string, params ...interface{}) {
	currentTime := time.Now()
	timestamp := currentTime.Format("2006-Jan-02 15:04:05")

	msg = fmt.Sprintf(msg, params...)
	msg = fmt.Sprintf("[%s] - %s", timestamp, msg)

	activity.resultLabel.Text = msg + "\n" + activity.resultLabel.Text
	activity.resultLabel.Refresh()
}

func (activity *ImagesActivity) processSelectedFiles() {
	if activity.sourceFolder == nil || activity.destinationFolder == nil {
		activity.log("Missing source or destination folder")
		return
	}

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
			activity.log("Failed to open the file", err.Error())
			return
		}
		defer sourceFile.Close()

		destinationFullPath := path.Join(destinationBasePath, name)
		destinationFile, err := os.OpenFile(destinationFullPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)

		if err != nil {
			activity.log("Failed to open the file %s", err.Error())
			return
		}

		defer destinationFile.Close()

		img, _, err := image.Decode(sourceFile)
		if err != nil {
			activity.log("Failed to decode the file %s", err.Error())
			return
		}

		// Converting image to grayscale
		grayImg := image.NewGray(img.Bounds())
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				grayImg.Set(x, y, img.At(x, y))
			}
		}

		err = jpeg.Encode(destinationFile, grayImg, nil)

		if err != nil {
			activity.log("Failed to encode the file %s", err.Error())
			return
		}

		activity.log("Processed %s successfully !", name)
	}

	return
}

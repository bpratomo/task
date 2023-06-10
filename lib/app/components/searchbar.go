package app

import (
	m "task/lib/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RenderSearchBox(app *tview.Application, tasks []m.Task) *tview.InputField {
	inputField := RenderSearchField(app)
	inputField.SetBorder(true).SetTitle("Search bar")

	return inputField
}

func RenderSearchField(app *tview.Application) *tview.InputField {
	inputField := tview.NewInputField().
		SetPlaceholder("Search or add a new Task").
		// SetFieldWidth(0).
		// SetAcceptanceFunc(tview.InputFieldInteger).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})

	return inputField

}

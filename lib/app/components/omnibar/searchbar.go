package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var inputField *tview.InputField
var placeholderText = "Start typing to search or add a new Task..."

func RenderSearchBox(changeCallback func(string), doneCallback func(string)) *tview.InputField {
	inputField = RenderSearchField()
	var bgStyle tcell.Style
	bgStyle.Background(tcell.ColorDefault)

	inputField.SetFieldBackgroundColor(tcell.ColorDefault)
	inputField.SetPlaceholderStyle(bgStyle)
	inputField.SetChangedFunc(changeCallback)
	inputField.SetDoneFunc(doneFuncFactory(doneCallback))
	// inputField.SetFieldTextColor(tcell.Color16)
	inputField.SetBorder(true).SetTitle("Search bar").SetTitleAlign(tview.AlignLeft)

	return inputField
}

func doneFuncFactory(callback func(string)) func(tcell.Key) {
	return func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			callback(inputField.GetText())
			inputField.SetText("")
		default:
		}
	}

}

func doneFunc(event *tcell.EventKey) {

}

func RenderSearchField() *tview.InputField {
	inputField := tview.NewInputField().
		SetPlaceholder(placeholderText)

	return inputField

}

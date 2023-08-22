package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	g "task/lib/app/global"
	r "task/lib/controllers"
)

var inputField *tview.InputField
var placeholderText = "Start typing to search or add a new Task..."
var globalState *g.GlobalState

func ConfigureOmnibox(state *g.GlobalState) *tview.InputField {
	globalState = state
	inputField = RenderSearchField()
	var bgStyle tcell.Style
	bgStyle.Background(tcell.ColorDefault)

	inputField.SetFieldBackgroundColor(tcell.ColorDefault)
	inputField.SetPlaceholderStyle(bgStyle)
	inputField.SetChangedFunc(onSearchbarChange)
	inputField.SetDoneFunc(doneFuncFactory(onSearchBarSubmit))
	inputField.SetBorder(true).SetTitle("Search bar").SetTitleAlign(tview.AlignLeft)
	inputField.SetFocusFunc(state.FocusFuncFactory(inputField))
	inputField.SetBlurFunc(state.BlurFuncFactory(inputField))

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

func RenderSearchField() *tview.InputField {
	inputField := tview.NewInputField().
		SetPlaceholder(placeholderText)

	return inputField
}

func onSearchBarSubmit(s string) {
	r.Create([]string{s})
	onSearchbarChange("")
}

func onSearchbarChange(s string) {
	globalState.FilterTaskString = s
	globalState.FilterProjectString = s
	globalState.RefreshData([]g.RefreshCategory{g.AllList})
}

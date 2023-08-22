package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	g "task/lib/app/global"
)

var helpText *tview.TextView
var placeholderText = ""
var globalState *g.GlobalState

func ConfigureHelpText(state *g.GlobalState) *tview.TextView {
	globalState = state
	helpText = RenderHelpText()
	var bgStyle tcell.Style
	bgStyle.Background(tcell.ColorDefault)
	helpText.SetBorder(true).SetTitle("Help text")
	helpText.SetFocusFunc(state.FocusFuncFactory(helpText))
	helpText.SetBlurFunc(state.BlurFuncFactory(helpText))

	return helpText
}

func RenderHelpText() *tview.TextView {
	helpText := tview.NewTextView()
	helpText.SetText("")

	return helpText
}

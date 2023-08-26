package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	g "task/lib/app/global"
)

var helpTextComponent *tview.TextView
var placeholderText = ""
var globalState *g.GlobalState

func ConfigureHelpText(state *g.GlobalState) *tview.TextView {
	globalState = state
	helpTextComponent = RenderHelpText()
	var bgStyle tcell.Style
	bgStyle.Background(tcell.ColorDefault)
	helpTextComponent.SetBorder(true).SetTitle("Help text")
	helpTextComponent.SetFocusFunc(state.FocusFuncFactory(helpTextComponent))
	helpTextComponent.SetBlurFunc(state.BlurFuncFactory(helpTextComponent))
	state.ReRenderHelpText = ReRenderHelpText

	return helpTextComponent
}

func RenderHelpText() *tview.TextView {
	helpTextComponent := tview.NewTextView()
	helpText := retrieveHelpText()
	helpTextComponent.SetText(helpText)

	return helpTextComponent
}

func ReRenderHelpText() {
	helpText := retrieveHelpText()
	helpTextComponent.SetText(helpText)
}

func retrieveHelpText() string {
	focusedComponent := globalState.GetFocus()
	switch focusedComponent {
	case globalState.ComponentPointers["projectList"]:
		return "j&k: navigate up & down  | l: go to tasks | a: add task | d: delete project and all associated tasks"
	case globalState.ComponentPointers["taskList"]:
		return "j&k: navigate up & down  | h: go to projects | a: add task | d: delete task"
	case globalState.ComponentPointers["omnibar"]:
		return "esc: exit omnibar | enter: add task"
	default:
		return ""

	}

}

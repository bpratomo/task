package app

import (
	h "task/lib/app/components/helptext"
	o "task/lib/app/components/omnibar"
	l "task/lib/app/components/task_and_project"
	t "task/lib/app/components/taskeditor"
	g "task/lib/app/global"
	m "task/lib/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var appFlex *tview.Flex
var taskFlex *tview.Flex
var omnibar *tview.InputField
var taskList *tview.List
var taskEditor *tview.Form
var pages *tview.Pages
var projectList *tview.List
var helpText *tview.TextView

var globalState g.GlobalState

func Run() {
	configure()

}

func configure() {
	app = tview.NewApplication()
	globalState.GetFocus = app.GetFocus

	globalState.RefreshData([]g.RefreshCategory{g.AllList})
	configureComponents()
	configureLayout()
	configureEventCapture()
}

func configureComponents() {
	taskList, projectList = l.ConfigureLists(&globalState, activateTaskEditor())
	omnibar = o.ConfigureOmnibox(&globalState)
	helpText = h.ConfigureHelpText(&globalState)
}

func configureLayout() {
	taskFlex = tview.NewFlex().
		AddItem(projectList, 0, 1, true).
		AddItem(taskList, 0, 4, false)

	appFlex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(omnibar, 3, 1, true).
		AddItem(taskFlex, 0, 10, false).
		AddItem(helpText, 3, 1, true)

	pages = tview.NewPages().AddPage("main", appFlex, true, true)
}

func configureEventCapture() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			deactivateTaskEditor()
			// appFlex.RemoveItem(omnibar)
			app.SetFocus(taskList)
			globalState.InputMode = false
			return event

		case tcell.KeyRight, tcell.KeyLeft:
			handleHorizontalMovement(event.Key())
			return event
		}

		switch {
		case event.Rune() == 'a' && globalState.InputMode == false:
			// appFlex.Clear()
			// appFlex.AddItem(omnibar, 3, 1, true)
			// appFlex.AddItem(taskFlex, 0, 10, false)
			app.SetFocus(omnibar)
			globalState.InputMode = true
			return nil

		case event.Rune() == 'l' && globalState.InputMode == false:
			handleHorizontalMovement(tcell.KeyRight)
			return nil
		case event.Rune() == 'h' && globalState.InputMode == false:
			handleHorizontalMovement(tcell.KeyLeft)
			return nil

		}

		return event
	})

	if err := app.SetRoot(pages, true).SetFocus(projectList).Run(); err != nil {
		panic(err)
	}
}

func onTaskEditorCancel() func() {
	return func() {
		deactivateTaskEditor()
	}
}

func onTaskEditorSubmit() func() {
	return func() {
		deactivateTaskEditor()
		l.ReRenderLists()
	}

}

func activateTaskEditor() func(m.Task) {
	return func(task m.Task) {
		globalState.TaskBeingEdited = task
		globalState.InputMode = true
		taskEditor = t.ConfigureTaskEditor(&globalState, onTaskEditorSubmit(), onTaskEditorCancel())
		pages.AddPage("editor", taskEditor, true, true)
		app.SetFocus(taskEditor)

	}

}

func deactivateTaskEditor() {
	globalState.TaskBeingEdited = m.Task{}
	globalState.InputMode = false
	pages.RemovePage("editor")
	app.SetFocus(taskList)
}

func handleHorizontalMovement(k tcell.Key) {
	switch app.GetFocus() {
	case projectList:
		app.SetFocus(taskList)

	case taskList:
		app.SetFocus(projectList)
	}

}

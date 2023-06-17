package app

import (
	"sort"
	o "task/lib/app/components/omnibar"
	l "task/lib/app/components/task_and_project"
	g "task/lib/app/global"
	r "task/lib/controllers"
	d "task/lib/database"
	m "task/lib/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var appFlex *tview.Flex
var taskFlex *tview.Flex
var projectMap map[m.Project]bool
var omnibar *tview.InputField
var taskList *tview.List
var projectList *tview.List

var globalState g.GlobalState

func configure() {
	app = tview.NewApplication()
	globalState.DisplayedTasks, projectMap = d.GetAll()
	globalState.DisplayedProjects = convertMapToList(projectMap)

	taskList, projectList = l.ConfigureLists(app, &globalState, refresh)

	omnibar = o.RenderSearchBox(app, onSearchbarChange(), onSearchBarSubmit())

	taskFlex = tview.NewFlex().
		AddItem(projectList, 0, 1, true).
		AddItem(taskList, 0, 4, false)

	appFlex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(taskFlex, 0, 10, false)

}

func Run() {
	configure()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			appFlex.RemoveItem(omnibar)
			app.SetFocus(taskList)
			globalState.InputMode = false
			return event

		case tcell.KeyRight, tcell.KeyLeft:
			handleMovement(event.Key())
			return event
		}

		switch {
		case event.Rune() == 'a' && globalState.InputMode == false:
			appFlex.Clear()
			appFlex.AddItem(omnibar, 3, 1, true)
			appFlex.AddItem(taskFlex, 0, 10, false)
			app.SetFocus(omnibar)
			globalState.InputMode = true
			return nil

		case event.Rune() == 'l' && globalState.InputMode == false:
			handleMovement(tcell.KeyRight)
			return nil
		case event.Rune() == 'h' && globalState.InputMode == false:
			handleMovement(tcell.KeyLeft)
			return nil

		}

		return event
	})

	if err := app.SetRoot(appFlex, true).SetFocus(projectList).Run(); err != nil {
		panic(err)
	}
}

func onSearchbarChange() func(string) {
	return func(s string) {
		globalState.FilterTaskString = s
		refresh()
		l.ReRenderLists()
	}
}

func refresh() {
	globalState.DisplayedTasks, projectMap = d.Get(globalState.FilterTaskString, globalState.FilterProjectString)
	globalState.DisplayedProjects = convertMapToList(projectMap)
}

func onSearchBarSubmit() func(string) {
	return func(s string) {
		r.Create([]string{s})
		onSearchbarChange()("")
		appFlex.RemoveItem(omnibar)
		app.SetFocus(taskList)
		globalState.InputMode = false

	}
}

func handleMovement(k tcell.Key) {
	switch app.GetFocus() {
	case projectList:
		app.SetFocus(taskList)
	case taskList:
		app.SetFocus(projectList)
	case projectList:
		switch k {
		case tcell.KeyRight:
			app.SetFocus(taskList)
		case tcell.KeyLeft:
			app.SetFocus(projectList)
		}
		globalState.InputMode = false
	}
}

func convertMapToList(ms map[m.Project]bool) []m.Project {
	keys := make([]m.Project, 0, len(ms))
	for key := range ms {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i int, j int) bool { return keys[i].Name < keys[j].Name })
	return keys
}

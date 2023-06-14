package app

import (
	"sort"
	c "task/lib/app/components"
	r "task/lib/controllers"
	d "task/lib/database"
	m "task/lib/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var appFlex *tview.Flex
var taskFlex *tview.Flex
var displayedTasks []m.Task
var displayedProjects []m.Project
var projectMap map[m.Project]bool
var inputField *tview.InputField
var taskList *tview.List
var projectList *tview.List

func configure() {
	app = tview.NewApplication()
	displayedTasks, projectMap = d.GetAll()
	displayedProjects = convertMapToList(projectMap)

	taskList = c.ConfigureTaskList(app, displayedTasks)
	projectList = c.ConfigureProjectList(app, displayedProjects)

	inputField = c.RenderSearchBox(app, onSearchbarChange(), onSearchBarSubmit())

	taskFlex = tview.NewFlex().
		AddItem(projectList, 0, 1, true).
		AddItem(taskList, 0, 4, false)

	appFlex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(inputField, 3, 1, true).
		AddItem(taskFlex, 0, 10, false)

}

func Run() {
	configure()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRight, tcell.KeyLeft:
			handleMovement(event.Key())
			return nil

		default:
			return event
		}
	})

	if err := app.SetRoot(appFlex, true).SetFocus(appFlex).Run(); err != nil {
		panic(err)
	}
}

func onSearchbarChange() func(string) {
	return func(s string) {
		displayedTasks, projectMap = d.Get(s)
		displayedProjects = convertMapToList(projectMap)
		c.ReRenderList(taskList, displayedTasks)
        c.ReRenderProjectList(projectList, displayedProjects)
	}
}

func onSearchBarSubmit() func(string) {
	return func(s string) {
		r.Create([]string{s})
		onSearchbarChange()("")
	}
}

func handleMovement(k tcell.Key) {
	var focusedIndex int
	for i := 0; i < appFlex.GetItemCount(); i++ {
		item := appFlex.GetItem(i)
		if item.HasFocus() {
			focusedIndex = i
			break
		}
	}
	var toBeFocusedIndex int
	switch k {
	case tcell.KeyRight:
		if focusedIndex < appFlex.GetItemCount()-1 {
			toBeFocusedIndex = focusedIndex + 1
		} else {
			toBeFocusedIndex = 0
		}

	case tcell.KeyLeft:
		if focusedIndex > 0 {
			toBeFocusedIndex = focusedIndex - 1
		} else {
			toBeFocusedIndex = appFlex.GetItemCount() - 1
		}

	}
	toBeFocusedItem := appFlex.GetItem(toBeFocusedIndex)
	app.SetFocus(toBeFocusedItem)

}

func convertMapToList(ms map[m.Project]bool) []m.Project {
	keys := make([]m.Project, 0, len(ms))
	for key := range ms {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i int, j int) bool { return keys[i].Name < keys[j].Name })
	return keys
}

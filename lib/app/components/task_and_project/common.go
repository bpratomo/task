package app

import (
	g "task/lib/app/global"
	"github.com/rivo/tview"
)

var state *g.GlobalState
var refresh func()

func ConfigureLists(app *tview.Application, globalState *g.GlobalState, refreshCallback func()) (*tview.List, *tview.List) {
	state = globalState
    refresh = refreshCallback

	taskList := ConfigureTaskList(app)
	projectList := ConfigureProjectList(app)

	return taskList, projectList

}

func ReRenderLists() {
    refresh()
	ReRenderTaskList()
	ReRenderProjectList()
}



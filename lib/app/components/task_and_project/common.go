package app

import (
	"github.com/rivo/tview"
	g "task/lib/app/global"
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

func GetNextIndex(list *tview.List, currentIdx int) int {
	if currentIdx+1 > list.GetItemCount()-1 {
		return 0
	} else {
		return currentIdx + 1
	}
}

func GetPrevIndex(list *tview.List, currentIdx int) int {
	if currentIdx-1 < 0 {
		return list.GetItemCount() - 1
	} else {
		return currentIdx - 1
	}
}

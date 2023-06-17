package app

import (
	"github.com/rivo/tview"
	g "task/lib/app/global"
    m "task/lib/models"
)

var state *g.GlobalState
var refresh func()
var activateTaskEditor func(m.Task)

func ConfigureLists(globalState *g.GlobalState, refreshCallback func(), activateTaskEditorCallback func(m.Task)) (*tview.List, *tview.List) {
	state = globalState
	refresh = refreshCallback
    activateTaskEditor = activateTaskEditorCallback

	taskList := ConfigureTaskList()
	projectList := ConfigureProjectList()

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

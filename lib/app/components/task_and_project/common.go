package app

import (
	"github.com/rivo/tview"
	g "task/lib/app/global"
	m "task/lib/models"
)

var state *g.GlobalState
var activateTaskEditor func(m.Task)

func ConfigureLists(globalState *g.GlobalState, activateTaskEditorCallback func(m.Task)) (*tview.List, *tview.List) {
	state = globalState
	activateTaskEditor = activateTaskEditorCallback

	taskList := ConfigureTaskList()
	projectList := ConfigureProjectList()

	state.AddRefreshCallback(g.TaskList, ReRenderTaskList)
	state.AddRefreshCallback(g.ProjectList, ReRenderProjectList)
	state.AddRefreshCallback(g.AllList, ReRenderLists)

	return taskList, projectList

}

func ReRenderLists() {
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

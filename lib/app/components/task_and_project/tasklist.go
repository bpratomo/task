package app

import (
	"strconv"
	c "task/lib/controllers"

	// conf "task/lib/config"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var selectedId string
var selectedIndex int
var taskList *tview.List

func ConfigureTaskList(app *tview.Application) *tview.List {
	taskList = tview.NewList()
	RenderTaskList()
	taskList.SetBorder(true).SetTitle("Task list").SetTitleAlign(tview.AlignLeft)
	taskList.SetChangedFunc(onTaskSelect)
	taskList.SetInputCapture(onTaskKeyPress)
	taskList.SetSelectedFocusOnly(true)
	taskList.ShowSecondaryText(false)

	return taskList

}

func ReRenderTaskList() {
	taskList.Clear()
	RenderTaskList()
}

func RenderTaskList() {
    if state == nil {return}
	for i, task := range state.DisplayedTasks {
		istr := strconv.Itoa(i + 1)
		irunes := []rune(istr)
		taskList.AddItem(task.Title+" - "+task.Project.Name, strconv.Itoa(task.ID), irunes[0], nil)
	}
}

func onTaskSelect(index int, mainText string, secondaryText string, shortcut rune) {
	selectedId = secondaryText
	selectedIndex = index
}

func onTaskKeyPress(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'd':
		c.Delete([]string{selectedId})
        ReRenderLists()
		return event
	default:
		return event
	}
}

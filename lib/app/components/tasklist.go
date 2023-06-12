package app

import (
	"strconv"
	c "task/lib/controllers"
	m "task/lib/models"
	// conf "task/lib/config"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var selectedId string
var selectedIndex int
var curList *tview.List

func ConfigureTaskList(app *tview.Application, tasks []m.Task) *tview.List {
	curList = tview.NewList()
	RenderList(curList, tasks)
	curList.SetBorder(true).SetTitle("Task list").SetTitleAlign(tview.AlignLeft)
	curList.SetChangedFunc(selectionFunc)
	curList.SetInputCapture(handleKeyInput)
	curList.SetSelectedFocusOnly(true)
	curList.ShowSecondaryText(false)

	return curList

}

func ReRenderList(list *tview.List, tasks []m.Task) {
	list.Clear()
	RenderList(list, tasks)
}

func RenderList(list *tview.List, tasks []m.Task) {
	for i, task := range tasks {
		istr := strconv.Itoa(i + 1)
		irunes := []rune(istr)
		list.AddItem(task.Title+" - "+task.Project.Name, strconv.Itoa(task.ID), irunes[0], nil)
	}
}

func selectionFunc(index int, mainText string, secondaryText string, shortcut rune) {
	selectedId = secondaryText
	selectedIndex = index
}

func handleKeyInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'd':
		c.Delete([]string{selectedId})
		curList.RemoveItem(selectedIndex)
		return event
	default:
		return event
	}
}

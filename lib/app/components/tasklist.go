package app

import (
	"strconv"
	c "task/lib/controllers"
	m "task/lib/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var selectedId string
var selectedIndex int
var curList *tview.List

func ConfigureTaskList(app *tview.Application, tasks []m.Task) *tview.List {
	list := RenderList(app, tasks)
	curList = list
	list.SetBorder(true).SetTitle("Task list").SetTitleAlign(tview.AlignLeft)
	list.SetChangedFunc(selectionFunc)
	list.SetInputCapture(handleKeyInput)
	list.SetSelectedFocusOnly(true)
	list.ShowSecondaryText(false)

	return list

}

func RenderList(app *tview.Application, tasks []m.Task) *tview.List {
	list := tview.NewList()

	for i, task := range tasks {
		istr := strconv.Itoa(i + 1)
		irunes := []rune(istr)
		list.AddItem(task.Title, strconv.Itoa(task.ID), irunes[0], nil)
	}

	return list
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

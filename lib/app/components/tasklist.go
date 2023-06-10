package app

import (
	m "task/lib/models"

	"github.com/rivo/tview"
)

func RenderListBox(app *tview.Application, tasks []m.Task) *tview.List {
	list := RenderList(app, tasks)
	list.SetBorder(true).SetTitle("Task list")

	return list

}

func RenderList(app *tview.Application, tasks []m.Task) *tview.List {
	list := tview.NewList()

	for _, task := range tasks {
		list.AddItem(task.Title, "", rune(task.ID), nil)
	}
	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	return list
}

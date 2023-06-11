package app

import (
	"fmt"

	c "task/lib/app/components"
	r "task/lib/controllers"
	d "task/lib/database"
	m "task/lib/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var flex *tview.Flex
var displayedTasks []m.Task
var inputField *tview.InputField
var list *tview.List

func Run() {

	app = tview.NewApplication()
	displayedTasks = d.Get("")

	list = c.ConfigureTaskList(app, displayedTasks)
	inputField = c.RenderSearchBox(app, updateTaskList(), submitTask())

	flex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(inputField, 3, 1, true).
		AddItem(list, 0, 10, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {

		case tcell.KeyRight, tcell.KeyLeft:
			handleMovement(event.Key(), flex, app)
			return nil

		default:
			return event

		}

	})

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}

	fmt.Println(app)
}

func updateTaskList() func(string) {
	return func(s string) {
		displayedTasks = d.Get(s)
		c.ReRenderList(list, displayedTasks)
	}
}

func submitTask() func(string) {
	return func(s string) {
		r.Create([]string{s})
		updateTaskList()("")
		app.SetFocus(list)
		app.SetFocus(flex)

	}
}

func focusOnTasks() {
	if flex == nil {
		return
	}
	app.SetFocus(flex.GetItem(1))
}

func handleMovement(k tcell.Key, flex *tview.Flex, app *tview.Application) {
	var focusedIndex int
	for i := 0; i < flex.GetItemCount(); i++ {
		item := flex.GetItem(i)
		if item.HasFocus() {
			focusedIndex = i
			break
		}
	}
	var toBeFocusedIndex int
	switch k {
	case tcell.KeyRight:
		if focusedIndex < flex.GetItemCount()-1 {
			toBeFocusedIndex = focusedIndex + 1
		} else {
			toBeFocusedIndex = 0
		}

	case tcell.KeyLeft:
		if focusedIndex > 0 {
			toBeFocusedIndex = focusedIndex - 1
		} else {
			toBeFocusedIndex = flex.GetItemCount() - 1
		}

	}
	toBeFocusedItem := flex.GetItem(toBeFocusedIndex)
	app.SetFocus(toBeFocusedItem)

}

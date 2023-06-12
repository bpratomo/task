package app

import (
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

func configure() {
	app = tview.NewApplication()
	displayedTasks = d.GetAll()

	list = c.ConfigureTaskList(app, displayedTasks)
	inputField = c.RenderSearchBox(app, onSearchbarChange(), onSearchBarSubmit())

	flex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(inputField, 3, 1, true).
		AddItem(list, 0, 10, false)

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

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}

func onSearchbarChange() func(string) {
	return func(s string) {
		displayedTasks = d.Get(s)
		c.ReRenderList(list, displayedTasks)
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

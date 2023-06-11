package app

import (
	"fmt"

	c "task/lib/app/components"
	d "task/lib/database"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var inputMode = false

func Run() {

	app = tview.NewApplication()
	tasks := d.Get("")
	list := c.RenderListBox(app, tasks)
	inputField := c.RenderSearchBox(app, tasks)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(inputField, 0, 1, true).
		AddItem(list, 0, 10, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {

		case tcell.KeyRight, tcell.KeyLeft:
			handleMovement(event.Key(), flex, app)
			return nil

		case tcell.KeyCtrlC:
			app.Stop()
		}

		return event

	})

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}

	fmt.Println(app)
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

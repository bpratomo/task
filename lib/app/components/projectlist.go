package app

import (
	"strconv"
	c "task/lib/controllers"
	m "task/lib/models"
	// conf "task/lib/config"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var selectedProjectId string
var selectedProjectIndex int
var curProjectList *tview.List

func ConfigureProjectList(app *tview.Application, projects []m.Project) *tview.List {
	curList = tview.NewList()
	RenderProjectList(curList, projects)
	curList.SetBorder(true).SetTitle("Projects").SetTitleAlign(tview.AlignLeft)
	// curList.SetChangedFunc(selectionFunc)
	curList.SetInputCapture(handleKeyInput)
	curList.SetSelectedFocusOnly(true)
	curList.ShowSecondaryText(false)

	return curList

}

func ReRenderProjectList(list *tview.List, projects []m.Project) {
	list.Clear()
	RenderProjectList(list, projects)
}

func RenderProjectList(list *tview.List, projects []m.Project) {
	for i, project := range projects {
		istr := strconv.Itoa(i + 1)
		irunes := []rune(istr)
		list.AddItem(project.Name, "", irunes[0], nil)
	}
}

func onProjectSelect(index int, mainText string, secondaryText string, shortcut rune) {
	selectedId = secondaryText
	selectedIndex = index
}

func onProjectKeyPress(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'd':
		c.Delete([]string{selectedId})
		curList.RemoveItem(selectedIndex)
		return event
	default:
		return event
	}
}

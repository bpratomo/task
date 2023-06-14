package app

import (
	"strconv"
	c "task/lib/controllers"
	// conf "task/lib/config"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var selectedProjectId string
var selectedProjectIndex int
var curProjectList *tview.List

func ConfigureProjectList(app *tview.Application) *tview.List {
	curProjectList = tview.NewList()
    RenderProjectList()
	curProjectList.SetBorder(true).SetTitle("Projects").SetTitleAlign(tview.AlignLeft)
	// curProjectList.SetChangedFunc(selectionFunc)
	curProjectList.SetInputCapture(onProjectKeyPress)
	curProjectList.SetSelectedFocusOnly(true)
	curProjectList.ShowSecondaryText(false)

	return curProjectList

}

func ReRenderProjectList() {
	curProjectList.Clear()
	RenderProjectList()
}

func RenderProjectList() {
    if state == nil{return}
	for i, project := range state.DisplayedProjects {
		istr := strconv.Itoa(i + 1)
		irunes := []rune(istr)
		curProjectList.AddItem(project.Name, "", irunes[0], nil)
	}
}

func onProjectSelect(index int, mainText string, secondaryText string, shortcut rune) {
	selectedProjectId = secondaryText
	selectedProjectIndex = index
}

func onProjectKeyPress(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'd':
		c.Delete([]string{selectedProjectId})
        ReRenderProjectList()
		return event
	default:
		return event
	}
}

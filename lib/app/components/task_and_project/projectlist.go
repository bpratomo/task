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

func ConfigureProjectList() *tview.List {
	curProjectList = tview.NewList()
	curProjectList.SetBorder(true).SetTitle("Projects").SetTitleAlign(tview.AlignLeft)
	curProjectList.SetChangedFunc(onProjectSelect)
	curProjectList.SetInputCapture(onProjectKeyPress)
	curProjectList.SetSelectedFocusOnly(true)
	curProjectList.ShowSecondaryText(false)
	RenderProjectList()

	return curProjectList

}

func ReRenderProjectList() {
	curProjectList.Clear()
	RenderProjectList()
}

func RenderProjectList() {
	if state == nil {
		return
	}
	curProjectList.AddItem("Show All", "", rune(strconv.Itoa(1)[0]), nil)
	for i, project := range state.DisplayedProjects {
		istr := strconv.Itoa(i + 2)
		irunes := []rune(istr)
		curProjectList.AddItem(project.Name, project.Name, irunes[0], nil)
	}
}

func onProjectSelect(index int, mainText string, secondaryText string, shortcut rune) {
	selectedProjectId = secondaryText
	selectedProjectIndex = index
	state.FilterProjectString = secondaryText
	refresh()
	ReRenderTaskList()
}

func onProjectKeyPress(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'd':
		c.Delete([]string{selectedProjectId})
		ReRenderProjectList()
		return event
	case 'j':
		nextIndex := GetNextIndex(curProjectList, selectedProjectIndex)
		curProjectList.SetCurrentItem(nextIndex)
		return event
	case 'k':
		nextIndex := GetPrevIndex(curProjectList, selectedProjectIndex)
		curProjectList.SetCurrentItem(nextIndex)
		return event

	default:
		return event
	}
}

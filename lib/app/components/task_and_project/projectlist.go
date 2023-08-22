package app

import (
	"strconv"
	c "task/lib/controllers"

	g "task/lib/app/global"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var selectedProjectId string
var selectedProjectIndex int
var curProjectList *tview.List

func ConfigureProjectList() *tview.List {
	selectedProjectIndex = 0
	curProjectList = tview.NewList()
	curProjectList.SetBorder(true).SetTitle("Projects").SetTitleAlign(tview.AlignLeft)
	curProjectList.SetChangedFunc(onProjectSelect)
	curProjectList.SetInputCapture(onProjectKeyPress)
	curProjectList.SetSelectedFocusOnly(true)
	curProjectList.ShowSecondaryText(false)
	curProjectList.SetFocusFunc(state.FocusFuncFactory(curProjectList))
	curProjectList.SetBlurFunc(state.BlurFuncFactory(curProjectList))

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
	// state.RefreshData()
	if index == selectedProjectIndex {
		return
	}
	selectedProjectId = secondaryText
	state.FilterProjectString = secondaryText
	state.RefreshData([]g.RefreshCategory{g.TaskList})
	ReRenderTaskList()
}

func onProjectKeyPress(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'd':
		c.DeleteProject([]string{selectedProjectId})
		state.RefreshData([]g.RefreshCategory{g.AllList})
		return event
	case 'j':
		selectedProjectIndex = curProjectList.GetCurrentItem()
		nextIndex := GetNextIndex(curProjectList, selectedProjectIndex)
		curProjectList.SetCurrentItem(nextIndex)
		return nil
	case 'k':
		selectedProjectIndex = curProjectList.GetCurrentItem()
		nextIndex := GetPrevIndex(curProjectList, selectedProjectIndex)
		curProjectList.SetCurrentItem(nextIndex)
		return nil

	default:
		return event
	}
}

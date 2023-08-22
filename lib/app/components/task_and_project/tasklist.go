package app

import (
	"encoding/json"
	"strconv"
	c "task/lib/controllers"
	m "task/lib/models"

	g "task/lib/app/global"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var selectedTask m.Task
var selectedIndex int
var taskList *tview.List

func ConfigureTaskList() *tview.List {
	taskList = tview.NewList()
	RenderTaskList()
	taskList.SetBorder(true).SetTitle("Task list").SetTitleAlign(tview.AlignLeft)
	taskList.SetChangedFunc(onTaskSelect)
	taskList.SetInputCapture(onTaskKeyPress)
	taskList.SetSelectedFocusOnly(true)
	taskList.ShowSecondaryText(false)
	taskList.SetFocusFunc(state.FocusFuncFactory(taskList))
	taskList.SetBlurFunc(state.BlurFuncFactory(taskList))

	return taskList

}

func ReRenderTaskList() {
	taskList.Clear()
	if len(state.FilterTaskString) > 0 {
		taskList.SetTitle("Task list - FILTERED").SetTitleColor(tcell.Color(tcell.ColorLightPink))
	} else {
		taskList.SetTitle("Task list").SetTitleColor(tcell.ColorWhite)
	}

	RenderTaskList()
}

func RenderTaskList() {
	if state == nil {
		return
	}
	if len(state.DisplayedTasks) == 0 {
		taskList.AddItem("No task yet. Add a new one with the omnibar", "", '0', nil)
	}

	for i, task := range state.DisplayedTasks {
		istr := strconv.Itoa(i + 1)
		irunes := []rune(istr)
		content, _ := json.Marshal(task)
		taskList.AddItem("Project :"+task.Project.Name+" - "+task.Title, string(content), irunes[0], nil)
		if i == 0 {
			selectedTask = task
		}
	}
}

func onTaskSelect(index int, mainText string, secondaryText string, shortcut rune) {
	json.Unmarshal([]byte(secondaryText), &selectedTask)
	selectedIndex = index
}

func onTaskKeyPress(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'e':
		activateTaskEditor(selectedTask)
		return event
	case 'd':
		c.Delete([]string{strconv.Itoa(selectedTask.ID)})
		state.RefreshData([]g.RefreshCategory{g.AllList})
		return event
	case 'j':
		nextIndex := GetNextIndex(taskList, selectedIndex)
		taskList.SetCurrentItem(nextIndex)
		return event
	case 'k':
		nextIndex := GetPrevIndex(taskList, selectedIndex)
		taskList.SetCurrentItem(nextIndex)
		return event

	default:
		return event
	}
}

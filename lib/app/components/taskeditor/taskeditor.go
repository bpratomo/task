package app

import (
	c "task/lib/controllers"

	g "task/lib/app/global"

	"github.com/rivo/tview"
)

var taskEditor *tview.Form
var state *g.GlobalState
var titleInputField *tview.InputField
var projectInputField *tview.InputField
var doneFunc func()

func ConfigureTaskEditor(globalState *g.GlobalState, doneCallback func(), cancelFunc func()) *tview.Form {
	state = globalState
	doneFunc = doneCallback
	taskEditor = tview.NewForm()
	taskEditor.AddInputField("Task Title", state.TaskBeingEdited.Title, 0, nil, nil)
	taskEditor.AddInputField("Project", state.TaskBeingEdited.Project.Name, 0, nil, nil)
	titleInputField = taskEditor.GetFormItemByLabel("Task Title").(*tview.InputField)
	projectInputField = taskEditor.GetFormItemByLabel("Project").(*tview.InputField)
	taskEditor.AddButton("Save", onTaskSubmit)
	taskEditor.AddButton("Cancel", cancelFunc)

	return taskEditor

}

func onTaskSubmit() {
	updatedTask := state.TaskBeingEdited
	updatedTask.Title = titleInputField.GetText()
	updatedTask.Project.Name = projectInputField.GetText()

	c.Update(updatedTask)
	doneFunc()

}

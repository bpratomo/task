package app

import (
	m "task/lib/models"
)

type GlobalState struct {
	TaskBeingEdited     m.Task
	FilterTaskString    string
	FilterProjectString string
	DisplayedTasks      []m.Task
	DisplayedProjects   []m.Project
	InputMode           bool

	RefreshData func()
}

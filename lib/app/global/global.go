package app

import (
	m "task/lib/models"
)

type GlobalState struct {
	DisplayedTasks    []m.Task
	DisplayedProjects []m.Project
	InputMode         bool
}

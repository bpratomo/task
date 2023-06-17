package app

import (
	m "task/lib/models"
)

type GlobalState struct {
    FilterTaskString string
    FilterProjectString string
	DisplayedTasks    []m.Task
	DisplayedProjects []m.Project
	InputMode         bool
}


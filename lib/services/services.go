package services

import (
	"regexp"
	"strings"
	m "task/lib/models"
)

type TaskSubmission struct {
	task    m.Task
	project m.Project
}

func ParseTaskSubmission(s string) m.Task {
	project, projectStr := extractProject(s)
	task := m.Task{
		Title:   strings.ReplaceAll(s, projectStr, ""),
		Project: project,
	}
	return task
}

func extractProject(s string) (m.Project, string) {
	r1, _ := regexp.Compile("#:[^ ]*.")
	r2, _ := regexp.Compile("#{[^}]*.")
	p1 := r1.FindString(s)
	p2 := r2.FindString(s)

	switch {
	case len(p1) > 0:
		return m.Project{Name: p1[2 : len(p1)-1]}, p1
	case len(p2) > 0:
		return m.Project{Name: p2[2 : len(p2)-1]}, p2
	default:
		return m.Project{}, ""
	}
}
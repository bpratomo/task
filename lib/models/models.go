package models

type TaskAuthor struct {
	Name string
}

type Task struct {
	ID      int
	Title   string
	Project Project
}

type Project struct {
	Name string
}

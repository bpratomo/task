package models


type TaskAuthor struct {
	Name string
}

type Task struct {
	ID     int
	Title  string
	Author TaskAuthor
}


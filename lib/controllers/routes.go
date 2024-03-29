package controllers

import (
	"fmt"
	"strconv"
	"strings"
)

import (
	d "task/lib/database"
	m "task/lib/models"
	s "task/lib/services"
)

var latestTaskId = 0

func Create(titleSlice []string) error {
	title := strings.Join(titleSlice, " ")
	t := s.ParseTaskSubmission(title)
	nextId := d.GetNextId()

	t.ID = nextId
	return d.Update(nextId, t)

}

func GetAll(filterSlice []string) {
	filter := strings.Join(filterSlice, " ")
	tasks, _ := d.Get(filter, "")

	for _, task := range tasks {
		fmt.Printf("%v: %v \n", task.ID, task.Title)
	}
}

func Update(t m.Task) error {
	return d.Update(t.ID, t)

}

func UpdateCli(params []string) error {
	if len(params) < 2 {
		fmt.Println("Not enough parameters. Please insert task id and title to update")
	}
	id := params[0]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	titleSlice := params[1:]
	title := strings.Join(titleSlice, " ")
	t := s.ParseTaskSubmission(title)
    t.ID = idInt

	err = d.Update(idInt, t)
	return err
}

func DeleteProject(projects []string) error {
	for _, project := range projects {
		d.DeleteProject(project)
	}
	return nil
}

func Delete(ids []string) error {
	for _, id := range ids {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		err = d.Delete(idInt)
		if err != nil {
			return err
		}
	}
	return nil
}

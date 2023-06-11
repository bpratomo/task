package controllers

import (
	"fmt"
	"strconv"
	"strings"
)

import (
    d "task/lib/database"
    m "task/lib/models"
)

var latestTaskId = 0

func Create(titleSlice []string) error {
	title := strings.Join(titleSlice, " ")
    nextId := d.GetNextId()

	t := m.Task{Title: title, ID: nextId}
	return d.Update(nextId,t)

}

func GetAll(filterSlice []string) {
	filter := strings.Join(filterSlice, " ")
	tasks := d.Get(filter)

	for _, task := range tasks {
		fmt.Printf("%v: %v \n", task.ID, task.Title)
	}
}

func Update(params []string) error {
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

	t := m.Task{ID: idInt, Title: title}
	err = d.Update(idInt, t)
	return err
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

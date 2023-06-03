package controller

import (
	"fmt"
	"strconv"
	"strings"
)

var latestTaskId = 0

func create(titleSlice []string) error {
	title := strings.Join(titleSlice, " ")
    nextId := dbGetNextId()

	t := Task{Title: title, ID: nextId}
	return dbUpdate(nextId,t)

}

func getAll(filterSlice []string) {
	filter := strings.Join(filterSlice, " ")
	tasks := dbGet(filter)

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

	t := Task{ID: idInt, Title: title}
	err = dbUpdate(idInt, t)
	return err
}

func delete(ids []string) error {
	for _, id := range ids {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		err = dbDelete(idInt)
		if err != nil {
			return err
		}
	}
	return nil
}

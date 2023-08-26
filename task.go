package main

import (
	"errors"
	"log"
	"os"
	"task/lib"
)

func main() {
	arg := os.Args[1:]
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configPath := homeDir+"/.config/task_app"
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(configPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	lib.ProcessRequest(arg)
}

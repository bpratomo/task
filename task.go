package main

import (
	// "os"
	"os"
	"task/controller"
)

func main() {
    arg:= os.Args[1:]
    controller.ProcessRequest(arg)
}

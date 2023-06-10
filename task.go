package main

import (
	// "os"
	"os"
	"task/lib"
)

func main() {
    arg:= os.Args[1:]
    lib.ProcessRequest(arg)
}

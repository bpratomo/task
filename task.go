package main

import (
	"os"
	"task/lib"
)

func main() {
    arg:= os.Args[1:]
    lib.ProcessRequest(arg)
}

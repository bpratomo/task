package controller

import "fmt"

func init() {
	connect()

}

func ProcessRequest(s []string) {
	verb := s[0]
	var payload string
	if len(s) > 1 {
		payload = s[1]
	} else {
		payload = ""
	}

	switch verb {
	case "create", "c", "a":
		create(payload)

	case "get", "g":
		getAll(payload)

	case "delete", "d":
		delete(payload)

	default:
		fmt.Println("Verb not valid")
	}

}

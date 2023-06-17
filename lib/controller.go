package lib

import (
	"task/lib/app"
	c "task/lib/controllers"
)

func parseArgs(s []string) (verb string, payload []string) {
	if len(s) > 0 {
		verb = s[0]
	} else {
		verb = ""
	}

	if len(s) > 1 {
		payload = s[1:]
	} else {
		payload = nil
	}

	return verb, payload
}

func ProcessRequest(s []string) {
	verb, payload := parseArgs(s)

	switch verb {
	case "create", "c", "a":
		c.Create(payload)

	case "get", "g":
		c.GetAll(payload)

	case "delete", "d":
		c.Delete(payload)

	case "update", "u":
		c.UpdateCli(payload)

	default:
		app.Run()
	}

}

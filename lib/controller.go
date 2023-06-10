package lib

import "task/lib/app"



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
		create(payload)

	case "get", "g":
		getAll(payload)

	case "delete", "d":
		delete(payload)

	case "update", "u":
		Update(payload)

	default:
		app.Run()
	}

}

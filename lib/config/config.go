package config

type config struct {
	showTaskBy TaskViewOption
}

type TaskViewOption int

const (
	Project TaskViewOption = 0
)

var Config = config{
	showTaskBy: Project,
}

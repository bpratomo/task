package app

import (
	u "task/lib/app/utils"
	d "task/lib/database"
	m "task/lib/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type RefreshCategory int

const (
	TaskList    RefreshCategory = 0
	ProjectList RefreshCategory = 1
	AllList     RefreshCategory = 2
)

type GlobalState struct {
	TaskBeingEdited     m.Task
	FilterTaskString    string
	FilterProjectString string
	DisplayedTasks      []m.Task
	DisplayedProjects   []m.Project
	InputMode           bool
	RefreshCallbacks    map[RefreshCategory]func()
	GetFocus            func() tview.Primitive
}

func (g *GlobalState) FocusFuncFactory(t tview.Primitive) func() {
	switch t.(type) {
	case *tview.InputField:
		return func() {
			t.(*tview.InputField).SetBorderColor(tcell.ColorRed)
		}

	case *tview.List:
		return func() {
			t.(*tview.List).SetBorderColor(tcell.ColorRed)
		}

	}
	return nil

}

func (g *GlobalState) BlurFuncFactory(t tview.Primitive) func() {
	switch t.(type) {
	case *tview.InputField:
		return func() {
			t.(*tview.InputField).SetBorderColor(tcell.ColorWhite)
		}

	case *tview.List:
		return func() {
			t.(*tview.List).SetBorderColor(tcell.ColorWhite)
		}

	}
	return nil

}

func (g *GlobalState) RefreshData(ts []RefreshCategory) {
	var projectMap map[m.Project]bool
	g.DisplayedTasks, projectMap = d.Get(g.FilterTaskString, g.FilterProjectString)
	g.DisplayedProjects = u.ConvertMapToList(projectMap)

	for _, t := range ts {
		f, ok := g.RefreshCallbacks[t]
		if !ok {
			continue
		}
		f()
	}
}

func (g *GlobalState) AddRefreshCallback(r RefreshCategory, f func()) {
	if g.RefreshCallbacks == nil {
		g.RefreshCallbacks = make(map[RefreshCategory]func())
	}
	g.RefreshCallbacks[r] = f

}

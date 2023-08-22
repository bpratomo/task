package app

import (
	"sort"
	m "task/lib/models"
)

func ConvertMapToList(ms map[m.Project]bool) []m.Project {
	keys := make([]m.Project, 0, len(ms))
	for key := range ms {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i int, j int) bool { return keys[i].Name < keys[j].Name })
	return keys
}

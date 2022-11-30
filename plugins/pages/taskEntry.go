package plugin

import "github.com/BrandenWilliams/dubyah/libs/tasklists"

type tasksEntry struct {
	ListTitle string `json:"listTitle"`
	EntryID   string `json:"entryID"`
	Tasks     []task
}

type task struct {
	// The text of the task itself
	TaskText string `json:"taskText"`
	// Task List position
	TaskPosition int `json:"taskPosition"`
	// Bool representing if the task complete
	IsCompleted bool `json:"isCompleted"`
}

func makeTaskEntry(e tasklists.Entry) (tle tasksEntry) {
	tle.ListTitle = e.ListTitle
	tle.EntryID = e.ID

	for _, e := range e.Tasks {
		tle.Tasks = append(tle.Tasks, makeTask(e))
	}

	return tle
}

func makeTask(tl tasklists.Tasks) (ntl task) {
	ntl.TaskText = tl.TaskText
	ntl.TaskPosition = tl.TaskPosition
	ntl.IsCompleted = tl.IsCompleted
	return
}

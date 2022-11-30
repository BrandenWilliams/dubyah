package plugin

import (
	"strconv"

	"github.com/BrandenWilliams/dubyah/libs/tasklists"
)

type TasksEntry struct {
	// EntryID for related tasklist
	EntryID string `json:"entryID" form:"entryID"`
	// The text of the task itself
	TaskText string `json:"taskText" form:"taskText"`
	// Task List position
	TaskPosition string `json:"taskPosition" form:"taskPosition"`
	// Bool representing if the task complete
	IsCompleted bool `json:"isCompleted" form:"isCompleted"`
}

func (e TasksEntry) makeTasksEntry() (ae tasklists.Tasks, err error) {
	ae.TaskText = e.TaskText
	ae.IsCompleted = e.IsCompleted
	if len(e.TaskPosition) != 0 {
		ae.TaskPosition = 1
	}
	ae.TaskPosition, err = strconv.Atoi(e.TaskPosition)
	return
}

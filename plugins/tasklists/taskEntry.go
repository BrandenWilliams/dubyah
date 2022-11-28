package plugin

import "github.com/BrandenWilliams/dubyah/libs/tasklists"

type TasksEntry struct {
	// EntryID for related tasklist
	EntryID string `json:"entryID" form:"entryID"`
	// Title of task
	Title string `json:"title" form:"title"`
	// The text of the task itself
	TaskText string `json:"taskText" form:"taskText"`
	// Task List position
	TaskPosition int `json:"taskPosition" form:"taskPosition"`
	// Bool representing if the task complete
	IsCompleted bool `json:"isCompleted" form:"isCompleted"`
}

func (e TasksEntry) makeTasksEntry() (ae tasklists.Tasks) {
	ae.Title = e.Title
	ae.TaskText = e.Title
	ae.IsCompleted = e.IsCompleted
	ae.TaskPosition = e.TaskPosition

	return
}

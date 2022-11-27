package plugin

import "github.com/BrandenWilliams/dubyah/libs/tasks"

type TasksEntry struct {
	Title       string `json:"title" form:"title"`
	TaskText    string `json:"taskText" form:"taskText"`
	IsCompleted bool   `json:"isCompleted"`
}

func (e TasksEntry) makeTasksEntry() (ae tasks.Entry) {
	ae.Title = e.Title
	ae.TaskText = e.Title
	ae.IsCompleted = e.IsCompleted

	return
}

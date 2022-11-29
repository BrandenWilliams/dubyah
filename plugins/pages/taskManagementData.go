package plugin

import "github.com/BrandenWilliams/dubyah/libs/tasklists"

type TaskManagementData struct {
	CoreData

	Tle  *tasklists.Entry
	Tles []*tasklists.Entry

	TasksEntry     tasksEntry
	TaskListsEntry taskListsEntry

	err error
}

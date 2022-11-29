package plugin

import "github.com/BrandenWilliams/dubyah/libs/tasklists"

type taskListsEntry struct {
	TaskLists []taskList
}

type taskList struct {
	EntryID   string `json:"entryID"`
	ListTitle string `json:"listTitle"`
}

func makeTaskListsEntry(e []*tasklists.Entry) (tle taskListsEntry) {
	for _, e := range e {
		tle.TaskLists = append(tle.TaskLists, makeTaskLists(e))
	}

	return tle
}

func makeTaskLists(tl *tasklists.Entry) (ntl taskList) {
	ntl.EntryID = tl.ID
	ntl.ListTitle = tl.ListTitle

	return
}

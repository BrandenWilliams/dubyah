package plugin

import "github.com/BrandenWilliams/dubyah/libs/tasklists"

type TasklistsEntry struct {
	UserID    string `json:"userID" form:"userID"`
	ListTitle string `json:"listTitle" form:"listTitle"`
}

func (e TasklistsEntry) makeTasklistsEntry() (ae tasklists.Entry) {
	ae.UserID = e.UserID
	ae.ListTitle = e.ListTitle

	return
}

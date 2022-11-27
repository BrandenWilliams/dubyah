package tasks

import (
	"strconv"

	"github.com/hatchify/errors"
	"github.com/mojura/mojura"
)

const (
	// ErrEmptyUserID is returned when the User ID for an Entry is empty
	ErrEmptyUserID = errors.Error("invalid user ID, cannot be empty")
	// ErrEmptyTitle is returned when the title for an Entry is empty
	ErrEmptyTitle = errors.Error("invalid title, cannot be empty")
	// ErrEmptyTaskText is returned when the task text for an Entry is empty
	ErrEmptyTaskText = errors.Error("invalid task text, cannot be empty")
)

type Entry struct {
	// Include mojura.Entry to auto-populate fields/methods needed to match the
	mojura.Entry

	// UserID which Entry is related to, used as a relational reference
	UserID string `json:"userID"`
	// Title of task
	Title string `json:"title"`
	// The text of the task itself
	TaskText string `json:"taskText"`
	// Bool representing if the task complete
	IsCompleted bool `json:"isCompleted"`
}

// GetRelationships will return the relationship IDs associated with the Entry
func (e *Entry) GetRelationships() (r mojura.Relationships) {
	r.Append(e.UserID)
	r.Append(strconv.FormatBool(e.IsCompleted))
	return
}

// Validate will ensure an Entry is valid
func (e *Entry) Validate() (err error) {
	// An error list allows us to collect all the errors and return them as a group
	var errs errors.ErrorList
	// Check to see if User ID is set
	if len(e.UserID) == 0 {
		errs.Push(ErrEmptyUserID)
	}

	if len(e.Title) == 0 {
		errs.Push(ErrEmptyTitle)
	}

	if len(e.TaskText) == 0 {
		errs.Push(ErrEmptyTaskText)
	}

	// Note: If error list is empty, a nil value is returned
	return errs.Err()
}

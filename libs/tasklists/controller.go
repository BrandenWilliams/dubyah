package tasklists

import (
	"context"
	"fmt"

	"github.com/mojura/mojura"
	"github.com/mojura/mojura/filters"
)

const (
	RelationshipUserID  = "userID"
	RelationshipEntryID = "EntryID"
)

var relationships = []string{
	RelationshipUserID,
	RelationshipEntryID,
}

// New will return a new instance of the Controller
func New(opts mojura.Opts) (cp *Controller, err error) {
	var c Controller
	opts.Name = "tasklists"
	if c.m, err = mojura.New[Entry](opts, relationships...); err != nil {
		return
	}

	c.ReadWrapper = mojura.MakeReadWrapper(c.m)
	cp = &c
	return
}

type Controller struct {
	m           *mojura.Mojura[Entry, *Entry]
	ReadWrapper mojura.ReadWrapper[Entry, *Entry]
}

// New will insert a new Entry to the back-end
func (c *Controller) New(ctx context.Context, userID string, e Entry) (created *Entry, err error) {
	// Set entry's user ID
	e.UserID = userID
	// Validate entry
	if err = e.Validate(); err != nil {
		err = fmt.Errorf("error validating")
		return
	}

	err = c.m.Transaction(ctx, func(txn *mojura.Transaction[Entry, *Entry]) (err error) {
		created, err = c.new(txn, e)
		return
	})

	return
}

// Get will retrieve an Entry which has the same ID as the provided entryID
func (c *Controller) Get(entryID string) (entry *Entry, err error) {
	// Attempt to get Entry with the provided ID, pass reference to entry for which values to be applied
	if entry, err = c.m.Get(entryID); err != nil {
		return
	}

	return
}

// GetByUserID will retrieve all Entries associated with given user
func (c *Controller) GetByUserID(userID string) (entries []*Entry, err error) {
	userFilter := filters.Match(RelationshipUserID, userID)
	opts := mojura.NewFilteringOpts(userFilter)
	entries, _, err = c.m.GetFiltered(opts)
	return
}

// Update will update the Entry for a given user ID
func (c *Controller) AddTask(ctx context.Context, entryID string, t Tasks) (updated *Entry, err error) {
	err = c.m.Transaction(ctx, func(txn *mojura.Transaction[Entry, *Entry]) (err error) {
		updated, err = c.addTask(txn, entryID, t)
		return
	})

	return
}

func (c *Controller) DeleteTask(ctx context.Context, entryID string, t Tasks) (deleted *Entry, err error) {
	err = c.m.Transaction(ctx, func(txn *mojura.Transaction[Entry, *Entry]) (err error) {
		deleted, err = c.deleteTask(txn, entryID, t.TaskPosition)
		return
	})

	return
}

func (c *Controller) UpdateTaskText(ctx context.Context, entryID string, taskPosition int, e Entry) (updated *Entry, err error) {
	err = c.m.Transaction(ctx, func(txn *mojura.Transaction[Entry, *Entry]) (err error) {
		updated, err = c.updateTaskText(txn, entryID, taskPosition, &e)
		return
	})

	return
}

func (c *Controller) UpdateTaskPositionUp(ctx context.Context, entryID string, currentPosition int) (updated *Entry, err error) {
	err = c.m.Transaction(ctx, func(txn *mojura.Transaction[Entry, *Entry]) (err error) {
		updated, err = c.moveTaskPositionUp(txn, entryID, currentPosition)
		return
	})

	return
}

func (c *Controller) UpdateTaskPositionDown(ctx context.Context, entryID string, currentPosition int) (updated *Entry, err error) {
	err = c.m.Transaction(ctx, func(txn *mojura.Transaction[Entry, *Entry]) (err error) {
		updated, err = c.moveTaskPositionDown(txn, entryID, currentPosition)
		return
	})

	return
}

func (c *Controller) DeleteTaskList(ctx context.Context, userID string) (removed *Entry, err error) {
	err = c.m.Transaction(ctx, func(txn *mojura.Transaction[Entry, *Entry]) (err error) {
		removed, err = c.delete(txn, userID)
		return
	})

	return
}

// Close will close the controller and it's underlying dependencies
func (c *Controller) Close() (err error) {
	// Since we only have one dependency, we can just call this func directly
	return c.m.Close()
}

func (c *Controller) new(txn *mojura.Transaction[Entry, *Entry], e Entry) (created *Entry, err error) {
	// Insert Entry into mojura.Mojura and return the results
	if created, err = txn.New(e); err != nil {
		return
	}

	return
}

// Update helper method
func (c *Controller) update(txn *mojura.Transaction[Entry, *Entry], entryID string, fn func(*Entry) error) (updated *Entry, err error) {
	var orig *Entry
	if orig, err = txn.Get(entryID); err != nil {
		return
	}

	if err = fn(orig); err != nil {
		return
	}

	// Attempt to validate Entry
	if err = orig.Validate(); err != nil {
		// Entry is not valid, return validation error
		return
	}

	// Insert Entry into mojura.Mojura and return the results
	if _, err = txn.Update(orig.ID, fn); err != nil {
		return
	}

	updated = orig
	return
}

// Delete will remove an Entry for a given entryID
func (c *Controller) delete(txn *mojura.Transaction[Entry, *Entry], entryID string) (deleted *Entry, err error) {
	// Remove Entry from mojura.Mojura
	if deleted, err = txn.Delete(entryID); err != nil {
		return
	}

	return
}

func (c *Controller) updateTaskText(txn *mojura.Transaction[Entry, *Entry], entryID string, taskPosition int, e *Entry) (updated *Entry, err error) {
	updated, err = c.update(txn, entryID, func(orig *Entry) (err error) {
		for n, e := range orig.Tasks {
			if e.TaskPosition == taskPosition {
				orig.Tasks[n].TaskText = e.TaskText
			}
		}

		return
	})

	return
}

func (c *Controller) moveTaskPositionUp(txn *mojura.Transaction[Entry, *Entry], entryID string, currentPosition int) (updated *Entry, err error) {
	updated, err = c.update(txn, entryID, func(orig *Entry) (err error) {
		var newTasks []Tasks
		for _, te := range orig.Tasks {

			switch position := te.TaskPosition; {

			case position == currentPosition:
				te.TaskPosition = te.TaskPosition + 1
				newTasks = append(newTasks, te)
			case position == (currentPosition + 1):
				te.TaskPosition = te.TaskPosition - 1
				newTasks = append(newTasks, te)
			default:
				newTasks = append(newTasks, te)
			}
		}

		orig.Tasks = newTasks

		return
	})

	return
}

func (c *Controller) moveTaskPositionDown(txn *mojura.Transaction[Entry, *Entry], entryID string, currentPosition int) (updated *Entry, err error) {
	updated, err = c.update(txn, entryID, func(orig *Entry) (err error) {
		var newTasks []Tasks
		for _, te := range orig.Tasks {
			switch position := te.TaskPosition; {
			case position == currentPosition:
				te.TaskPosition = te.TaskPosition - 1
				newTasks = append(newTasks, te)
			case position == (currentPosition - 1):
				te.TaskPosition = te.TaskPosition + 1
				newTasks = append(newTasks, te)
			default:
				newTasks = append(newTasks, te)
			}
		}

		orig.Tasks = newTasks

		return
	})

	return
}

func (c *Controller) addTask(txn *mojura.Transaction[Entry, *Entry], entryID string, t Tasks) (updated *Entry, err error) {
	updated, err = c.update(txn, entryID, func(orig *Entry) (err error) {
		orig.Tasks = append(orig.Tasks, t)

		return
	})

	return
}

func (c *Controller) deleteTask(txn *mojura.Transaction[Entry, *Entry], entryID string, taskPosition int) (updated *Entry, err error) {
	updated, err = c.update(txn, entryID, func(orig *Entry) (err error) {
		var newTasks []Tasks
		for _, e := range orig.Tasks {
			switch pos := e.TaskPosition; {
			case pos > taskPosition:
				newTasks = append(newTasks, e)
			case pos == taskPosition:
				continue
			case pos < taskPosition:
				e.TaskPosition = e.TaskPosition - 1
				newTasks = append(newTasks, e)

			}
			if e.TaskPosition != taskPosition {
				e.TaskPosition = e.TaskPosition - 1
				newTasks = append(newTasks, e)
			}
		}

		orig.Tasks = newTasks

		return
	})

	return
}

package plugin

import (
	"fmt"
	"log"

	"github.com/BrandenWilliams/dubyah/libs/tasklists"
	"github.com/gdbu/jump"
	"github.com/gdbu/scribe"
	"github.com/mojura/mojura"
	"github.com/vroomy/common"
	"github.com/vroomy/vroomy"
)

var p Plugin

func init() {
	if err := vroomy.Register("tasklists", &p); err != nil {
		log.Fatal(err)
	}
}

type Plugin struct {
	vroomy.BasePlugin

	tasklists *tasklists.Controller
	out       scribe.Scribe
	// Dependencies
	Jump *jump.Jump  `vroomy:"jump"`
	Opts mojura.Opts `vroomy:"mojura-opts"`
}

// Methods to match plugins.Plugin interface below

// Load will be called by Vroomy on initialization
func (p *Plugin) Load(env vroomy.Environment) (err error) {
	// Initialize Inventory controller
	if p.tasklists, err = tasklists.New(p.Opts); err != nil {
		err = fmt.Errorf("error initializing logs controller: %v", err)
		return
	}

	return
}

// Backend returns the underlying backend to the plugin
func (p *Plugin) Backend() interface{} {
	return p.tasklists
}

// Close will close the plugin
func (p *Plugin) Close() (err error) {
	return p.tasklists.Close()
}

// Handlers below

// New is the handler for creating a new Entry
func (p *Plugin) NewTaskList(ctx common.Context) {
	var (
		e   TasklistsEntry
		err error
	)

	// Parse request body as JSON
	if err = ctx.Bind(&e); err != nil {
		// Error parsing request body, return error
		err = fmt.Errorf("error parsing request body: %v", err)
		p.out.Notification(err.Error())
		return
	}

	userID := ctx.Param("userID")

	var created *tasklists.Entry
	// Attempt to insert parsed Entry into the tasks.Controller
	if created, err = p.tasklists.New(ctx.Request().Context(), userID, e.makeTasklistsEntry()); err != nil {
		// Error inserting new Entry, return error
		err = fmt.Errorf("error creating new entry: %v", err)
		p.out.Notification(err.Error())
		return
	}

	// Create resource key
	resourceKey := jump.NewResourceKey("tasks", created.ID)

	// Set resource permissions for the current user ID
	if err = p.Jump.Permissions().SetPermissions(resourceKey, userID, jump.PermRW); err != nil {
		return
	}

	// Entry was successfully inserted, return newly created ID
	ctx.WriteJSON(200, created)
}

func (p *Plugin) AddNewTask(ctx common.Context) {
	var (
		e   TasksEntry
		err error
	)

	e.EntryID = ctx.Param("entryID")

	// Parse request body as JSON
	if err = ctx.Bind(&e); err != nil {
		// Error parsing request body, return error
		err = fmt.Errorf("error parsing request body: %v", err)
		p.out.Notification(err.Error())
		return
	}

	var updated *tasklists.Entry
	// Attempt to insert parsed Entry into the tasks.Controller
	if updated, err = p.tasklists.AddTask(ctx.Request().Context(), e.EntryID, e.makeTasksEntry()); err != nil {
		// Error inserting new Entry, return error
		err = fmt.Errorf("error creating new entry: %v", err)
		p.out.Notification(err.Error())
		return
	}

	userID := ctx.Get("userID")
	// Create resource key
	resourceKey := jump.NewResourceKey("tasklists", updated.ID)

	// Set resource permissions for the current user ID
	if err = p.Jump.Permissions().SetPermissions(resourceKey, userID, jump.PermRW); err != nil {
		return
	}

	// Entry was successfully inserted, return newly created ID
	ctx.WriteJSON(200, updated)
}

func (p *Plugin) UpdateTaskPositionUp(ctx common.Context) {
	var (
		e   TasksEntry
		err error
	)

	// Parse request body as JSON
	if err = ctx.Bind(&e); err != nil {
		// Error parsing request body, return error
		err = fmt.Errorf("error parsing request body: %v", err)
		p.out.Notification(err.Error())
		return
	}

	var updated *tasklists.Entry
	if updated, err = p.tasklists.UpdateTaskPositionUp(ctx.Request().Context(), e.EntryID, e.TaskPosition); err != nil {
		err = fmt.Errorf("error updating task position up: %v", err)
		p.out.Notification(err.Error())
		return
	}

	ctx.WriteJSON(200, updated)
}

func (p *Plugin) UpdateTaskPositionDown(ctx common.Context) {
	var (
		e   TasksEntry
		err error
	)

	// Parse request body as JSON
	if err = ctx.Bind(&e); err != nil {
		// Error parsing request body, return error
		err = fmt.Errorf("error parsing request body: %v", err)
		p.out.Notification(err.Error())
		return
	}

	var updated *tasklists.Entry
	if updated, err = p.tasklists.UpdateTaskPositionDown(ctx.Request().Context(), e.EntryID, e.TaskPosition); err != nil {
		err = fmt.Errorf("error updating task position down: %v", err)
		p.out.Notification(err.Error())
		return
	}

	ctx.WriteJSON(200, updated)
}

func (p *Plugin) DeleteTask(ctx common.Context) {
	var (
		te      TasksEntry
		deleted *tasklists.Entry
		err     error
	)

	entryID := ctx.Param("entryID")

	// Parse request body as JSON
	if err = ctx.Bind(&te); err != nil {
		// Error parsing request body, return error
		err = fmt.Errorf("error parsing request body: %v", err)
		p.out.Notification(err.Error())
		return
	}

	if deleted, err = p.tasklists.DeleteTask(ctx.Request().Context(), entryID, te.makeTasksEntry()); err != nil {
		err = fmt.Errorf("error deleting task: %v", err)
		p.out.Notification(err.Error())
		return
	}

	ctx.WriteJSON(200, deleted)
}

func (p *Plugin) DeleteTaskList(ctx common.Context) {
	var (
		deleted *tasklists.Entry
		err     error
	)

	entryID := ctx.Param("entryID")

	if deleted, err = p.tasklists.DeleteTaskList(ctx.Request().Context(), entryID); err != nil {
		err = fmt.Errorf("error deleting task list: %v", err)
		p.out.Notification(err.Error())
		return
	}

	ctx.WriteJSON(200, deleted)
}

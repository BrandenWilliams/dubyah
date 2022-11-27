package plugin

import (
	"fmt"
	"log"

	"github.com/BrandenWilliams/dubyah/libs/tasks"
	"github.com/gdbu/jump"
	"github.com/mojura/mojura"
	"github.com/vroomy/common"
	"github.com/vroomy/vroomy"
)

var p Plugin

func init() {
	if err := vroomy.Register("tasks", &p); err != nil {
		log.Fatal(err)
	}
}

type Plugin struct {
	vroomy.BasePlugin

	tasks *tasks.Controller

	// Dependencies
	Jump *jump.Jump  `vroomy:"jump"`
	Opts mojura.Opts `vroomy:"mojura-opts"`
}

// Methods to match plugins.Plugin interface below

// Load will be called by Vroomy on initialization
func (p *Plugin) Load(env vroomy.Environment) (err error) {
	// Initialize Inventory controller
	if p.tasks, err = tasks.New(p.Opts); err != nil {
		err = fmt.Errorf("error initializing logs controller: %v", err)
		return
	}

	return
}

// Backend returns the underlying backend to the plugin
func (p *Plugin) Backend() interface{} {
	return p.tasks
}

// Close will close the plugin
func (p *Plugin) Close() (err error) {
	return p.tasks.Close()
}

// Handlers below

// New is the handler for creating a new Entry
func (p *Plugin) New(ctx common.Context) {
	var (
		e   TasksEntry
		err error
	)

	// Parse request body as JSON
	if err = ctx.Bind(&e); err != nil {
		// Error parsing request body, return error
		err = fmt.Errorf("error parsing request body: %v", err)
		return
	}

	userID := ctx.Get("userID")

	var created *tasks.Entry
	// Attempt to insert parsed Entry into the tasks.Controller
	if created, err = p.tasks.New(ctx.Request().Context(), userID, e.makeTasksEntry()); err != nil {
		// Error inserting new Entry, return error
		err = fmt.Errorf("error creating new entry: %v", err)
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

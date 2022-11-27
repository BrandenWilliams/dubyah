package plugin

import (
	"fmt"
	"log"

	"github.com/BrandenWilliams/dubyah/libs/templates"
	"github.com/BrandenWilliams/dubyah/plugins/errorpages"
	"github.com/gdbu/jump"
	"github.com/vroomy/common"
	"github.com/vroomy/vroomy"
)

var p Plugin

type Plugin struct {
	vroomy.BasePlugin

	// Dependencies
	Templates  *templates.Templates `vroomy:"templates"`
	Jump       *jump.Jump           `vroomy:"jump"`
	ErrorPages *errorpages.Plugin   `vroomy:"errorPages"`
}

func init() {
	if err := vroomy.Register("onboarding", &p); err != nil {
		log.Fatal(err)
	}
}

func (p *Plugin) Load(env vroomy.Environment) (err error) {

	return
}

// New will create a user
func (p *Plugin) New(ctx common.Context) {
	var (
		req    CreateRequest
		userID string
		err    error
	)

	// Bind signup form
	if err = ctx.Bind(&req); err != nil {
		err = fmt.Errorf("error parsing request: %v", err)
		p.ErrorPages.RenderError(ctx, err)
		return
	}

	// Check if password and repeat password match
	if req.Password != req.RepeatPassword {
		err = fmt.Errorf("passwords not equal")
		p.ErrorPages.RenderError(ctx, err)
		return
	}

	// Create user
	if userID, err = createUser(ctx.Request().Context(), req); err != nil {
		p.ErrorPages.RenderError(ctx, err)
		return
	}

	ctx.WriteJSON(200, userID)
}

// Delete will delete a user
func (p *Plugin) Delete(ctx common.Context) {
	var err error
	userID := ctx.Param("userID")
	if err = deleteUser(ctx.Request().Context(), userID); err != nil {
		err = fmt.Errorf("error deleting user: %v", err)
		p.ErrorPages.RenderError(ctx, err)
		return
	}

	ctx.WriteNoContent()
}

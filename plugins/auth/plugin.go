package plugin

import (
	"fmt"
	"log"

	"github.com/BrandenWilliams/dubyah/libs/templates"
	"github.com/BrandenWilliams/dubyah/plugins/errorpages"

	"github.com/gdbu/jump"
	"github.com/gdbu/jump/users"
	"github.com/hatchify/errors"
	"github.com/vroomy/common"
	"github.com/vroomy/vroomy"
)

var p plugin

const (
	errCannotAddToGroup      = errors.Error("you do not have access to add users to groups")
	errCannotRemoveFromGroup = errors.Error("you do not have access to remove users from groups")
	errEmptyUserID           = errors.Error("userID is empty")
)

type plugin struct {
	vroomy.BasePlugin

	authPages AuthPages

	// Dependencies
	Templates  *templates.Templates `vroomy:"templates"`
	ErrorPages *errorpages.Plugin   `vroomy:"errorPages"`
	Jump       *jump.Jump           `vroomy:"jump"`
}

func init() {
	if err := vroomy.Register("auth", &p); err != nil {
		log.Fatal(err)
	}
}

// Load will be called by vroomy on initialization
func (p *plugin) Load(env vroomy.Environment) (err error) {

	if err = p.Templates.ParseAndWatchTemplate("login", &p.authPages.Login); err != nil {
		return
	}

	if err = p.Templates.ParseAndWatchTemplate("signup", &p.authPages.SignUp); err != nil {
		return
	}

	isMirror := env["mojura-sync-mode"] == "mirror"
	if isMirror {
		return
	}

	if err = p.Jump.SetPermission("administration", "administration", jump.PermRWD, jump.PermRWD); err != nil {
		return
	}

	return
}

// Login is the login handler
func (p *plugin) Login(ctx common.Context) {
	var (
		d     LoginData
		login users.User
		err   error
	)

	q := ctx.Request().URL.Query()
	d.RedirectURL = q.Get("redirect")

	if d.RedirectURL == "" {
		d.RedirectURL = "/"
	}

	contentType := ctx.Request().Header.Get("Content-Type")
	switch contentType {
	case "application/x-www-form-urlencoded":
		if err := ctx.Request().ParseForm(); err != nil {
			err = fmt.Errorf("error parsing form: %v", err)
			p.ErrorPages.RenderError(ctx, err)
			return
		}

		login.Email = ctx.Request().Form.Get("email")
		login.Password = ctx.Request().Form.Get("password")

	default:
		if err = ctx.Bind(&login); err != nil {
			p.ErrorPages.RenderError(ctx, err)
			return
		}
	}

	if login.ID, err = p.Jump.Login(ctx, login.Email, login.Password); err != nil {
		d.LoginErr = "Invalid Credentials"
		// TODO: Respond differently based on content type
		rendered := p.authPages.Login.Render(d)
		ctx.WriteString(400, "text/html", rendered)
		return
	}

	var user *users.User
	if user, err = p.Jump.GetUser(login.ID); err != nil {
		err = fmt.Errorf("error getting user %s: %v", login.ID, err)
		// TODO: Respond differently based on content type
		p.ErrorPages.RenderError(ctx, err)
		return
	}

	// TODO: Respond differently based on content type
	ctx.WriteJSON(200, user)
}

// Logout is the logout handler
func (p *plugin) SignUp(ctx common.Context) {

	rendered := p.authPages.SignUp.Render()
	ctx.WriteString(200, "text/html", rendered)
}

// Logout is the logout handler
func (p *plugin) Logout(ctx common.Context) {
	var err error
	if err = p.Jump.Logout(ctx); err != nil {
		p.ErrorPages.RenderError(ctx, err)
		return
	}

	ctx.WriteNoContent()
}

func (p *plugin) LoginPage(ctx common.Context) {
	var d LoginData

	q := ctx.Request().URL.Query()
	d.RedirectURL = q.Get("redirect")

	if d.RedirectURL == "" {
		d.RedirectURL = "/"
	}

	rendered := p.authPages.Login.Render(d)
	ctx.WriteString(200, "text/html", rendered)
}

// AddToGroup will add a user to a group. Currently the following groups are supported:
//   - admins
//   - inventory-managers
//   - content-managers
func (p *plugin) AddToGroup(ctx common.Context) {
	var (
		req groupRequest
		err error
	)

	if err = ctx.Bind(&req); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	var isAdmin bool
	userID := ctx.Get("userID")
	if isAdmin, err = p.Jump.Groups().HasGroup(userID, "admins"); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	if !isAdmin {
		ctx.WriteJSON(401, errCannotAddToGroup)
		return
	}

	switch req.Group {
	case "admins":
	case "inventory-managers":
	case "content-managers":

	default:
		err = fmt.Errorf("group <%s> is not supported", req.Group)
		ctx.WriteJSON(400, err)
		return
	}

	if len(req.UserID) == 0 {
		ctx.WriteJSON(401, errEmptyUserID)
		return
	}

	if err = p.Jump.AddToGroup(req.UserID, req.Group); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteJSON(200, "OK")
}

// RemoveFromGroup will remove a user from a group
func (p *plugin) RemoveFromGroup(ctx common.Context) {
	var (
		req groupRequest
		err error
	)

	if err = ctx.Bind(&req); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	var isAdmin bool
	userID := ctx.Get("userID")
	if isAdmin, err = p.Jump.Groups().HasGroup(userID, "admins"); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	if !isAdmin {
		ctx.WriteJSON(401, errCannotRemoveFromGroup)
		return
	}

	if err = p.Jump.RemoveFromGroup(req.UserID, req.Group); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteJSON(200, "OK")
}

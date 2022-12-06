package plugin

import (
	"fmt"
	"log"

	"github.com/BrandenWilliams/dubyah/libs/templates"
	"github.com/BrandenWilliams/dubyah/plugins/errorpages"
	"github.com/gdbu/jump"
	"github.com/gdbu/jump/users"
	"github.com/vroomy/common"
	"github.com/vroomy/vroomy"
)

var p plugin

type plugin struct {
	vroomy.BasePlugin

	pages pages

	// Dependencies
	Templates  *templates.Templates `vroomy:"templates"`
	Jump       *jump.Jump           `vroomy:"jump"`
	ErrorPages *errorpages.Plugin   `vroomy:"errorPages"`
}

func init() {
	if err := vroomy.Register("auth", &p); err != nil {
		log.Fatal(err)
	}
}

// Load will be called by vroomy on initialization
func (p *plugin) Load(env vroomy.Environment) (err error) {
	if err = p.Templates.ParseAndWatchTemplate("login", &p.pages.login); err != nil {
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

	// TODO: Respond differently based on content type
	contentType := ctx.Request().Header.Get("Content-Type")
	switch contentType {
	case "application/x-www-form-urlencoded":
		fmt.Printf("testing www.form\n")
		if err := ctx.Request().ParseForm(); err != nil {
			err = fmt.Errorf("error parsing form: %v", err)
			p.ErrorPages.RenderError(ctx, err)
			return
		}

		login.Email = ctx.Request().Form.Get("email")
		login.Password = ctx.Request().Form.Get("password")

	default:
		if err = ctx.Bind(&login); err != nil {
			fmt.Printf("error binding: %v\n", err)
			p.ErrorPages.RenderError(ctx, err)
			return
		}
	}

	if login.ID, err = p.Jump.Login(ctx, login.Email, login.Password); err != nil {
		d.LoginErr = "Invalid Credentials"
		rendered := p.pages.login.Render(d)
		ctx.WriteString(400, "text/html", rendered)

		return
	}

	var user *users.User
	if user, err = p.Jump.GetUser(login.ID); err != nil {
		err = fmt.Errorf("error getting user %s: %v", login.ID, err)
		p.ErrorPages.RenderError(ctx, err)
		return
	}
	fmt.Printf("success! userid: %v\n", user.ID)
	ctx.WriteJSON(200, user)
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

	rendered := p.pages.login.Render(d)
	ctx.WriteString(200, "text/html", rendered)
}

/* TODO add group management
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
*/

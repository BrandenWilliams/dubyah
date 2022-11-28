package plugin

import (
	"fmt"
	"log"

	"github.com/BrandenWilliams/dubyah/libs/templates"
	"github.com/BrandenWilliams/dubyah/plugins/meta"
	"github.com/gdbu/scribe"
	"github.com/vroomy/common"
	"github.com/vroomy/vroomy"
)

var p Plugin

func init() {
	p.out = scribe.New("Pages")

	if err := vroomy.Register("pages", &p); err != nil {
		err = fmt.Errorf("error registering, %v", err)
		log.Fatal(err)
	}
}

type Plugin struct {
	vroomy.BasePlugin

	pages    Pages
	CoreData CoreData

	Templates *templates.Templates `vroomy:"templates"`
	Meta      *meta.Plugin         `vroomy:"meta"`
	out       *scribe.Scribe
}

// Load will be called by vroomy on initialization
func (p *Plugin) Load(env vroomy.Environment) (err error) {

	// Common Pages
	if err = p.Templates.ParseAndWatchTemplate("index", &p.pages.Homepage); err != nil {
		return
	}

	if err = p.Templates.ParseAndWatchTemplate("techsupport", &p.pages.TechSupport); err != nil {
		return
	}

	if err = p.Templates.ParseAndWatchTemplate("websites", &p.pages.Websites); err != nil {
		return
	}

	if err = p.Templates.ParseAndWatchTemplate("resume", &p.pages.Resume); err != nil {
		return
	}

	if err = p.Templates.ParseAndWatchTemplate("404", &p.pages.NotFound); err != nil {
		return
	}

	// TaskManager pages
	if err = p.Templates.ParseAndWatchTemplate("taskManagement", &p.pages.TaskListsManagement); err != nil {
		return
	}

	if err = p.Templates.ParseAndWatchTemplate("taskList", &p.pages.Tasklist); err != nil {
		return
	}

	// Onboarding pages
	if err = p.Templates.ParseAndWatchTemplate("signUp", &p.pages.SignUp); err != nil {
		return
	}

	if err = p.Templates.ParseAndWatchTemplate("login", &p.pages.Login); err != nil {
		return
	}

	return
}

// Homepage is the handler for serving the homepage
func (p *Plugin) Homepage(ctx common.Context) {
	var d CoreData
	d.Meta = p.Meta.New(ctx)
	d.PageTitle = "Homepage"

	rendered := p.pages.Homepage.Render(d)
	ctx.WriteString(200, "text/html", rendered)
}

func (p *Plugin) TechSupport(ctx common.Context) {
	var d CoreData
	d.Meta = p.Meta.New(ctx)
	d.PageTitle = "Technical Support"

	rendered := p.pages.TechSupport.Render(d)
	ctx.WriteString(200, "text/html", rendered)
}

func (p *Plugin) Websites(ctx common.Context) {
	var d CoreData
	d.Meta = p.Meta.New(ctx)
	d.PageTitle = "Websites"

	rendered := p.pages.Websites.Render(d)
	ctx.WriteString(200, "text/html", rendered)
}

func (p *Plugin) Resume(ctx common.Context) {
	var d CoreData
	d.Meta = p.Meta.New(ctx)
	d.PageTitle = "Resume"

	rendered := p.pages.Resume.Render(d)
	ctx.WriteString(200, "text/html", rendered)
}

func (p *Plugin) NotFound(ctx common.Context) {
	var d CoreData
	d.Meta = p.Meta.New(ctx)
	d.PageTitle = "404 Not Found"

	rendered := p.pages.NotFound.Render(d)
	ctx.WriteString(200, "text/html", rendered)
}

// Auth pages
func (p *Plugin) SignUp(ctx common.Context) {
	var d CoreData
	d.Meta = p.Meta.New(ctx)
	d.PageTitle = "Sign Up"

	rendered := p.pages.SignUp.Render(d)
	ctx.WriteString(200, "text/html", rendered)
}

func (p *Plugin) LoginPage(ctx common.Context) {
	var d CoreData
	d.Meta = p.Meta.New(ctx)
	d.PageTitle = "Login Up"

	rendered := p.pages.Login.Render(d)
	ctx.WriteString(200, "text/html", rendered)
}

// Task Mangement pages
func (p *Plugin) TaskListsManagement(ctx common.Context) {
	var d CoreData
	d.Meta = p.Meta.New(ctx)
	d.PageTitle = "Task Mangement"

	rendered := p.pages.TaskListsManagement.Render(d)
	ctx.WriteString(200, "text/html", rendered)
}

func (p *Plugin) TaskList(ctx common.Context) {
	var d CoreData
	d.Meta = p.Meta.New(ctx)
	d.PageTitle = "Task list"

	rendered := p.pages.Tasklist.Render(d)
	ctx.WriteString(200, "text/html", rendered)
}

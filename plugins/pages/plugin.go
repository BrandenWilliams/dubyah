package plugin

import (
	"fmt"
	"log"

	"github.com/BrandenWilliams/dubyah/libs/templates"
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

	out *scribe.Scribe
}

// Load will be called by vroomy on initialization
func (p *Plugin) Load(env vroomy.Environment) (err error) {

	if err = p.Templates.ParseAndWatchTemplate("index", &p.pages.Homepage); err != nil {
		return
	}

	if err = p.Templates.ParseAndWatchTemplate("techsupport", &p.pages.TechSupport); err != nil {
		return
	}

	if err = p.Templates.ParseAndWatchTemplate("websites", &p.pages.Websites); err != nil {
		return
	}

	if err = p.Templates.ParseAndWatchTemplate("stackshowcase", &p.pages.StackShowcase); err != nil {
		return
	}

	if err = p.Templates.ParseAndWatchTemplate("resume", &p.pages.Resume); err != nil {
		return
	}

	return
}

// Homepage is the handler for serving the homepage
func (p *Plugin) Homepage(ctx common.Context) {
	p.CoreData.PageTitle = "Homepage"

	rendered := p.pages.Homepage.Render(p.CoreData)
	ctx.WriteString(200, "text/html", rendered)
}

// Homepage is the handler for serving the homepage
func (p *Plugin) TechSupport(ctx common.Context) {
	p.CoreData.PageTitle = "Technical Support"

	rendered := p.pages.TechSupport.Render(p.CoreData)
	ctx.WriteString(200, "text/html", rendered)
}

// Homepage is the handler for serving the homepage
func (p *Plugin) Websites(ctx common.Context) {
	p.CoreData.PageTitle = "Websites"

	rendered := p.pages.Websites.Render(p.CoreData)
	ctx.WriteString(200, "text/html", rendered)
}

// Homepage is the handler for serving the homepage
func (p *Plugin) StackShowcase(ctx common.Context) {
	p.CoreData.PageTitle = "Stack Show Case"

	rendered := p.pages.StackShowcase.Render(p.CoreData)
	ctx.WriteString(200, "text/html", rendered)
}

// Homepage is the handler for serving the homepage
func (p *Plugin) Resume(ctx common.Context) {
	p.CoreData.PageTitle = "Resume"

	rendered := p.pages.Resume.Render(p.CoreData)
	ctx.WriteString(200, "text/html", rendered)
}

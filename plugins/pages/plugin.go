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

	pages pages

	Templates *templates.Templates `vroomy:"templates"`

	out *scribe.Scribe
}

// Load will be called by vroomy on initialization
func (p *Plugin) Load(env vroomy.Environment) (err error) {

	if err = p.Templates.ParseAndWatchTemplate("index", &p.pages.homepage); err != nil {
		return
	}

	return
}

// Homepage is the handler for serving the homepage
func (p *Plugin) Homepage(ctx common.Context) {

	rendered := p.pages.homepage.Render()
	ctx.WriteString(200, "text/html", rendered)
}

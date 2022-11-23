package plugin

import (
	"fmt"

	"github.com/BrandenWilliams/dubyah/libs/templates"
	"github.com/gdbu/scribe"
	"github.com/vroomy/vroomy"
)

var p Plugin

func init() {
	p.out = scribe.New("Pages")
	if err := vroomy.Register("pages", &p); err != nil {
		err = fmt.Errorf("error registering, %v", err)
		fmt.Printf("%v", err)
	}
}

type Plugin struct {
	vroomy.BasePlugin

	pages pages

	Templates *templates.Templates `vroomy:"templates"`

	out *scribe.Scribe
}

// Load will be called by vroomy on initialization
func (p *Plugin) Load(env map[string]string) (err error) {

	if err = p.Templates.ParseAndWatchTemplate("index", &p.pages.homepage); err != nil {
		return
	}

	return
}

func (p *Plugin) Backend() interface{} {
	return p
}

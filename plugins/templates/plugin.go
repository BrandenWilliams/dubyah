package luxutils

import (
	"log"

	"github.com/gdbu/scribe"
	"github.com/vroomy/vroomy"

	"github.com/BrandenWilliams/dubyah/libs/templates"
)

var p Plugin

type Plugin struct {
	vroomy.BasePlugin

	out       *scribe.Scribe
	templates *templates.Templates
}

func init() {
	p.out = scribe.New("Templates")

	if err := vroomy.Register("templates", &p); err != nil {
		log.Fatal(err)
	}
}

// Init is called by vroomy during the plugin initialization pass
func (p *Plugin) Load(env vroomy.Environment) (err error) {
	p.templates = templates.New()
	return
}

// Backend exposes this plugin's data layer to other plugins
func (p *Plugin) Backend() interface{} {
	return p.templates
}

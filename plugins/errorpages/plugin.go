package errorpages

import (
	"log"
	"strings"

	"github.com/BrandenWilliams/dubyah/libs/templates"
	"github.com/vroomy/common"
	"github.com/vroomy/vroomy"
)

func init() {
	if err := vroomy.Register("errorPages", &p); err != nil {
		log.Fatal(err)
	}
}

var p Plugin

type Plugin struct {
	vroomy.BasePlugin

	pages Pages

	Templates *templates.Templates `vroomy:"templates"`
}

func (p *Plugin) Backend() interface{} {
	return p
}

// Load will be called by vroomy on initialization
func (p *Plugin) Load(env vroomy.Environment) (err error) {
	if err = p.Templates.ParseAndWatchTemplate("errorpage", &p.pages.ErrorPage); err != nil {
		return
	}

	return
}

func (p *Plugin) RenderError(ctx common.Context, err error) {
	var (
		ep ErrorPayload
	)

	ep.Error = err.Error()

	accept := ctx.Request().Header.Get("accept")
	splitAccept := strings.Split(accept, ",")

	switch splitAccept[0] {
	case "text/html":
		renderError := p.pages.ErrorPage.Render(ep)
		ctx.WriteString(400, "text/html", renderError)
		return
	default:
		ctx.WriteJSON(400, err)
		return
	}
}

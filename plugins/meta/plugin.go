package meta

import (
	"log"
	"net/url"

	"github.com/gdbu/jump"
	"github.com/vroomy/common"
	"github.com/vroomy/vroomy"
)

func init() {
	if err := vroomy.Register("meta", &p); err != nil {
		log.Fatal(err)
	}
}

var p Plugin

type Plugin struct {
	vroomy.BasePlugin

	Jump *jump.Jump `vroomy:"jump"`
}

func (p *Plugin) Backend() interface{} {
	return p
}

func (p *Plugin) New(ctx common.Context) *Meta {
	m := p.makeMeta(ctx)
	return &m
}

func (p *Plugin) makeMeta(ctx common.Context) (m Meta) {
	m.ViewingUserID = ctx.Get("userID")
	m.CurrentURL = url.QueryEscape(ctx.Request().URL.String())
	m.Query = ctx.Request().URL.Query().Get("q")
	m.IsLoggedIn = m.ViewingUserID != ""
	m.IsAdmin, _ = p.Jump.Groups().HasGroup(m.ViewingUserID, "admins")
	return
}

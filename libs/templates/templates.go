package templates

import (
	"fmt"
	"time"

	"github.com/gdbu/poller"
	"github.com/gdbu/scribe"
	"github.com/hoisie/mustache"
)

func New() *Templates {
	var t Templates
	return &t
}

type Templates struct {
	out scribe.Scribe
}

func (t *Templates) ParseAndWatchTemplate(key string, tmpl *mustache.Template) (err error) {
	filename := "./private_html/tmpls/" + key + ".html"
	if err = t.parseTemplate(key, filename, tmpl); err != nil {
		return
	}

	if err = t.watchTemplate(key, filename, tmpl); err != nil {
		return
	}

	return
}

func (t *Templates) parseTemplate(key, filename string, tmpl *mustache.Template) (err error) {
	var parsed *mustache.Template
	if parsed, err = mustache.ParseFile(filename); err != nil {
		err = fmt.Errorf("error parsing %s: %v", key, err)
		return
	}

	t.out.Notification("Template parsed, updating!")
	*tmpl = *parsed
	return
}

func (t *Templates) watchTemplate(key, filename string, tmpl *mustache.Template) (err error) {
	var pol *poller.Poller
	if pol, err = poller.New(filename, func(evt poller.Event) {
		if !isRelevantEvent(evt) {
			return
		}

		t.out.Notification("File changes, parsing template!")
		if err := t.parseTemplate("index", filename, tmpl); err != nil {
			t.out.Errorf("error parsing template %s: %v", key, err)
			return
		}
	}); err != nil {
		return
	}

	go pol.Run(time.Millisecond * 100)
	return
}

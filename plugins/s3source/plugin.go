package plugin

import (
	"context"
	"io"
	"log"

	"github.com/mojura/kiroku"
	"github.com/vroomy/vroomy"
)

var p Plugin

func init() {
	if err := vroomy.Register("mojura-source", &p); err != nil {
		log.Fatal(err)
	}
}

type Plugin struct {
	vroomy.BasePlugin
	kiroku.Source
}

// Backend exposes this plugin's data layer to other plugins
func (p *Plugin) Backend() interface{} {
	return &source{}
}

type source struct{}

func (s *source) Export(ctx context.Context, filename string, r io.Reader) error { return nil }

func (s *source) Import(ctx context.Context, filename string, w io.Writer) error { return nil }

func (s *source) Get(ctx context.Context, filename string, fn func(io.Reader) error) error {
	return nil
}

func (s *source) GetNext(ctx context.Context, prefix, lastFilename string) (filename string, err error) {
	return "", nil
}

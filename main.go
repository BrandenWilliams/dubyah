package main

import (
	"context"
	"log"

	"github.com/hatchify/closer"
	"github.com/vroomy/vroomy"

	// External plugins
	_ "github.com/vroomy-ext/digitalocean-s3-plugin"
	_ "github.com/vroomy-ext/fileserver-plugin"
	_ "github.com/vroomy-ext/jump-plugin"

	// Internal plugins
	_ "github.com/BrandenWilliams/dubyah/plugins/pages"
	_ "github.com/BrandenWilliams/dubyah/plugins/templates"
)

func main() {
	var (
		svc *vroomy.Vroomy
		err error
	)

	if svc, err = vroomy.New("./config.toml"); err != nil {
		log.Fatal(err)
	}

	c := closer.New()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		_ = c.Wait()
		cancel()
	}()

	if err = svc.Listen(ctx); err != nil && err != context.Canceled {
		log.Fatal(err)
	}

	if err = svc.Close(); err != nil {
		log.Fatal(err)
	}
}

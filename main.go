package main

import (
	"context"
	"log"

	"github.com/hatchify/closer"
	"github.com/vroomy/vroomy"

	// External plugins
	_ "github.com/vroomy-ext/fileserver-plugin"
	_ "github.com/vroomy-ext/jump-plugin"
	_ "github.com/vroomy-ext/mojura-opts-plugin"

	// Internal plugins
	_ "github.com/BrandenWilliams/dubyah/plugins/auth"
	_ "github.com/BrandenWilliams/dubyah/plugins/errorpages"
	_ "github.com/BrandenWilliams/dubyah/plugins/meta"
	_ "github.com/BrandenWilliams/dubyah/plugins/onboarding"
	_ "github.com/BrandenWilliams/dubyah/plugins/pages"
	_ "github.com/BrandenWilliams/dubyah/plugins/s3source"
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

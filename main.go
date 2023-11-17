package main

import (
	"context"

	"codeberg.org/woodpecker-plugins/go-plugin"
)

func main() {
	p := &plugin.Plugin{}
	plugin.New(plugin.Options{
		Name:        "plugin-extend-env",
		Description: "Export semver variables to an env file",
		Version:     "v0.0.1",
		Execute: func(ctx context.Context) error {
			err := handler(&p.Metadata)
			if err !=nil {
				return err
			}
			return nil
		},
	}).Run()
}

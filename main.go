package main

import (
	"codeberg.org/woodpecker-plugins/go-plugin"
)

func main() {
	p := &Plugin{}
	p.Plugin = plugin.New(plugin.Options{
		Name:        "var-to-env",
		Description: "Export semver variables to an env file",
		Version:     "v0.0.1",
		Execute:     p.execute,
	})

	p.Run()
}

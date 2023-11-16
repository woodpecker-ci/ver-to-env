package main

import (
	"context"
	"strconv"

	"codeberg.org/woodpecker-plugins/go-plugin"
	"github.com/Masterminds/semver/v3"
	"github.com/joho/godotenv"
)

func main() {
	plugin.New(plugin.Options{
		Name:        "var-to-env",
		Description: "Export semver variables to an env file",
		Version:     "v0.0.1",
		Execute: func(ctx context.Context) error {
			var env map[string]string
			env, err := godotenv.Read()
			if err != nil {
				return err
			}

			v, err := semver.NewVersion(env["CI_COMMIT_TAG"])
			if err != nil {
				env["CI_COMMIT_TAG_IS_SEMVER"] = "false"
			}
			env["CI_COMMIT_TAG_SEMVER"] = v.String()
			env["CI_COMMIT_TAG_SEMVER_MAJOR"] = strconv.FormatUint(v.Major(), 10)
			env["CI_COMMIT_TAG_SEMVER_MINOR"] = strconv.FormatUint(v.Minor(), 10)
			env["CI_COMMIT_TAG_SEMVER_PATCH"] = strconv.FormatUint(v.Patch(), 10)
			if v.Prerelease() != "" {
				env["CI_COMMIT_TAG_SEMVER_PRERELEASE"] = v.Prerelease()
				env["CI_COMMIT_TAG_IS_PRERELEASE"] = "true"
			}

			return godotenv.Write(env, ".env")
		},
	}).Run()
}

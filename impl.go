package main

import (
	"strconv"

	"codeberg.org/woodpecker-plugins/go-plugin"
	"github.com/Masterminds/semver/v3"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func handler(m *plugin.Metadata) error {
	eenv := ExtendedEnv{}
	e, err := eenv.LoadEnv()
	if err != nil {
		log.Error().Msg("Unable to load existing .env file!")
	}
	v, err := semver.NewVersion(e["CI_COMMIT_TAG"])

	if err != nil {
		ev["CI_COMMIT_TAG_IS_SEMVER"] = "false"
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
}

package main

import (
	"github.com/Masterminds/semver/v3"
	"github.com/joho/godotenv"
)

type (
	// EnvMap is a type alias form map[string]string
	//EnvMap map[string]string

	// Metadata for semver data on runtime.
	SemVer struct {
		IsValid      bool
		IsPrerelease bool
		Version      semver.Version
	}

	// Holds full extended env representation on runtime.
	ExtendedEnv struct {
		Semver SemVer
		// Embed additional structs to extend
	}
)

func New() *ExtendedEnv {
	eenv := &ExtendedEnv{
		Semver: ,
	}
	return eenv
}

// Get defined SemVer specific env vars.
func (m *ExtendedEnv) getSemVerEnv() map[string]string {
	env := map[string]string{
		"CI_COMMIT_TAG_IS_SEMVER":         "",
		"CI_COMMIT_TAG_SEMVER":            "",
		"CI_COMMIT_TAG_SEMVER_MAJOR":      "",
		"CI_COMMIT_TAG_SEMVER_MINOR":      "",
		"CI_COMMIT_TAG_SEMVER_PATCH":      "",
		"CI_COMMIT_TAG_SEMVER_PRERELEASE": "",
		"CI_COMMIT_TAG_IS_PRERELEASE":     "",
	}
	return env
}

// Gets all sub env scope values as map.
func (m *ExtendedEnv) GetFullEnv() map[string]string {
	env := map[string]string{}
	mergeEnv(&env, m.getSemVerEnv())
	return env
}

// Merges input map into remaining map pointer.
func mergeEnv(extEnv *map[string]string, input map[string]string) {
	for k, v := range input {
		(*extEnv)[k] = v
	}
}

// Load current .env file into current scope.
func (m *ExtendedEnv) LoadEnv() (map[string]string, error) {
	env, err := godotenv.Read()
	if err != nil {
		return nil, err
	}
	return env, nil
}

// Save update Env setting into .env file.
func (m *ExtendedEnv) UpdateEnv(extEnv *map[string]string) error {
	return godotenv.Write(m, ".env")
}

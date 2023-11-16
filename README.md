# plugin-extend-env

<p align="center">
  <a href="https://ci.woodpecker-ci.org/woodpecker-ci/plugin-extend-env" title="Build Status">
    <img src="https://ci.woodpecker-ci.org/api/badges/woodpecker-ci/plugin-extend-env/status.svg">
  </a>
  <a href="https://goreportcard.com/badge/github.com/woodpecker-ci/plugin-extend-env" title="Go Report Card">
    <img src="https://goreportcard.com/badge/github.com/woodpecker-ci/plugin-extend-env">
  </a>
  <a href="https://app.fossa.com/projects/git%2Bgithub.com%2Fwoodpecker-ci%2Fplugin-extend-env?ref=badge_shield" alt="FOSSA Status">
    <img src="https://app.fossa.com/api/projects/git%2Bgithub.com%2Fwoodpecker-ci%2Fplugin-extend-env.svg?type=shield"/>
  </a>
  <a href="https://godoc.org/github.com/woodpecker-ci/plugin-extend-env" title="GoDoc">
    <img src="https://godoc.org/github.com/woodpecker-ci/plugin-extend-env?status.svg">
  </a>
  <a href="https://hub.docker.com/r/woodpeckerci/plugin-extend-env" title="Docker pulls">
    <img src="https://img.shields.io/docker/pulls/woodpeckerci/plugin-extend-env">
  </a>
  <a href="https://opensource.org/licenses/Apache-2.0" title="License: Apache-2.0">
    <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg">
  </a>
</p>

The extend env plugin extends an existing or creates a new `.env` file with additional variables like semver information.

## Usage

```yml
steps:
  extend-env:
    image: woodpeckerci/plugin-extend-env
```

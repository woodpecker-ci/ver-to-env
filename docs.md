---
name: Extend env plugin
authors: Woodpecker Authors
icon: https://woodpecker-ci.org/img/logo.svg
description: Extend your .env file with additional variables like semver information.
tags: [env, semver]
containerImage: woodpeckerci/plugin-extend-env
containerImageUrl: https://hub.docker.com/r/woodpeckerci/plugin-extend-env
url: https://github.com/woodpecker-ci/plugin-extend-env
---

The extend env plugin extends an existing or creates a new `.env` file with additional variables like semver information.

The below pipeline configuration demonstrates simple usage:

```yml
steps:
  extend-env:
    image: woodpeckerci/plugin-extend-env
```

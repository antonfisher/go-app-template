# go-app-template

[![Build Status](https://travis-ci.org/antonfisher/go-app-template.svg?branch=master)](https://travis-ci.org/antonfisher/go-app-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/antonfisher/go-app-template)](https://goreportcard.com/report/github.com/antonfisher/go-app-template)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

# Go Application Template

- pre-configured `Makefile`
- build with version based on Git branch and commit (``)

## Instructions
- clone repository
- remove `.git` folder
    ```bash
    rm -r .git
    ```
- replace REPOSITORY name in
    - `Makefile`
    - `README.md`

## Build

```bash
make build
```

## Run

```bash
bin/go-app-template
```

Options:

```bash
# help
bin/go-app-template --help # print all available flags

# version
bin/go-app-template --version # print application version
```

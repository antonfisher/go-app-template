# Go Application Template

[![Build Status](https://travis-ci.org/antonfisher/go-app-template.svg?branch=master)](https://travis-ci.org/antonfisher/go-app-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/antonfisher/go-app-template)](https://goreportcard.com/report/github.com/antonfisher/go-app-template)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

Lightweight Go application template includes:
- file structure backbone
- pre-configured `Makefile`
- build version based on a Git branch and the last commit (like `go-app-template@master-26a6e8a`)
- `--version` and `--help` commands
- application waits for user interruption (<kbd>Ctrl</kbd> + <kbd>C</kbd>)
- travis-ci config

## Instructions
- clone repository
    ```bash
    cd /tmp
    git clone https://github.com/antonfisher/go-app-template.git
    mv do-app-template YOUR_PROJECT_FOLDER
    ```
- remove old git folder and init new git repository
    ```bash
    cd YOUR_PROJECT_FOLDER
    rm -r .git README.md
    git init
    ```
- replace `NAME` and `REPOSITORY` in `Makefile`
- replace imports in `cmd/main.go` file
- remove or update `LICENSE.md` file
- build project
    ```bash
    make
    ````

## Build

```bash
make clean
make build
```

## Run

```bash
bin/go-app-template
```

Options:

```bash
# help - prints all available flags
$ bin/go-app-template --help
Usage of bin/go-app-template:
  -version
        print application version

# version - prints application version
$ bin/go-app-template --version
go-app-template@master-26a6e8a
```

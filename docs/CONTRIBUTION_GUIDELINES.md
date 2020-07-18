# Development guide

## Prerequisites

You need to have the following tools installed:
- :mouse2: go
- :whale: docker
- direnv
- make
- [golint](https://github.com/golang/lint)

Also it would be easier to be on a Unix (or Unix-like) OS. The following environment variables need to be set - GOBIN, GOPATH.

_**:bangbang: Note:** This project is currently being developed for a `postgre` db._

## :wrench: Workflow

We are following a standard workflow:
- each new feature is developed in a `feature branch`
- there is no `dev branch` and everything is merged directly into `master`
- run all the tests before making a PR
- good to request reviews
- hacky and untested scripts should go to the `hack` folder in the root of the project

Before committing:
- run `make clean-code`
- run `shellcheck` on scripts

## :hammer: Build

To build just run:

```bash
go get ./...
# Export the env vars as described below.
make build
```

## Technologies

The server is written in `go` and tests are in `ginkgo`. For db migrations and orm we are using [gorm](https://gorm.io/). The project is built with a `GNU Makefile`.

## Makefile

A lot of the work is done through `Makefile`. The following make commands are supported:

```bash
make # installs a local binary to your GOBIN and builds a docker image
make build # the same as make
make clean # remove local binary from GOBIN and delete local docker image

make check-compliance # checks if you have the prerequisites set

make build-binary # installs a local binary to your GOBIN
make clean-binary # removes your local binary from GOBIN

make build-docker-image # builds a docker image
make clean-docker-image # delete local docker image
make run-docker-image # run docker image
make push-docker-image # push docker image to a docker registry
```

## Environment

Most of the variable you need would be loaded from `.envrc`. You would additionally need:
- `DOCKERHUB_USERNAME` - Your dockerhub username
- `DOCKERHUB_PASSWORD` - Your dockerhub password
- `DB_HOST` - Your db host
- `DB_PORT` - Your db port
- `DB_USER` - Your postgre db user
- `DB_PASS` - Your postgre db password
- `DB_NAME` - Your postgre db name

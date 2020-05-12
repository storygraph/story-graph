# :jack_o_lantern: :lollipop: Story Graph
Story graph is a state-control system for storytellers. It models a story as a set of **storylines with states at every given moment.**. Each state is modeled via a graph where the vertices and the edges are respectively the story objects and their corresponding relations. As the story goes by the **state is modified** via applying a **delta of changes** in moments called **events**.

**Basically this is Git for storytellers.**

## How can I use this?

You can run it locally via docker or you can use the [k8s release](https://github.com/storygraph/story-graph-release).

## Local development

### Prerequisites
You need to have the following tools installed:
- :mouse2: go
- :whale: docker
- direnv
- make
Also you would need a Unix (or Unix-like) OS. The following environment variables need to be set - GOBIN, GOPATH.

### Make commands

For app builds:
```bash
make # installs a local binary to your GOBIN and builds a docker image
make build # the same as make
make clean # remove local binary from GOBIN and delete local docker image

make check-compliance # checks if you have the prerequisites set

make build-binary # installs a local binary to your GOBIN
make clean-binary # removes your local binary from GOBIN

make build-docker-image
make clean-docker-image # delete local docker image
make run-docker-image
make push-docker-image
```

Running tests:
```bash
make run-tests
```

# :jack_o_lantern: :lollipop: Story Graph
Story graph is a state-control system for storytellers. It models a story as a set of **storylines with states at every given moment.**. Each state is modeled via a graph where the vertices and the edges are respectively the story objects and their corresponding relations. As the story goes by the **state is modified** via applying a **delta of changes** in moments called **events**.

**Basically this is Git for storytellers.**

## How can I use this?
Currently you have two alternatives:

### Makefile
#### Prerequisites
You need to have the following tools installed:
- :mouse2: go
- :whale: docker
- direnv
- make
Also you would need a Unix (or Unix-like) OS. The following environment variables need to be set - GOBIN, GOPATH.

#### Use cases
The following make commands are supported:
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

### Kubernetes
#### Prerequisites
You need to have the following tools installed:
- kubectl
- k14s utilities, like ytt and kapp
You would need a Unix (or Unix-like) OS and a running Kubernetes cluster.

#### Use cases
You can deploy like so:
```bash
ytt -f chart | kapp deploy -a sg -f- --yes # installs storygraph to your k8s cluster
kubectl get deployments -n storygraph # view the storygraph deployment
kubectl -n storygraph port-forward service/storygraph-service 8080:8080 # forward port to localhost:8080
```
Now you should be able to access your storygraph instance like so:
```bash
curl localhost:8080/greet 
```
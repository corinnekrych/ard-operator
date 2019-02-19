# ADR Operator

ADR stands for [Architecture Decision Record](https://adr.github.io/). 
The concept is simple: let's make the decisions that driv your project explicit.
In this (simple) implementation of ADR, we stick to the KISS priciple:
- ARD is written in plain text or similar markdown, asciidoc.
- ARD is kept where it belongs: source code.

The intend of this repo is to provide an OpenShift operator that takes 
your project repositories, search for `docs` folder and deploy 
the documentation on OpenShit.

This repository was initially boostrapped using [CoreOS operator](https://github.com/operator-framework/operator-sdk). 

## Build

### Pre-requisites
- [operator-sdk v0.5.0](https://github.com/operator-framework/operator-sdk#quick-start) 
- [dep][dep_tool] version v0.5.0+.
- [git][git_tool]
- [go][go_tool] version v1.10+.
- [docker][docker_tool] version 17.03+.
- [kubectl][kubectl_tool] version v1.11.0+ or [oc] version 3.11
- Access to a kubernetes v.1.11.0+ cluster or openshift cluster version 3.11

### Build
```
make build
```

## Deploy

### Minishift


[dep_tool]:https://golang.github.io/dep/docs/installation.html
[git_tool]:https://git-scm.com/downloads
[go_tool]:https://golang.org/dl/
[docker_tool]:https://docs.docker.com/install/
[kubectl_tool]:https://kubernetes.io/docs/tasks/tools/install-kubectl/
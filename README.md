# ADR Operator

ADR stands for [Architecture Decision Record](https://adr.github.io/). 
The concept is simple: let's make the decisions that driv your project explicit.
In this (simple) implementation of ADR, we stick to the KISS principle:
- ARD is written in plain text or similar markdown, asciidoc.
- ARD is kept where it belongs: source code.

The intend of this repo is to provide an OpenShift operator that takes 
your project repositories, search for `docs` folder and deploy 
the documentation on OpenShit.

This repository was initially bootstrapped using [CoreOS operator](https://github.com/operator-framework/operator-sdk). 

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
## Deployment

### Set up Minishift (one-off)
* create a new profile to test the operator
```
minishift profile set adr-operator
```
* enable the admin-user add-on
```
minishift addon enable admin-user
```
* optionally, configure the VM 
```
minishift config set cpus 4
minishift config set memory 8GB
minishift config set vm-driver virtualbox
```
* start the instance
```
minishift start
```
* login with the admin account
```
oc login -u system:admin
```
* deploy RBAC
```
make deploy-rbac
```

### Deploy the CRD/Operator
#### Option 1: Quay.io
* You can build and push operator to a container hub
```
# login to Quay.io
docker login -u corinnekrych quay.io
# push the image to Docker Hub
docker push quay.io/corinnekrych/adr-operator
```
> NOTE: here we use quay.io, make sure once you've pushed your image that 
the image is visible as public(by default, it's pushed as private).
* Deploy to minishift
```
make deploy-operator
```
#### Option 2: Minishift container hub
* Or alternatively you can build and push to docker's minishift
```
eval $(minishift docker-env)
docker login -u developer -p $(oc whoami -t) $(minishift openshift registry)
operator-sdk build 172.30.1.1:5000/corinnekrych/adr-operator:v0.0.1
docker push 172.30.1.1:5000/corinnekrych/adr-operator:v0.0.1
```
> Note: change deploy/operator.yam image.
* Deploy to minishift
```
make deploy-operator
```

#### Option 3: Dev mode
In dev mode, no need to package in container and deploy, simply run your operator locally:
```
make local
```

### Deploy the CR for testing
* Make sure minishift is running and you're on myproject
```
oc project myproject
```
* clean all generated resources from previsous run
```
make clean
```
* Deploy CR 
```
make deploy-test
```
* Check the k8s resources were created:
```
oc get is,bc,svc,archdecisionrecord,build
NAME                                            DOCKER REPO                                TAGS      
imagestream.image.openshift.io/nodejs-output    172.30.1.1:5000/myproject/nodejs-output    latest
imagestream.image.openshift.io/nodejs-runtime   172.30.1.1:5000/myproject/nodejs-runtime   latest

NAME                                      TYPE      FROM         LATEST
buildconfig.build.openshift.io/myadr-bc   Source    Git@master   1

NAME                                        AGE
archdecisionrecord.corinnekrych.org/myadr   2m

NAME                                  TYPE      FROM          STATUS
build.build.openshift.io/myadr-bc-1   Source    Git@85ac14e   Complete
```

[dep_tool]:https://golang.github.io/dep/docs/installation.html
[git_tool]:https://git-scm.com/downloads
[go_tool]:https://golang.org/dl/
[docker_tool]:https://docs.docker.com/install/
[kubectl_tool]:https://kubernetes.io/docs/tasks/tools/install-kubectl/
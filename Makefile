VERSION     	?= unset
GITCOMMIT    	:= $(shell git rev-parse --short HEAD 2>/dev/null)
PROJECT_NAME 	:= adr-operator
OC_PROJECT      ?= myproject
BIN_DIR      	:= ./build/_output/bin
REPO_ORG     	:= corinnekrych
REPO_PATH    	:= github.com/$(REPO_ORG)/$(PROJECT_NAME)
PKGS         	:= $(shell go list  ./... | grep -v $(PROJECT)/vendor)
CONTAINER_CENTRAL_REPO := quay.io
M            	= $(shell printf "\033[34;1m▶\033[0m")
GO           	?= go
GOFMT        	?= $(GO)fmt

.PHONY: build
build: ; $(info $(M) Build operator docker images with $(CONTAINER_CENTRAL_REPO) )
	operator-sdk generate k8s
	operator-sdk generate openapi
	#operator-sdk build $(CONTAINER_CENTRAL_REPO)/$(REPO_ORG)/${PROJECT_NAME}

.PHONY: local
local: build deploy-crd; $(info $(M) Run Operator locally)
	@-oc new-project $(OC_PROJECT)
	operator-sdk up local --namespace=$(OC_PROJECT)

.PHONY: deploy-rbac
deploy-rbac: ; $(info $(M) Setup service account and deploy RBAC )
	oc create -f deploy/service_account.yaml
	oc create -f deploy/role.yaml
	oc create -f deploy/role_binding.yaml

.PHONY: deploy-crd
deploy-crd: ; $(info $(M) Deploy CRD )
	@-oc delete crd archdecisionrecords.corinnekrych.org
	@-oc create -f deploy/crds/corinnekrych_v1alpha1_archdecisionrecord_crd.yaml

.PHONY: deploy-operator
deploy-operator: deploy-crd ; $(info $(M) Deploy Operator )
	oc create -f deploy/operator.yaml

.PHONY: deploy-test
deploy-test: ; $(info $(M) Deploy a CR as testr )
	@-oc delete imagestream.image.openshift.io/nodejs-generated-xxxx buildconfig.build.openshift.io/myadr deploymentconfig.apps.openshift.io/myadr
	oc create -f deploy/crds/corinnekrych_v1alpha1_archdecisionrecord_cr.yaml

.PHONY: clean
clean: ; $(info $(M) Clean deployment )
	@-oc delete deployment adr-operator
	@-oc delete crd archdecisionrecords.corinnekrych.org
	@-oc delete pod

.PHONY: format
format: ; $(info $(M) Checking code style )
	@fmtRes=$$($(GOFMT) -d $$(find . -path ./vendor -prune -o -name '*.go' -print)); \
	if [ -n "$${fmtRes}" ]; then \
		echo "gofmt checking failed!"; echo "$${fmtRes}"; echo; \
		exit 1; \
	fi

.PHONY: lint
lint:
	golint $(PKGS)

.PHONY: test
test: ; $(info $(M) Running unit test )
	go test ./pkg/...
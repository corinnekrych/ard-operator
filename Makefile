VERSION     	?= unset
GITCOMMIT    	:= $(shell git rev-parse --short HEAD 2>/dev/null)
PROJECT_NAME 	:= adr-operator
BIN_DIR      	:= ./build/_output/bin
REPO_ORG     	:= corinnekrych
REPO_PATH    	:= github.com/$(REPO_ORG)/$(PROJECT_NAME)
PKGS         	:= $(shell go list  ./... | grep -v $(PROJECT)/vendor)
CONTAINER_CENTRAL_REPO := quay.io
M            	= $(shell printf "\033[34;1m▶\033[0m")
GO           	?= go
GOFMT        	?= $(GO)fmt

.PHONY: build
build: format ; $(info $(M) Build operator docker images with $(CONTAINER_CENTRAL_REPO) )
	operator-sdk generate k8s
	operator-sdk build $(CONTAINER_CENTRAL_REPO)/$(REPO_ORG)/${PROJECT_NAME}

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
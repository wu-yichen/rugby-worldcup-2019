export GOBIN ?= $(PWD)/bin
export PATH := $(GOBIN):$(PATH)
GOIMPORTS := $(GOBIN)/goimports
GINKGO := $(GOBIN)/ginkgo
GOCOV := $(GOBIN)/gocov
GOCOV_HTML := $(GOBIN)/gocov-html
GOLANGCI_LINT := $(GOBIN)/golangci-lint

.PHONY: all
all: test integration_test

COMMON_GINKGO_ARGS := -r -keepGoing --randomizeAllSpecs --randomizeSuites --failOnPending --trace --race --progress

.PHONY: test
test: tools
	$(GOBIN)/ginkgo --skip "\[INTEGRATION\]" $(COMMON_GINKGO_ARGS)

.PHONY: integration_test
integration_test: tools
	$(GOBIN)/ginkgo --focus "\[INTEGRATION\]" $(COMMON_GINKGO_ARGS)

.PHONY: cover
cover: tools
	mkdir -p coverage
	$(GOBIN)/gocov test ./... | gocov-html > coverage/index.html

sources = $(shell find . -name '*.go' -not -path './vendor/*')
.PHONY: goimports
goimports: tools
	$(GOBIN)/goimports -w $(sources)

.PHONY: lint
lint: tools
	$(GOBIN)/golangci-lint run $(ARGS)

.PHONY: clean
clean:
	rm -f $(GOIMPORTS) $(GINKGO) $(GOCOV) $(GOCOV_HTML) $(GOLANGCI_LINT)

.PHONY: release
release:
	mkdir -p $(CURDIR)/release
	cp $(CURDIR)/.env $(CURDIR)/release/
	cp $(CURDIR)/go.mod go.sum $(CURDIR)/release/
	CGO_ENABLED=1 GOOS=$(os) GOARCH=$(arch) go build -o release/wcwhen ./cli.go

$(GOIMPORTS):
	go install golang.org/x/tools/cmd/goimports

$(GINKGO):
	go install github.com/onsi/ginkgo/ginkgo

$(GOCOV):
	go install github.com/axw/gocov/gocov

$(GOCOV_HTML):
	go install github.com/matm/gocov-html

$(GOLANGCI_LINT):
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

tools: $(GOIMPORTS) $(GINKGO) $(GOCOV) $(GOCOV_HTML) $(GOLANGCI_LINT)

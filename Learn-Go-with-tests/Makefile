ifndef GO_MODULE
$(error GO_MODULE has not been defined)
endif

ifeq ($(PWD), $(GO_MODULE))
WD ::= $(PWD)
else
WD ::= $(realpath $(GO_MODULE))
endif

ifndef WD
$(error could not find module $(GO_MODULE))
endif


MODULE = $(notdir $(GO_MODULE))

prefix = podman run --interactive --tty --rm \
       -w/usr/local/go/src/$(GO_MODULE) \
       -v$(WD)/:/usr/local/go/src/$(GO_MODULE)/:Z \
       golang:alpine

run:
	$(prefix) go run $(MODULE)
test:
	$(prefix) go test $(MODULE)
build:
	$(prefix) go build -v $(MODULE)

sh:
	$(prefix) sh


prefix = podman run --interactive --tty --rm \
       -w/usr/local/go/src/app \
       -v$(PWD)/:/usr/local/go/src/app/:Z \
       golang:alpine

build:
	$(prefix) go build -v app
sh:
	$(prefix) sh

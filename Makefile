all:	xrpc feed

feed:
	go install -tags jwx_es256k ./...

xrpc:
	go install github.com/agentio/slink/cmd/slink-generate@latest
	$(shell go env GOPATH)/bin/slink-generate xrpc -i lexicons-bluesky -m xrpc.json

submodules:
	git submodule update --init --recursive
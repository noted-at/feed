all:	xrpc feed

feed:
	go install -tags jwx_es256k ./...

xrpc:
	slink generate xrpc -m xrpc.json -l debug

submodules:
	git submodule update --init --recursive
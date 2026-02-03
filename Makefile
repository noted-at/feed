all:
	go install -tags jwx_es256k ./...

xrpc:
	slink generate xrpc -m xrpc.json -l debug

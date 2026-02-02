all:
	go install ./...

xrpc:
	slink generate xrpc -m xrpc.json -l debug

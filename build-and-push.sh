#!/bin/sh

PLATFORMS=linux/amd64,linux/arm64

docker buildx build \
	--sbom=true \
	--push \
	--provenance=true \
	--platform $PLATFORMS \
	--tag docker.io/agentio/feed:latest \
	.


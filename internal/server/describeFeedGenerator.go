package server

import (
	"net/http"

	"github.com/noted-at/feed/gen/xrpc"
	"github.com/noted-at/feed/internal/common"
)

func describeFeedGeneratorHandler(w http.ResponseWriter, r *http.Request) {
	response := &xrpc.AppBskyFeedDescribeFeedGenerator_Output{
		Did: "did:plc:2aa6fahtis66ymq7rd2yxdwu",
		Feeds: []*xrpc.AppBskyFeedDescribeFeedGenerator_Feed{
			{
				Uri: "at://did:plc:2aa6fahtis66ymq7rd2yxdwu/app.bsky.feed.generator/static",
			},
		},
	}
	common.RespondWithJSON(w, response)
}

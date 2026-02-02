package server

import (
	"encoding/json"
	"net/http"

	"github.com/agentio/sting/pkg/auth"
	"github.com/agentio/sting/pkg/sting"
	"github.com/charmbracelet/log"
	"github.com/noted-at/feed/gen/xrpc"
)

func getFeedSkeletonHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query()
	log.Infof("feed: %s", p["feed"])
	log.Infof("limit: %s", p["limit"])
	claims, err := auth.Authenticate(r.Header.Get("authorization"))
	if err == nil {
		b, _ := json.MarshalIndent(claims, "", "  ")
		log.Infof("claims: %s", b)
	}
	response := &xrpc.AppBskyFeedGetFeedSkeleton_Output{
		Feed: []*xrpc.AppBskyFeedDefs_SkeletonFeedPost{
			{
				Post: "at://did:plc:bnr33h7nafe5nk4zzlshvana/app.bsky.feed.post/3mdjb3t3bfc25",
			},
			{
				Post: "at://did:plc:bnr33h7nafe5nk4zzlshvana/app.bsky.feed.post/3ltkivqtpfk2l",
			},
			{
				Post: "at://did:plc:bnr33h7nafe5nk4zzlshvana/app.bsky.feed.post/3lqo6xka3622g",
			},
		},
	}
	sting.RespondWithJSON(w, response)
}

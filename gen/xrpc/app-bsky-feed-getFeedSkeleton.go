// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Code produced by slink and slink itself are released under the AGPL.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.feed.getFeedSkeleton

import (
	"context"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyFeedGetFeedSkeleton_Description = "Get a skeleton of a feed provided by a feed generator. Auth is optional, depending on provider requirements, and provides the DID of the requester. Implemented by Feed Generator Service."

type AppBskyFeedGetFeedSkeleton_Output struct {
	Cursor *string                             `json:"cursor,omitempty"`
	Feed   []*AppBskyFeedDefs_SkeletonFeedPost `json:"feed,omitempty"`
	ReqId  *string                             `json:"reqId,omitempty"`
}

// Get a skeleton of a feed provided by a feed generator. Auth is optional, depending on provider requirements, and provides the DID of the requester. Implemented by Feed Generator Service.
func AppBskyFeedGetFeedSkeleton(ctx context.Context, c slink.Client, cursor string, feed string, limit int64) (*AppBskyFeedGetFeedSkeleton_Output, error) {
	var output AppBskyFeedGetFeedSkeleton_Output
	params := map[string]any{
		"cursor": cursor,
		"feed":   feed,
		"limit":  limit,
	}
	if err := c.Do(ctx, slink.Query, "", "app.bsky.feed.getFeedSkeleton", params, nil, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

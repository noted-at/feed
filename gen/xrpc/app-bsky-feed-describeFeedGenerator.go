// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Code produced by slink and slink itself are released under the AGPL.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.feed.describeFeedGenerator

import (
	"context"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyFeedDescribeFeedGenerator_Feed_Description = ""

// AppBskyFeedDescribeFeedGenerator_Feed is a record with Lexicon type app.bsky.feed.describeFeedGenerator#feed
type AppBskyFeedDescribeFeedGenerator_Feed struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Uri           string `json:"uri"`
}

const AppBskyFeedDescribeFeedGenerator_Links_Description = ""

// AppBskyFeedDescribeFeedGenerator_Links is a record with Lexicon type app.bsky.feed.describeFeedGenerator#links
type AppBskyFeedDescribeFeedGenerator_Links struct {
	LexiconTypeID  string  `json:"$type,omitempty"`
	PrivacyPolicy  *string `json:"privacyPolicy,omitempty"`
	TermsOfService *string `json:"termsOfService,omitempty"`
}

const AppBskyFeedDescribeFeedGenerator_Description = "Get information about a feed generator, including policies and offered feed URIs. Does not require auth; implemented by Feed Generator services (not App View)."

type AppBskyFeedDescribeFeedGenerator_Output struct {
	Did   string                                   `json:"did"`
	Feeds []*AppBskyFeedDescribeFeedGenerator_Feed `json:"feeds,omitempty"`
	Links *AppBskyFeedDescribeFeedGenerator_Links  `json:"links,omitempty"`
}

// Get information about a feed generator, including policies and offered feed URIs. Does not require auth; implemented by Feed Generator services (not App View).
func AppBskyFeedDescribeFeedGenerator(ctx context.Context, c slink.Client) (*AppBskyFeedDescribeFeedGenerator_Output, error) {
	var output AppBskyFeedDescribeFeedGenerator_Output
	params := map[string]any{}
	if err := c.Do(ctx, slink.Query, "", "app.bsky.feed.describeFeedGenerator", params, nil, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Code produced by slink and slink itself are released under the AGPL.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.feed.defs

import (
	"encoding/json"
	"fmt"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyFeedDefs_SkeletonFeedPost_Description = ""

// AppBskyFeedDefs_SkeletonFeedPost is a record with Lexicon type app.bsky.feed.defs#skeletonFeedPost
type AppBskyFeedDefs_SkeletonFeedPost struct {
	LexiconTypeID string                                   `json:"$type,omitempty"`
	FeedContext   *string                                  `json:"feedContext,omitempty"`
	Post          string                                   `json:"post"`
	Reason        *AppBskyFeedDefs_SkeletonFeedPost_Reason `json:"reason,omitempty"`
}

// AppBskyFeedDefs_SkeletonFeedPost_Reason is a union with these possible values:
// - AppBskyFeedDefs_SkeletonReasonRepost (app.bsky.feed.defs#skeletonReasonRepost)
// - AppBskyFeedDefs_SkeletonReasonPin (app.bsky.feed.defs#skeletonReasonPin)
type AppBskyFeedDefs_SkeletonFeedPost_Reason struct {
	Wrapper AppBskyFeedDefs_SkeletonFeedPost_Reason_Wrapper
}

// Value wrappers must conform to AppBskyFeedDefs_SkeletonFeedPost_Reason_Wrapper
type AppBskyFeedDefs_SkeletonFeedPost_Reason_Wrapper interface {
	isAppBskyFeedDefs_SkeletonFeedPost_Reason()
}

// AppBskyFeedDefs_SkeletonFeedPost_Reason__AppBskyFeedDefs_SkeletonReasonRepost wraps values of type *AppBskyFeedDefs_SkeletonReasonRepost
type AppBskyFeedDefs_SkeletonFeedPost_Reason__AppBskyFeedDefs_SkeletonReasonRepost struct {
	Value *AppBskyFeedDefs_SkeletonReasonRepost
}

func (t AppBskyFeedDefs_SkeletonFeedPost_Reason__AppBskyFeedDefs_SkeletonReasonRepost) isAppBskyFeedDefs_SkeletonFeedPost_Reason() {
}

// AppBskyFeedDefs_SkeletonFeedPost_Reason__AppBskyFeedDefs_SkeletonReasonPin wraps values of type *AppBskyFeedDefs_SkeletonReasonPin
type AppBskyFeedDefs_SkeletonFeedPost_Reason__AppBskyFeedDefs_SkeletonReasonPin struct {
	Value *AppBskyFeedDefs_SkeletonReasonPin
}

func (t AppBskyFeedDefs_SkeletonFeedPost_Reason__AppBskyFeedDefs_SkeletonReasonPin) isAppBskyFeedDefs_SkeletonFeedPost_Reason() {
}

func (u AppBskyFeedDefs_SkeletonFeedPost_Reason) MarshalJSON() ([]byte, error) {
	switch v := u.Wrapper.(type) {
	case AppBskyFeedDefs_SkeletonFeedPost_Reason__AppBskyFeedDefs_SkeletonReasonRepost:
		return slink.MarshalWithLexiconType("app.bsky.feed.defs#skeletonReasonRepost", v.Value)
	case AppBskyFeedDefs_SkeletonFeedPost_Reason__AppBskyFeedDefs_SkeletonReasonPin:
		return slink.MarshalWithLexiconType("app.bsky.feed.defs#skeletonReasonPin", v.Value)
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}
func (u *AppBskyFeedDefs_SkeletonFeedPost_Reason) UnmarshalJSON(data []byte) error {
	switch slink.LexiconTypeFromJSONBytes(data) {
	case "app.bsky.feed.defs#skeletonReasonRepost":
		var v AppBskyFeedDefs_SkeletonReasonRepost
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyFeedDefs_SkeletonFeedPost_Reason__AppBskyFeedDefs_SkeletonReasonRepost{Value: &v}
		return nil
	case "app.bsky.feed.defs#skeletonReasonPin":
		var v AppBskyFeedDefs_SkeletonReasonPin
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyFeedDefs_SkeletonFeedPost_Reason__AppBskyFeedDefs_SkeletonReasonPin{Value: &v}
		return nil
	default:
		return nil
	}
}

const AppBskyFeedDefs_SkeletonReasonPin_Description = ""

// AppBskyFeedDefs_SkeletonReasonPin is a record with Lexicon type app.bsky.feed.defs#skeletonReasonPin
type AppBskyFeedDefs_SkeletonReasonPin struct {
	LexiconTypeID string `json:"$type,omitempty"`
}

const AppBskyFeedDefs_SkeletonReasonRepost_Description = ""

// AppBskyFeedDefs_SkeletonReasonRepost is a record with Lexicon type app.bsky.feed.defs#skeletonReasonRepost
type AppBskyFeedDefs_SkeletonReasonRepost struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Repost        string `json:"repost"`
}

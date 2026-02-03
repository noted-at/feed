package server

import (
	"net/http"

	"github.com/agentio/slink/pkg/slink"
)

func getDidJSONHandler(w http.ResponseWriter, r *http.Request) {
	doc := &DidDocument{
		Context: "https://www.w3.org/ns/did/v1",
		Id:      "did:web:feed.noted.at",
		Service: []DidService{
			{
				ID:              "#bsky_fg",
				Type:            "BskyFeedGenerator",
				ServiceEndpoint: "https://feed.noted.at",
			},
		},
	}
	slink.RespondWithJSON(w, doc)
}

type DidDocument struct {
	Context string       `json:"@context"`
	Id      string       `json:"id"`
	Service []DidService `json:"service"`
}

type DidService struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

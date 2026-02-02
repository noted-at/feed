package server

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	b, err := json.Marshal(doc)
	if err != nil {
		http.Error(w, "oops", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, "%s\n", string(b))
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

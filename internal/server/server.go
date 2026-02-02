package server

import (
	"net"
	"net/http"

	"github.com/charmbracelet/log"
)

func Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/.well-known/did.json", getDidJSONHandler)
	mux.HandleFunc("/xrpc/app.bsky.feed.describeFeedGenerator", describeFeedGeneratorHandler)
	mux.HandleFunc("/xrpc/app.bsky.feed.getFeedSkeleton", getFeedSkeletonHandler)
	server := &http.Server{
		Handler: mux,
	}
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		return err
	}
	log.Infof("starting feed server")
	return server.Serve(listener)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	const indexHtml = "There's nothing to see here.\n"
	if r.URL.Path == "/" && r.Method == "GET" {
		_, _ = w.Write([]byte(indexHtml))
		return
	}
	http.Error(w, "Not found", http.StatusNotFound)
}

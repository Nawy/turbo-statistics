package main

import (
	"net/http"
	controller "turbo-statistics/controller"
)

func main() {
	http.HandleFunc("/health", controller.HealthHandler)
	http.HandleFunc("/post/vote", controller.PostVoteHandler)
	http.HandleFunc("/post/meta", controller.PostUpdateMetaHandler)
	http.ListenAndServe(":8080", nil)
}

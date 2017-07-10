package controller

import (
	"net/http"
	HTTPUtils "turbo-statistics/utils"
)

// PostVoteHandler post upvote
func PostVoteHandler(w http.ResponseWriter, r *http.Request) {
	if !HTTPUtils.IsHttpMethod(r, HTTPUtils.POST) {
		HTTPUtils.WriteBadGateway(w, "You should use POST method")
		return
	}
}

// PostUpdateMetaHandler post update meta information about post
func PostUpdateMetaHandler(w http.ResponseWriter, r *http.Request) {
	if !HTTPUtils.IsHttpMethod(r, HTTPUtils.POST) {
		HTTPUtils.WriteBadGateway(w, "You should use POST method")
		return
	}
}

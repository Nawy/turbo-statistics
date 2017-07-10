package controller

import (
	"net/http"
	HTTPUtils "turbo-statistics/utils"
)

// HealthHandler return status of service
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if !HTTPUtils.IsHttpMethod(r, HTTPUtils.GET) {
		HTTPUtils.WriteBadGateway(w, "You should use GET method")
		return
	}

}

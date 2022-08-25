package server

import (
	"github.com/golang/gddo/httputil/header"
	"net/http"
)

func CanAccept(r *http.Request, values ...string) bool {
	for _, spec := range header.ParseAccept(r.Header, "Accept") {
		for _, value := range values {
			if spec.Value == value || spec.Value == "*/*" {
				return true
			}
		}
	}

	return false
}

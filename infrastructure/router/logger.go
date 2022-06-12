package router

import (
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

type httpReqInfo struct {
	method       string
	uri          string
	referer      string
	userAgent    string
	realIP       string
	forwardedFor string
	remoteAddr   string
}

func logHandler(h http.Handler) http.Handler {
	fn := func(res http.ResponseWriter, req *http.Request) {
		ri := httpReqInfo{
			method:       req.Method,
			uri:          req.URL.String(),
			referer:      req.Header.Get("Referer"),
			userAgent:    req.Header.Get("User-Agent"),
			realIP:       req.Header.Get("X-Real-Ip"),
			forwardedFor: req.Header.Get("X-Forwarded-For"),
			remoteAddr:   req.RemoteAddr,
		}
		log.Println(spew.Sprintf("%v", ri))

		h.ServeHTTP(res, req)
	}

	return http.HandlerFunc(fn)
}

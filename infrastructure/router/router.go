package router

import (
	"fmt"
	"go-bootcamp/config"
	"go-bootcamp/interface/controller"
	"log"
	"net/http"
	"time"

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
	fn := func(w http.ResponseWriter, req *http.Request) {
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

		h.ServeHTTP(w, req)
	}
	return http.HandlerFunc(fn)

}

func Init(c controller.AppController) (server *http.Server, address string) {
	mux := &http.ServeMux{}
	mux.HandleFunc("/pokemon", c.Pokemon.GetPokemon)
	mux.HandleFunc("/pokemon/", c.Pokemon.GetPokemonById)
	mux.HandleFunc("/create-pokemon", c.Pokemon.CreatePokemon)

	handler := logHandler(mux)

	port := config.Config.Server.Port
	host := config.Config.Server.Host

	address = fmt.Sprintf("%v:%v", host, port)
	server = &http.Server{
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
		Addr:         address,
	}

	return
}

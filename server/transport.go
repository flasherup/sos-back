package server

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPTSTransport(logger log.Logger, staticFolder string) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Methods("GET").Path("/eth/calculate/{" + KeyAccount + "}/").HandlerFunc(CalculateHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticFolder)))
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	});
}

package utils

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var logger = func(method, uri, name string, start time.Time) {
	log.Printf("\"method\":%q  \"uri\":%q  \"name\":%q  \"time\":%q", method, uri, name, time.Since(start))
}

func Logging(h httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		h(w, r, ps)
		logger(r.Method, r.URL.Path, name, start)
	}
}

func CommonHeaders(h httprouter.Handle, name string) httprouter.Handle {
	return Logging(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; chartset=UTF-8")
		h(w, r, ps)
	}, name)
}

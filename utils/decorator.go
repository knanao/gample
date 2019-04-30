package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

func IdShouldBeInt(h httprouter.Handle, name string) httprouter.Handle {
	return CommonHeaders(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idParam := ps.ByName("todoId")
		_, err := strconv.Atoi(idParam)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(500)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				return
			}
			return
		}
		h(w, r, ps)
	}, name)
}

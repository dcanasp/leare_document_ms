package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Main(port int) {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	k := r.PathPrefix("/read").Subrouter()
	s := r.PathPrefix("/create").Subrouter()
	l := r.PathPrefix("/update").Subrouter()
	m := r.PathPrefix("/delete").Subrouter()

	Get(k)
	Create(s)
	Create(l)
	Delete(m)

	// http.ListenAndServe(":3012", r)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)

}

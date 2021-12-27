package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)


func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Sprintln(r)
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars[KeyAccount])
}

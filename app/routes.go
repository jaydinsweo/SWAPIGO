package apiroute

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// AppRouter function handle the api router
// When access path  localhost:8080/people, .localhost:8080/planets
// it will print out a text
func AppRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/people", GetPeople)
	return r
}

// GetPeople returns a text
func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get People")

}

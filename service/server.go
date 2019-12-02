package server

import (
	"net/http"
	"newgo/test/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	routes.Addroutes(r)
	http.ListenAndServe(":8001", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r))

}

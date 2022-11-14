package router

import (
	"dbms/configuration"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	Mux = mux.NewRouter()
)

// Router class
type router struct{}

func NewRouter() RouterInterface {
	return &router{}
}

// General post method
func (r *router) Post(url string, handle http.HandlerFunc) {
   Mux.HandleFunc(url,handle).Methods("POST")
}

// General get method
func (r *router) Get(url string, handle http.HandlerFunc) {
   Mux.HandleFunc(url,handle).Methods("GET")
}

// General server method
func (server *router) Serve(port string, ctx *configuration.Context) {
	application := &http.Server{
		Addr:    port,
		Handler: Mux,
	}

	application.ListenAndServe()
}

// General prefixing method
func (r *router) PathPrefix(url string) *mux.Router{
	return Mux.PathPrefix(url).Subrouter()
}

// Cors handling
func (r *router) Cors(f http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	})

	return corsHandler.Handler(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

	}))
}

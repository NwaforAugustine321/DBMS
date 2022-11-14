package router

import (
	"dbms/configuration"
	"net/http"

	"github.com/gorilla/mux"
)


type RouterInterface interface{
	Get(url string, handle  http.HandlerFunc)
	Post(url string, handle http.HandlerFunc)
	Serve(port string, ctx *configuration.Context)
	PathPrefix(url string) *mux.Router
	Cors(f http.Handler) http.Handler
}

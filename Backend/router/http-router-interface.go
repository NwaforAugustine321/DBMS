package router

import (
	"dbms/configuration"
	"net/http"
)


type RouterInterface interface{
	Get(url string, handle  http.HandlerFunc)
	Post(url string, handle http.HandlerFunc)
	Serve(port string, ctx *configuration.Context)
	PathPrefix(url string,prefix string,handle http.HandlerFunc)
	Cors(f http.Handler) http.Handler
}

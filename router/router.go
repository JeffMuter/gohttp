package router

import (
	"gohttp/pages"
	"net/http"
)

func Router() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", pages.ServeHomepage)

	return mux

}

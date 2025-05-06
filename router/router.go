package router

import (
	"gohttp/pages"
	"net/http"
)

func Router() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", pages.ServeHomepage)
	mux.HandleFunc("/getForm", pages.ServeForm)
	return mux
}

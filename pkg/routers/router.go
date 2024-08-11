package routers

import (
	"net/http"

	"github.com/url-shortner/pkg/controller"
)

var RegisterRouter = func(mux *http.ServeMux) {

	mux.HandleFunc("GET /heathcheck/", controller.Heathcheck)
	mux.HandleFunc("GET /", controller.Resolve)
	mux.HandleFunc("POST /shorten/", controller.Shorten)

}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/url-shortner/pkg/routers"
)

func main() {
	mux := http.NewServeMux()

	routers.RegisterRouter(mux)

	fmt.Println("Server has started at port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))

}

package main

import (
	"github.com/hadv/morilla/api"

	"net/http"
)

func main() {
	router := api.NewMovieServer()
	http.ListenAndServe(":8080", router)
}

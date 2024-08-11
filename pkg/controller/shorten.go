package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ShortenRequest struct {
	LongUrl string `json:"longurl"`
	Expire  string `json:"expire"`
}

func Heathcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "tiny url service is up")
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the controller")
	w.Header().Set("Content-type", "application/json")

	var data ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid Payload", http.StatusBadRequest)
		return
	}
	fmt.Printf("Inside the controller %v", data)

	json.NewEncoder(w).Encode(data)
}

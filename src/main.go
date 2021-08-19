package main

import (
	"encoding/json"
    "net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
    http.HandleFunc("/", ServerResponse)
    http.ListenAndServe(":3000", nil)
}

func ServerResponse(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Léo", []string{"play instruments", "programming", "youtuber"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
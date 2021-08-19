package main

import (
	"encoding/json"
	"net/http"
)

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

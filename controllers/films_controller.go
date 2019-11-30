package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	crawler "github.com/M3nthis/space-cinema-api/services"
)

// GetFilms takes today’s movie schedule
// and returns it in Json format
func GetFilms(w http.ResponseWriter, r *http.Request) {
	elencoFilms, err := crawler.SearchTitles(`https://www.thespacecinema.it/i-nostri-cinema/vimercate/al-cinema`, `#loader-target`)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(*elencoFilms)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))
}

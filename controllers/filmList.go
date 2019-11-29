package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/M3nthis/space-cinema-api/models"
	crawler "github.com/M3nthis/space-cinema-api/utils"
)

// FilmList takes today’s movie schedule
// and returns it in Json format
func FilmList(w http.ResponseWriter, r *http.Request) {
	elencoFilms, err := crawler.SearchTitles(`https://www.thespacecinema.it/i-nostri-cinema/vimercate/al-cinema`, `#loader-target`)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var titoli []Film
	for _, titolo := range elencoFilms {
		film := Film{Nome: titolo}
		titoli = append(titoli, film)
	}
	resp, err := json.Marshal(titoli)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))
}

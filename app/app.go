package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/M3nthis/space-cinema-api/controllers"
)

func StartApp() {
	fmt.Println("Server starting...")
	http.HandleFunc("/film-list", controllers.GetFilms)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

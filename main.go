package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/M3nthis/space-cinema-api/controllers"
)

func main() {
	fmt.Println("Server starting...")
	http.HandleFunc("/film-list", controllers.FilmList)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

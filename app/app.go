package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/M3nthis/space-cinema-api/controllers"
	"github.com/M3nthis/space-cinema-api/domain"
	"github.com/chromedp/chromedp"
)

// StartApp starts the app
func StartApp() {
	Ctx, cancel := chromedp.NewContext(context.Background())
	domain.Ctx = &Ctx
	defer cancel()

	fmt.Println("Server starting...")
	http.HandleFunc("/film-list", controllers.GetFilms)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

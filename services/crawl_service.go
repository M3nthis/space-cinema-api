package services

import (
	"fmt"
	"strings"

	"github.com/M3nthis/space-cinema-api/domain"
	. "github.com/M3nthis/space-cinema-api/domain"
	"github.com/chromedp/chromedp"
)

// SearchTitles carica il contenuto portato dall'url
// e lo inserisce in una stringa
func SearchTitles(url, target string) (titles *[]Film, err error) {
	res, err := crawl(url, target)
	if err != nil {
		return nil, err
	}

	el := strings.Split(res, "\n")
	var titlesSlice []Film
	titlesSlice = append(titlesSlice, Film{Nome: el[0]})
	for i, text := range el {
		if isTitle(text, i, &el) {
			titlesSlice = append(titlesSlice, Film{Nome: el[i+1]})
		} else if isHour(text) {
			titlesSlice[len(titlesSlice)-1].Orari += fmt.Sprintf("%s ", text)
		} else if isLength(text) {
			titlesSlice[len(titlesSlice)-1].Durata = formatLength(text)
		}
	}
	titles = &titlesSlice
	err = nil

	return
}

func crawl(url, target string) (string, error) {
	// Ensure the first tab is created
	if err := chromedp.Run(*domain.Ctx); err != nil {
		return "", err
	}

	// Create Tab
	ctx2, cancel := chromedp.NewContext(*domain.Ctx)
	defer cancel()

	// Run task
	var res string
	err := chromedp.Run(ctx2,
		chromedp.Navigate(url),
		chromedp.Text(target, &res, chromedp.NodeVisible, chromedp.ByID),
	)
	if err != nil {
		return "", err
	}

	return res, nil
}

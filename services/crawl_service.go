package utils

import (
	"context"
	"strings"

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
		if text == "Scheda film" && i+1 < len(el) && el[i+1] != "TOP" {
			titlesSlice = append(titlesSlice, Film{Nome: el[i+1]})
		}
		if text == "TOP" && i+1 < len(el) {
			titlesSlice = append(titlesSlice, Film{Nome: el[i+1]})
		}
	}
	titles = &titlesSlice
	err = nil

	return
}

func crawl(url, target string) (string, error) {
	// Create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Run task
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Text(target, &res, chromedp.NodeVisible, chromedp.ByID),
	)
	if err != nil {
		return "", err
	}

	return res, nil
}

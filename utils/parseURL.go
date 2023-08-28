package utils

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/pterm/pterm"
)

func ParseURL(url string, comprehensive bool) (string, error) {

	var baseUrl string
	var body string

	baseUrl, _ = CleanURL(&url)

	fmt.Println("")
	pterm.DefaultBasicText.Println("Base URL: ", pterm.LightCyan(baseUrl))
	pterm.DefaultBasicText.Println("Entered URL: ", pterm.LightCyan(url))
	pterm.DefaultBasicText.Println("Comprehensive Search: ", pterm.LightCyan(comprehensive))

	c := colly.NewCollector(
	// colly.AllowedDomains(domain),
	)

	c.OnHTML("body", func(e *colly.HTMLElement) {
		body = e.Text
	})

	if comprehensive {
		// TODO
		err := c.Visit(baseUrl)
		if err != nil {
			return "", err
		}

	} else {
		err := c.Visit(url)
		if err != nil {
			return "", err
		}
	}

	return body, nil
}

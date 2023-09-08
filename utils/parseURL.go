package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/pterm/pterm"
)

func ParseURL(url string, comprehensive bool, depth int) (string, error) {

	var baseUrl string
	var body string
	var currBody string

	baseUrl, domain := CleanURL(&url)
	_ = os.Remove("logs.txt")

	pterm.DefaultBasicText.Println("Base URL: ", pterm.LightCyan(baseUrl))
	pterm.DefaultBasicText.Println("Entered URL: ", pterm.LightCyan(url))
	pterm.DefaultBasicText.Println("Comprehensive Search: ", pterm.LightCyan(comprehensive))

	if !comprehensive {
		fmt.Println("")

		c := colly.NewCollector()

		c.OnHTML("body", func(e *colly.HTMLElement) {
			body = strings.ReplaceAll(strings.ReplaceAll(e.Text, "\t", ""), "\n", "")
		})

		// TODO
		err := c.Visit(baseUrl)
		if err != nil {
			return "", err
		}

	} else {
		pterm.DefaultBasicText.Println("Search Depth: ", pterm.LightCyan(depth))

		domainsVisited := 0
		fmt.Println(domain)

		c := colly.NewCollector(
			colly.AllowedDomains(domain),
		)

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong:", err)
		})

		c.OnResponse(func(r *colly.Response) {
			domainsVisited++
			fmt.Println("Visited", r.Request.URL)
		})

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			if domainsVisited == depth {
				return
			}

			e.Request.Visit(e.Attr("href"))
		})

		c.OnHTML("body", func(e *colly.HTMLElement) {
			currBody = e.Text
			body += e.Text

			file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				log.Fatal(err)
			}

			defer file.Close()

			file.Write([]byte(e.Request.URL.Path))
			file.Write([]byte(currBody))
			file.Write([]byte("\n\n\n\n\n"))
		})

		if err := c.Visit(url); err != nil {
			return "", err
		}
	}

	return body, nil
}

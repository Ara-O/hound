package utils

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ParseURL(url string, comprehensive bool) (string, error) {

	var baseUrl string
	var body string

	baseUrl, domain := CleanURL(&url)

	fmt.Println("Base Url -", baseUrl)
	fmt.Println("Inputted Url -", url)
	fmt.Println("Domain -", domain)
	fmt.Println("Comprehensive -", comprehensive)

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.OnHTML("body", func(e *colly.HTMLElement) {
		body = e.Text
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	if comprehensive {
		c.Visit(baseUrl)
	} else {
		c.Visit(url)
	}

	return body, nil
}

package utils

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ParseURL(url string, comprehensive bool) (string, error) {

	var baseUrl string

	baseUrl, domain := CleanURL(&url)

	fmt.Println("Base Url -", baseUrl)
	fmt.Println("Inputted Url -", url)
	fmt.Println("Domain -", domain)
	fmt.Println("Comprehensive -", comprehensive)

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	// c.OnHTML("h1", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	if comprehensive {
		c.Visit(baseUrl)
	} else {
		c.Visit(url)
	}

	return "", nil
}

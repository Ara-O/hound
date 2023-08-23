package utils

import "fmt"

func ParseURL(url string, comprehensive bool) (string, error) {

	var baseUrl string

	if comprehensive {
		baseUrl = CleanURL(url)
	}

	fmt.Println("Base Url -", baseUrl)
	// c := colly.NewCollector(
	// 	colly.AllowedDomains("react.dev"),
	// )

	// // Find and visit all links
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })

	// c.OnHTML("h1", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	// c.Visit("https://react.dev")

	return "", nil
}

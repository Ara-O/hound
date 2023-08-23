package utils

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func CleanURL(urlParameter *string) (string, string) {
	//Should accept
	//https://www.araoladipo.tech
	//www.araoladipo.tech
	//https://araoladipo.tech
	//araoladipo.tech

	if !strings.HasPrefix(*urlParameter, "https://") && !strings.HasPrefix(*urlParameter, "http://") {
		*urlParameter = "https://" + *urlParameter
	}

	*urlParameter = strings.ReplaceAll(*urlParameter, "www.", "")

	// Parse the URL.
	parsedURL, err := url.Parse(*urlParameter)

	if err != nil {
		log.Fatal("Error parsing url")
	}

	// Build the base URL with scheme and host.
	baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)

	domain := parsedURL.Host

	if strings.HasPrefix(parsedURL.Host, "www.") {
		domain = strings.Replace(parsedURL.Host, "www.", "", 1)
	}

	return baseURL, domain
}

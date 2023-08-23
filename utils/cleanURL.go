package utils

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func CleanURL(urlParamater string) string {
	//Should accept
	//https://www.araoladipo.tech
	//www.araoladipo.tech
	//https://araoladipo.tech
	//araoladipo.tech

	if !strings.HasPrefix(urlParamater, "https://") && !strings.HasPrefix(urlParamater, "http://") {
		urlParamater = "https://" + urlParamater
	}

	// Parse the URL.
	parsedURL, err := url.Parse(urlParamater)

	if err != nil {
		log.Fatal("Error parsing url")
	}

	// Build the base URL with scheme and host.
	baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)

	return baseURL
}

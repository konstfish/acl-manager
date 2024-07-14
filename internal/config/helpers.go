package config

import (
	"net/url"
	"strings"
)

func isFullURL(str string) bool {
	parsedURL, err := url.Parse(str)
	if err != nil {
		return false
	}

	return parsedURL.Scheme != "" && parsedURL.Host != ""
}

func isDomain(str string) bool {
	parsedURL, err := url.Parse(str)
	if err != nil {
		return false
	}

	return parsedURL.Scheme == "" && strings.Contains(str, ".")
}

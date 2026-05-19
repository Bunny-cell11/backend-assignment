package utils

import (
	"errors"
	"net/url"
)

const MaxURLs = 20

func ValidateURL(input string) bool {

	parsed, err := url.ParseRequestURI(input)

	if err != nil {
		return false
	}

	return parsed.Scheme == "http" || parsed.Scheme == "https"
}

func ValidateMediaURLs(urls []string) error {

	if len(urls) > MaxURLs {
		return errors.New("too many URLs")
	}

	for _, u := range urls {

		if len(u) > 2048 {
			return errors.New("URL too long")
		}

		if !ValidateURL(u) {
			return errors.New("invalid URL")
		}
	}

	return nil
}

package url

import (
	"net/url"
)

func URLEncoder(urls string) string {
	encodeurl := url.QueryEscape(urls)
	return encodeurl
}

func URLDecoder(uri string) (string, error) {
	decodeurl, err := url.QueryUnescape(uri)
	if err != nil {
		return "", err
	}
	return decodeurl, nil
}

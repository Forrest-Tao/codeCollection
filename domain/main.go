package main

import (
	"fmt"
	"net/url"
	"strings"
)

func isValidHTTPURL(str string) bool {
	parsedURL, err := url.Parse(str)
	if err != nil {
		return false
	}
	return parsedURL.Scheme == "http" || parsedURL.Scheme == "https"
}

// getMainDomain 从 URL 中提取主域名
func getMainDomain(str string) (string, error) {
	parsedURL, err := url.Parse(str)
	if err != nil {
		return "", err
	}

	host := parsedURL.Hostname()
	parts := strings.Split(host, ".")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid URL, cannot determine main domain")
	}

	mainDomain := strings.Join(parts[len(parts)-2:], ".")
	return mainDomain, nil
}

const (
	blackListKey = "blackList:targetURL:%s"
)

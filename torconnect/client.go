package main

import (
	"net/http"
	"net/url"
	"golang.org/x/net/proxy"
)

func NewClient(addrs string) (*http.Client, error) {
	url, err := url.Parse("socks5://" + addrs)
	if err != nil {
		return nil, err
	}

	proxyUrl, err := proxy.FromURL(url, proxy.Direct)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{Dial: proxyUrl.Dial}

	return &http.Client{Transport: transport}, nil
}

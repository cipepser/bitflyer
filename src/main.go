package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Client struct {
	Url        *url.URL
	HTTPClient *http.Client

	Username, Password, Coin string
}

func NewClient(urlStr, username, password, coin string) (*Client, error) {
	if len(username) == 0 {
		return nil, errors.New("missing username")
	}
	if len(password) == 0 {
		return nil, errors.New("missing password")
	}
	if len(coin) == 0 {
		return nil, errors.New("missing coin")
	}

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", urlStr)
	}

	c := new(Client)
	c.Url = urlStr
	c.Username = username
	c.Password = password
	c.Coin = coin

	return &c
}

func main() {
	Url, err := url.Parse("https://example.com")
	c := NewClient(&Url, "hoge", "fuga", "btc")

	fmt.Printf("%s, %s, %s", c.Username, c.Password, c.Coin)
}

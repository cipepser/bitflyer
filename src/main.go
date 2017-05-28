package main

import "github.com/cipepser/bitflyer/src/sdk"

const (
	// URL is a end point of bitflyer api.
	URL = "https://api.bitflyer.jp"
)

func main() {
	c, _ := sdk.NewClient(URL, "user", "passwd", nil)

}

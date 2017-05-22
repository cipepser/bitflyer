package sdk

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/cipepser/bitflyer/src/sdk"
)

const (
	// URL is a end point of bitflyer api.
	URL     = "https://api.bitflyer.jp"
	timeout = 10
)

// ************** public API **************

// Board is a json struct for market board information.
type Board struct {
	MidPrice float64 `json:"mid_price"`
	Bids     []struct {
		Price float64 `json:"price"`
		Size  float64 `json:"size"`
	} `json:"bids"`
	Asks []struct {
		Price float64 `json:"price"`
		Size  float64 `json:"size"`
	} `json:"asks"`
}

// GetBoard gets makert board information.
func GetBoard(prodcut string) Board {
	c, _ := NewClient(URL, "user", "passwd", nil)

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	method := "GET"
	spath := "/v1/getboard"
	req, err := c.NewRequest(ctx, method, spath, nil)
	if err != nil {
		log.Fatal(err)
	}

	values := url.Values{}
	values.Add("product_code", prodcut)
	req.URL.RawQuery = values.Encode()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	board := Board{}
	err = DecodeBody(resp, &board)
	if err != nil {
		log.Fatal(err)
	}

	return board
}

// ************** private API **************

// Collateral is a json struct for private collateral information.
type Collateral struct {
	Collateral        float64 `json:"collateral"`
	OpenPositionPnl   float64 `json:"open_position_pnl"`
	RequireCollateral float64 `json:"require_collateral"`
	KeepRate          float64 `json:"keep_rate"`
}

// GetCollateral gets your private collateral information.
func GetCollateral() Collateral {
	c, _ := NewClient(URL, "user", "passwd", nil)

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	method := "GET"
	spath := "/v1/me/getcollateral"
	req, err := c.NewRequest(ctx, method, spath, nil)
	if err != nil {
		log.Fatal(err)
	}

	SetPrivateHeader(req, method, spath)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	collateral := Collateral{}
	err = sdk.DecodeBody(resp, &collateral)
	if err != nil {
		log.Fatal(err)
	}

	return collateral
}

// SetPrivateHeader sets authentication header to req.
// TODO: implement an authentication with body.
func SetPrivateHeader(req *http.Request, method, spath string) {
	key := os.Getenv("BFKEY")
	secret := os.Getenv("BFSECRET")

	timestamp := strconv.Itoa(int(time.Now().Unix()))
	sign := MakeHMAC(timestamp+method+spath, secret)

	req.Header.Set("ACCESS-KEY", key)
	req.Header.Set("ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("ACCESS-SIGN", sign)
	req.Header.Set("Content-Type", "application/json")
}

// MakeHMAC returns a HMAC by sha256.
func MakeHMAC(msg, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}

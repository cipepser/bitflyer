package sdk

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cipepser/bitflyer/src/sdk"
)

const (
	URL     = "https://api.bitflyer.jp"
	timeout = 10
)

// public API
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

func GetBoard() {

}

// private API
type Collateral struct {
	Collateral        float64 `json:"collateral"`
	OpenPositionPnl   float64 `json:"open_position_pnl"`
	RequireCollateral float64 `json:"require_collateral"`
	KeepRate          float64 `json:"keep_rate"`
}

func GetCollateral() {
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

	fmt.Println(resp)

	collateral := Collateral{}
	err = sdk.DecodeBody(resp, &collateral)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(collateral)

}

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

func MakeHMAC(msg, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}

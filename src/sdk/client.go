package sdk

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Client struct {
	URL                *url.URL
	HTTPClient         *http.Client
	Username, Password string
	Logger             *log.Logger
}

func NewClient(urlStr, username, password string, logger *log.Logger) (*Client, error) {
	if len(username) == 0 {
		return nil, errors.New("missing username")
	}
	if len(password) == 0 {
		return nil, errors.New("missing password")
	}

	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", urlStr)
	}

	c := &Client{
		URL:        parsedURL,
		HTTPClient: &http.Client{},
		Username:   username,
		Password:   password,
		Logger:     logger,
	}

	return c, err
}

func (c *Client) NewRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	// Basic認証
	// req.SetBasicAutho(c.Username, c.Password)

	// ヘッダ情報
	// if method == "POST" {
	key := os.Getenv("BFKEY")
	secret := os.Getenv("BFSECRET")

	timestamp := strconv.Itoa(int(time.Now().Unix()))
	sign := MakeHMAC(timestamp+method+spath, secret)

	req.Header.Set("ACCESS-KEY", key)
	req.Header.Set("ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("ACCESS-SIGN", sign)
	req.Header.Set("Content-Type", "application/json")
	// }

	return req, nil
}

func DecodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	return dec.Decode(out)
}

func MakeHMAC(msg, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}

func GetBoard() {

}

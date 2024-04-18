package repository

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/theborzet/connection_project/pkg/common/config"
)

type Repository interface {
	GetAccount() (*Account, error)
}

type MyAPI struct {
	Api    *http.Client
	Config *config.Config
}

// func NewRepository(api *http.Client) *MyAPI {
// 	return &MyAPI{
// 		Api: api,
// 	}
// }

func (ap MyAPI) Signature(timestamp string, api *MyAPI) string {
	preHashString := timestamp + api.Config.ApiKey + api.Config.RecvWindow
	h := hmac.New(sha256.New, []byte(api.Config.ApiSecret))
	h.Write([]byte(preHashString))
	return hex.EncodeToString(h.Sum(nil))
}

func (ap MyAPI) Timestamp() string {
	return fmt.Sprint(time.Now().UnixMilli())
}

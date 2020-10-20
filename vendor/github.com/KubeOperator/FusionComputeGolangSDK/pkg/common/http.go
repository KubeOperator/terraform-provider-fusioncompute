package common

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func NewHttpClient() *resty.Client {
	r := resty.New()
	r.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return r
}

func FormatHttpError(resp *resty.Response) error {
	text := fmt.Sprintf("code: %d,msg: %s", resp.StatusCode(), string(resp.Body()))
	return errors.New(text)
}

package client

import (
	"errors"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
	"github.com/go-resty/resty/v2"
)

type Session string

type FusionComputeClient interface {
	Connect() error
	DisConnect() error
	SetSession(token string)
	GetSession() Session
	GetHost() string
	GetUser() string
	GetPassword() string
	GetApiClient() (*resty.Client, error)
}

func NewFusionComputeClient(host string, user string, password string) FusionComputeClient {
	return &fusionComputeClient{
		user:     user,
		password: password,
		host:     host,
	}
}

type fusionComputeClient struct {
	session  Session
	user     string
	password string
	host     string
}

func (f *fusionComputeClient) SetSession(token string) {
	f.session = Session(token)
}

func (f *fusionComputeClient) GetSession() Session {
	return f.session
}

func (f *fusionComputeClient) Connect() error {
	a := NewAuth(f)
	err := a.Login()
	if err != nil {
		return err
	}
	return nil
}

func (f *fusionComputeClient) DisConnect() error {
	a := NewAuth(f)
	err := a.Logout()
	if err != nil {
		return err
	}
	return nil
}
func (f *fusionComputeClient) GetHost() string {
	return f.host
}
func (f *fusionComputeClient) GetUser() string {
	return f.user
}
func (f *fusionComputeClient) GetPassword() string {
	return f.password
}

func (f *fusionComputeClient) GetApiClient() (*resty.Client, error) {
	r := common.NewHttpClient()
	if f.GetSession() == "" {
		return nil, errors.New("no session exists,please login and try it again")
	}
	f.setDefaultHeader(r)
	r.SetHeader(XAuthToken, string(f.GetSession())).
		SetHostURL(f.host)
	return r, nil
}

func (f *fusionComputeClient) setDefaultHeader(client *resty.Client) {
	client.SetHeaders(map[string]string{
		"Accept":          "application/json;version=v8.0;charset=UTF-8;",
		"Accept-Language": "zh_CN:1.0",
	})
}

package site

import (
	"encoding/json"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
)

const (
	siteUrl = "/service/sites"
)

type Interface interface {
	ListSite() ([]Site, error)
	GetSite(siteUri string) (*Site, error)
}

func NewManager(client client.FusionComputeClient) Interface {
	return &manager{client: client}
}

type manager struct {
	client client.FusionComputeClient
}

func (m *manager) GetSite(siteUri string) (*Site, error) {
	var site Site
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().Get(siteUri)
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		err := json.Unmarshal(resp.Body(), &site)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, common.FormatHttpError(resp)
	}

	return &site, nil
}

func (m *manager) ListSite() ([]Site, error) {
	var sites []Site
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().Get(siteUrl)
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		var listSiteResponse ListSiteResponse
		err := json.Unmarshal(resp.Body(), &listSiteResponse)
		if err != nil {
			return nil, err
		}
		sites = listSiteResponse.Sites
	} else {
		return nil, common.FormatHttpError(resp)
	}
	return sites, nil
}

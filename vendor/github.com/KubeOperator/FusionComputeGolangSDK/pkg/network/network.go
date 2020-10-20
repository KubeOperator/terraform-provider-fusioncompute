package network

import (
	"encoding/json"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
	"path"
	"strings"
)

const (
	siteMask    = "<site_uri>"
	dvSwitchUrl = "<site_uri>/dvswitchs"
)

type Manager interface {
	ListDVSwitch() ([]DVSwitch, error)
	ListPortGroup(dvSwitchIdUri string) ([]PortGroup, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) ListPortGroup(dvSwitchIdUri string) ([]PortGroup, error) {

	var portGroups []PortGroup
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().Get(path.Join(dvSwitchIdUri, "portgroups"))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		var listPortGroupResponse ListPortGroupResponse
		err := json.Unmarshal(resp.Body(), &listPortGroupResponse)
		if err != nil {
			return nil, err
		}
		portGroups = listPortGroupResponse.PortGroups
	} else {
		return nil, common.FormatHttpError(resp)
	}
	return portGroups, nil
}

func (m *manager) ListDVSwitch() ([]DVSwitch, error) {
	var dvSwitchs []DVSwitch
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().Get(strings.Replace(dvSwitchUrl, siteMask, m.siteUri, -1))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		var listDVSwitchResponse ListDVSwitchResponse
		err := json.Unmarshal(resp.Body(), &listDVSwitchResponse)
		if err != nil {
			return nil, err
		}
		dvSwitchs = listDVSwitchResponse.DVSwitchs
	} else {
		return nil, common.FormatHttpError(resp)
	}
	return dvSwitchs, nil
}

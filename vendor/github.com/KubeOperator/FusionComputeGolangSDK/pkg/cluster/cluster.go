package cluster

import (
	"encoding/json"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
	"strings"
)

const (
	siteMask   = "<site_uri>"
	clusterUrl = "<site_uri>/clusters"
)

type Manager interface {
	ListCluster() ([]Cluster, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) ListCluster() ([]Cluster, error) {
	var clusters []Cluster
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().Get(strings.Replace(clusterUrl, siteMask, m.siteUri, -1))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		var listClusterResponse ListClusterResponse
		err := json.Unmarshal(resp.Body(), &listClusterResponse)
		if err != nil {
			return nil, err
		}
		clusters = listClusterResponse.Clusters
	} else {
		return nil, common.FormatHttpError(resp)
	}
	return clusters, nil

}

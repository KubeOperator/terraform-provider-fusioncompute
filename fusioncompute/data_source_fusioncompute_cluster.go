package fusioncompute

import (
	"errors"
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/cluster"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFusionComputeCluster() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the cluster.",
			},
			"site_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The urn of site.",
			},
		},
		Read: dataSourceFusionComputeClusterRead,
	}
}

func dataSourceFusionComputeClusterRead(d *schema.ResourceData, meta interface{}) error {
	clusterName := d.Get("name").(string)
	c := meta.(client.FusionComputeClient)
	siteUri := d.Get("site_uri").(string)
	m := cluster.NewManager(c, siteUri)
	clusters, err := m.ListCluster()
	if err != nil {
		return errors.New(fmt.Sprintf("error fetching site:%s", err))
	}
	var result *cluster.Cluster
	for i := range clusters {
		if clusters[i].Name == clusterName {
			result = &clusters[i]
		}
	}
	if result == nil {
		return errors.New(fmt.Sprintf("site: %s not exists", clusterName))
	}
	d.SetId(result.Urn)
	return nil
}

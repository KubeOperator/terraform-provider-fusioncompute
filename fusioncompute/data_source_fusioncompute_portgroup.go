package fusioncompute

import (
	"errors"
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/network"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFusionComputePortGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the portgroup.",
			},
			"site_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The uri of site.",
			},
			"dvswitch_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The uri of dvswitch.",
			},
		},
		Read: dataSourceFusionComputePortGroupRead,
	}
}

func dataSourceFusionComputePortGroupRead(d *schema.ResourceData, meta interface{}) error {
	portGroupName := d.Get("name").(string)
	c := meta.(client.FusionComputeClient)
	siteUri := d.Get("site_uri").(string)
	dvSwitchUri := d.Get("dvswitch_uri").(string)
	m := network.NewManager(c, siteUri)
	portGroups, err := m.ListPortGroup(dvSwitchUri)
	if err != nil {
		return errors.New(fmt.Sprintf("error fetching vm:%s", err))
	}
	var result *network.PortGroup
	for _, v := range portGroups {
		if v.Name == portGroupName {
			result = &v
		}
	}
	if result == nil {
		return errors.New(fmt.Sprintf("site: %s not exists", portGroupName))
	}
	d.SetId(result.Uri)
	return nil
}

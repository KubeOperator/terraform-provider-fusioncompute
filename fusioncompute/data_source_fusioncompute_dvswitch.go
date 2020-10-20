package fusioncompute

import (
	"errors"
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/network"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFusionComputeDVSwitch() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the vm.",
			},
			"site_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The urn of site.",
			},
		},
		Read: dataSourceFusionComputeDVSwitchRead,
	}
}

func dataSourceFusionComputeDVSwitchRead(d *schema.ResourceData, meta interface{}) error {
	vSwitchName := d.Get("name").(string)
	c := meta.(client.FusionComputeClient)
	siteUri := d.Get("site_uri").(string)
	m := network.NewManager(c, siteUri)
	vSwitchs, err := m.ListDVSwitch()
	if err != nil {
		return errors.New(fmt.Sprintf("error fetching vm:%s", err))
	}
	var result *network.DVSwitch
	for i := range vSwitchs {
		if vSwitchs[i].Name == vSwitchName {
			result = &vSwitchs[i]
		}
	}
	if result == nil {
		return errors.New(fmt.Sprintf("site: %s not exists", vSwitchName))
	}
	d.SetId(result.Uri)
	return nil
}

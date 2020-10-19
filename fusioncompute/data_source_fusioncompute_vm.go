package fusioncompute

import (
	"errors"
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/vm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFusionComputeVm() *schema.Resource {
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
		Read: dataSourceFusionComputeVmRead,
	}
}

func dataSourceFusionComputeVmRead(d *schema.ResourceData, meta interface{}) error {
	vmName := d.Get("name").(string)
	c := meta.(client.FusionComputeClient)
	siteUri := d.Get("site_uri").(string)
	m := vm.NewManager(c, siteUri)
	vms, err := m.ListVm(true)
	if err != nil {
		return errors.New(fmt.Sprintf("error fetching vm:%s", err))
	}
	var result *vm.Vm
	for _, v := range vms {
		if v.Name == vmName {
			result = &v
		}
	}
	if result == nil {
		return errors.New(fmt.Sprintf("site: %s not exists", vmName))
	}
	d.SetId(result.Uri)
	return nil
}

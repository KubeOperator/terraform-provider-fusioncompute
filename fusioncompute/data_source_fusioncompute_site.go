package fusioncompute

import (
	"errors"
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/site"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFusionComputeSite() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the site.",
			},
		},
		Read: dataSourceFusionComputeSiteRead,
	}
}

func dataSourceFusionComputeSiteRead(d *schema.ResourceData, meta interface{}) error {
	siteName := d.Get("name").(string)
	c := meta.(client.FusionComputeClient)
	m := site.NewManager(c)
	sites, err := m.ListSite()
	if err != nil {
		return errors.New(fmt.Sprintf("error fetching site:%s", err))
	}
	var result *site.Site
	for i := range sites {
		if sites[i].Name == siteName {
			result = &sites[i]
		}
	}
	if result == nil {
		return errors.New(fmt.Sprintf("site: %s not exists", siteName))
	}
	d.SetId(result.Uri)
	return nil
}

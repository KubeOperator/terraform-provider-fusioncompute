package fusioncompute

import (
	"errors"
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/storage"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFusionComputeDatastore() *schema.Resource {
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
		Read: dataSourceFusionComputeDatastoreRead,
	}
}

func dataSourceFusionComputeDatastoreRead(d *schema.ResourceData, meta interface{}) error {
	datastoreName := d.Get("name").(string)
	c := meta.(client.FusionComputeClient)
	siteUri := d.Get("site_uri").(string)
	m := storage.NewManager(c, siteUri)
	datastores, err := m.ListDataStore()
	if err != nil {
		return errors.New(fmt.Sprintf("error fetching site:%s", err))
	}
	var result *storage.Datastore
	for _, d := range datastores {
		if d.Name == datastoreName {
			result = &d
		}
	}
	if result == nil {
		return errors.New(fmt.Sprintf("site: %s not exists", datastoreName))
	}
	uri := result.Uri
	d.SetId(uri)
	return nil
}

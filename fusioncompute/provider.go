package fusioncompute

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FUSION_COMPUTE_USER", nil),
				Description: "The user name for FusionCompute API operations.",
			},

			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FUSION_COMPUTE_PASSWORD", nil),
				Description: "The user password for FusionCompute API operations.",
			},

			"server": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FUSION_COMPUTE_SERVER", nil),
				Description: "The  Server  for FusionCompute API operations.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"fusioncompute_site":      dataSourceFusionComputeSite(),
			"fusioncompute_cluster":   dataSourceFusionComputeCluster(),
			"fusioncompute_datastore": dataSourceFusionComputeDatastore(),
		},
	}
}

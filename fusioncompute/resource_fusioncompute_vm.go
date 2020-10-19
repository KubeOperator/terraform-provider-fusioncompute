package fusioncompute

import (
	"errors"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/task"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/vm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
	"time"
)

func resourceFusionComputeVm() *schema.Resource {
	s := map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"timeout": {
			Type:     schema.TypeInt,
			Required: false,
		},
		"num_cpus": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"memory": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"site_uri": {
			Type:     schema.TypeString,
			Required: true,
		},
		"cluster_uri": {
			Type:     schema.TypeString,
			Required: true,
		},
		"datastore_uri": {
			Type:     schema.TypeString,
			Required: true,
		},
		"template_uri": {
			Type:     schema.TypeString,
			Required: true,
		},
		"network_interface": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 10,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"portgroup_uri": {
					Type:     schema.TypeString,
					Optional: true,
				},
			}},
		},
		"disk": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"size": {
					Type:     schema.TypeInt,
					Optional: true,
				},
				"thin": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			}},
		},
		"customize": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"host_name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"set_dns": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"add_dns": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"network_interface": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"ipv4_address": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv4_netmask": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
					},
				},
			},
			},
		},
	}
	return &schema.Resource{
		Schema:       s,
		Read:         resourceFusionComputeVmRead,
		Update:       nil,
		Create:       resourceFusionComputeCreate,
		Delete:       nil,
		Importer:     nil,
		MigrateState: nil,
	}
}

func resourceFusionComputeCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(client.FusionComputeClient)
	siteUri := d.Get("site_uri").(string)
	m := vm.NewManager(c, siteUri)
	templateUri := d.Get("template_uri").(string)
	req := vm.CloneVmRequest{
		Name:          d.Get("name").(string),
		Location:      fusioncomputeUriToUrn(d.Get("cluster_uri").(string)),
		IsBindingHost: false,
		Config: vm.Config{
			Cpu: vm.Cpu{
				Quantity:    d.Get("num_cpus").(int),
				Reservation: d.Get("num_cpus").(int),
			},
			Memory: vm.Memory{
				QuantityMB:  d.Get("memory").(int),
				Reservation: d.Get("memory").(int),
			},
			Disks: nil,
			Nics:  nil,
		},
		VmCustomization: vm.Customization{},
	}
	for i, di := range d.Get("disk").([]map[string]interface{}) {
		req.Config.Disks = append(req.Config.Disks, vm.Disk{
			SequenceNum:  i,
			QuantityGB:   di["size"].(int),
			IsDataCopy:   false,
			DatastoreUrn: fusioncomputeUriToUrn(d.Get("datastore_uri").(string)),
			IsThin:       true,
		})
	}
	for _, ni := range d.Get("network_interface").([]map[string]interface{}) {
		req.Config.Nics = append(req.Config.Nics, vm.Nic{
			PortGroupUrn: fusioncomputeUriToUrn(ni["portgroup_uri"].(string)),
		})
	}
	ccs := d.Get("customize").([]map[string]interface{})
	if len(ccs) > 0 {
		hostname := ccs[0]["host_name"].(string)
		req.VmCustomization = vm.Customization{
			OsType:   "Linux",
			Hostname: hostname,
		}
		networkCcs := ccs[0]["network_interface"].([]map[string]interface{})
		for i, nccs := range networkCcs {
			req.VmCustomization.NicSpecification = append(req.VmCustomization.NicSpecification, vm.NicSpecification{
				SequenceNum: i,
				Ip:          nccs["ipv4_address"].(string),
				Netmask:     nccs["ipv4_netmask"].(string),
				Gateway:     ccs[0]["ipv4_gateway"].(string),
				Setdns:      ccs[0]["set_dns"].(string),
				Adddns:      ccs[0]["add_dns"].(string),
			})
		}
	}
	resp, err := m.CloneVm(templateUri, req)
	if err != nil {
		return err
	}
	timeout := d.Get("timeout").(int)

	tm := task.NewManager(c, siteUri)
	n := 0
	for {
		r, err := tm.Get(resp.TaskUri)
		if err != nil {
			return err
		}
		if r.Progress == 100 {
			break
		}
		if 5*n >= timeout*60 {
			return errors.New("create vm timeout")
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}

func resourceFusionComputeVmRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(client.FusionComputeClient)
	err := c.Connect()
	if err != nil {
		return err
	}
	return nil
}

func fusioncomputeUriToUrn(itemUri string) string {
	str := strings.Replace(itemUri, "/service", "urn", -1)
	return strings.Replace(str, "/", ":", -1)
}

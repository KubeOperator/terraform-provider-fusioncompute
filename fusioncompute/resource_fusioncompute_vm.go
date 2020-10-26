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
			Optional: true,
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
				"ipv4_gateway": {
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
		Schema: s,
		Read:   resourceFusionComputeVmRead,
		Update: resourceFusionComputeUpdate,
		Create: resourceFusionComputeCreate,
		Delete: resourceFusionComputeDelete,
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
	for i, di := range d.Get("disk").([]interface{}) {
		di := di.(map[string]interface{})
		req.Config.Disks = append(req.Config.Disks, vm.Disk{
			SequenceNum:  i + 1,
			QuantityGB:   di["size"].(int),
			IsDataCopy:   false,
			DatastoreUrn: fusioncomputeUriToUrn(d.Get("datastore_uri").(string)),
			IsThin:       true,
		})
	}
	for _, ni := range d.Get("network_interface").([]interface{}) {
		ni := ni.(map[string]interface{})
		req.Config.Nics = append(req.Config.Nics, vm.Nic{
			PortGroupUrn: fusioncomputeUriToUrn(ni["portgroup_uri"].(string)),
		})
	}
	ccs := d.Get("customize").([]interface{})
	firstCcs := ccs[0].(map[string]interface{})
	if len(ccs) > 0 {
		hostname := firstCcs["host_name"].(string)
		req.VmCustomization = vm.Customization{
			OsType:   "Linux",
			Hostname: hostname,
		}
		networkCcs := firstCcs["network_interface"].([]interface{})
		for i, nccs := range networkCcs {
			nccs := nccs.(map[string]interface{})
			req.VmCustomization.NicSpecification = append(req.VmCustomization.NicSpecification, vm.NicSpecification{
				SequenceNum: i + 1,
				Ip:          nccs["ipv4_address"].(string),
				Netmask:     nccs["ipv4_netmask"].(string),
				Gateway:     firstCcs["ipv4_gateway"].(string),
				Setdns:      firstCcs["set_dns"].(string),
				Adddns:      firstCcs["add_dns"].(string),
			})
		}
	}
	resp, err := m.CloneVm(templateUri, req)
	if err != nil {
		return err
	}
	d.SetId(resp.Uri)
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
		n++
		time.Sleep(5 * time.Second)
	}
	return resourceFusionComputeVmRead(d, meta)
}

func resourceFusionComputeVmRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(client.FusionComputeClient)
	err := c.Connect()
	if err != nil {
		return err
	}
	id := d.Id()
	siteUri := d.Get("site_uri").(string)
	vmm := vm.NewManager(c, siteUri)
	_, err = vmm.GetVM(id)
	if err != nil {
		return err
	}
	return nil
}

func resourceFusionComputeDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(client.FusionComputeClient)
	err := c.Connect()
	if err != nil {
		return err
	}
	id := d.Id()
	siteUri := d.Get("site_uri").(string)
	vmm := vm.NewManager(c, siteUri)
	resp, err := vmm.DeleteVm(id)
	if err != nil {
		return err
	}
	tm := task.NewManager(c, siteUri)
	timeout := d.Get("timeout").(int)
	// wait for delete
	n := 0
	for {
		t, err := tm.Get(resp.TaskUri)
		if err != nil {
			return err
		}
		if t.Progress == 100 {
			break
		}
		if 5*n >= timeout*60 {
			return errors.New("delete vm timeout")
		}
		n++
		time.Sleep(5 * time.Second)
	}
	return nil
}

func resourceFusionComputeUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func fusioncomputeUriToUrn(itemUri string) string {
	str := strings.Replace(itemUri, "/service", "urn", -1)
	return strings.Replace(str, "/", ":", -1)
}

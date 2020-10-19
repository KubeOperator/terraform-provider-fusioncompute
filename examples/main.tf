provider "fusioncompute" {
  user = "kubeoperator"
  password = "Calong@2015"
  server = "https://100.199.16.208:8443"
}

data "fusioncompute_site" "site" {
  name = "site"
}

data "fusioncompute_cluster" "cluster" {
  name = ""
  site_uri = "data.fusioncompute_site.site.id"
}

data "fusioncompute_dvswitch" "dvswitch" {
  name = ""
  site_uri = "data.fusioncompute_site.site.id"
}

data "fusioncompute_portgroup" "portgroup" {
  name = ""
  dvswitch_uri = "data.fusioncompute_dvswitch.dvswitch.id"
  site_uri = "data.fusioncompute_site.site.id"
}

data "fusioncompute_datastore" "datastore" {
  name = ""
  site_uri = "data.fusioncompute_site.site.id"
}

data "fusioncompute_vm" "template" {
  name = ""
  site_uri = "data.fusioncompute_site.site.id"
}

resource "fusioncompute_vm" "test" {
  name = ""

  timeout = 30


  num_cpus = 2
  memory = 1024

  site_uri = "data.fusioncompute_site.site.id"
  cluster_uri = "data.fusioncompute_cluster.cluster.id"
  datastore_uri = "data.fusioncompute_datastore.datastore.id"


  template_uri = "data.fusioncompute_vm.template.id"

  network_interface {
    portgroup_uri = "data.fusioncompute_portgroup.portgroup.id"
  }
  disk {
    size = 50
    thin = true
  }
  customize {
    host_name = ""
    network_interface {
      ipv4_address = ""
      ipv4_netmask = ""
    }
    ipv4_gateway = ""
    set_dns = ""
    add_dns = ""
  }
}


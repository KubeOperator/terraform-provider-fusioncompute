provider "fusioncompute" {
  user = "kubeoperator"
  password = "Calong@2015"
  server = "https://100.199.16.208:7443"
}

data "fusioncompute_site" "site" {
  name = "site"
}

data "fusioncompute_cluster" "cluster" {
  name = "ManagementCluster"
  site_uri = data.fusioncompute_site.site.id
}
data "fusioncompute_dvswitch" "dvswitch" {
  name = "ManagementDVS"
  site_uri = data.fusioncompute_site.site.id
}
data "fusioncompute_portgroup" "portgroup" {
  name = "managePortgroup"
  dvswitch_uri = data.fusioncompute_dvswitch.dvswitch.id
  site_uri = data.fusioncompute_site.site.id
}

data "fusioncompute_datastore" "datastore" {
  name = "Storage"
  site_uri = data.fusioncompute_site.site.id
}

data "fusioncompute_vm" "template" {
  name = "CentOS7.6-template"
  site_uri = data.fusioncompute_site.site.id
}


resource "fusioncompute_vm" "test" {
  name = "aaabbb"
  timeout = 30
  num_cpus = 2
  memory = 1024
  site_uri = data.fusioncompute_site.site.id
  cluster_uri = data.fusioncompute_cluster.cluster.id
  datastore_uri = data.fusioncompute_datastore.datastore.id
  template_uri = data.fusioncompute_vm.template.id
  network_interface {
    portgroup_uri = data.fusioncompute_portgroup.portgroup.id
  }
  disk {
    size = 50
    thin = true
  }
  customize {
    host_name = "aaaa"
    network_interface {
      ipv4_address = "100.199.10.88"
      ipv4_netmask = "255.255.255.0"
    }
    ipv4_gateway = "100.199.10.0"
    set_dns = "8.8.8.8"
    add_dns = "114.114.114.114"
  }
}


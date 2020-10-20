package vm

type Vm struct {
	Urn               string `json:"urn,omitempty,omitempty"`
	Uri               string `json:"uri,omitempty"`
	Uuid              string `json:"uuid,omitempty"`
	Name              string `json:"name,omitempty"`
	Arch              string `json:"arch,omitempty"`
	Description       string `json:"description,omitempty"`
	Group             string `json:"group,omitempty"`
	Location          string `json:"location,omitempty"`
	LocationName      string `json:"locationName,omitempty"`
	HostUrn           string `json:"hostUrn,omitempty"`
	Status            string `json:"status,omitempty"`
	PvDriverStatus    string `json:"pvDriverStatus,omitempty"`
	ToolInstallStatus string `json:"toolInstallStatus,omitempty"`
	CdRomStatus       string `json:"cdRomStatus,omitempty"`
	IsTemplate        bool   `json:"isTemplate,omitempty"`
	IsLinkClone       bool   `json:"isLinkClone,omitempty"`
	IsBindingHost     bool   `json:"isBindingHost,omitempty"`
	CreateTime        string `json:"createTime,omitempty"`
	ToolsVersion      string `json:"toolsVersion,omitempty"`
	HostName          string `json:"hostName,omitempty"`
	ClusterName       string `json:"clusterName,omitempty"`
	HugePage          string `json:"hugePage,omitempty"`
	Idle              int    `json:"idle,omitempty"`
	VmType            int    `json:"vmType,omitempty"`
	DrStatus          int    `json:"drStatus,omitempty"`
	RpoStatus         int    `json:"rpoStatus,omitempty"`
	InitSyncStatus    int    `json:"initSyncStatus,omitempty"`
}

type Customization struct {
	OsType             string             `json:"osType,omitempty"`
	Hostname           string             `json:"hostname,omitempty"`
	IsUpdateVmPassword bool               `json:"isUpdateVmPassword,omitempty"`
	Password           string             `json:"password,omitempty"`
	NicSpecification   []NicSpecification `json:"nicSpecification,omitempty"`
}

type NicSpecification struct {
	SequenceNum int    `json:"sequenceNum,omitempty"`
	Ip          string `json:"ip,omitempty"`
	Netmask     string `json:"netmask,omitempty"`
	Gateway     string `json:"gateway,omitempty"`
	Setdns      string `json:"setdns,omitempty"`
	Adddns      string `json:"adddns,omitempty"`
}

type ListVmResponse struct {
	Total int  `json:"total,omitempty"`
	Vms   []Vm `json:"vms,omitempty"`
}

type CloneVmRequest struct {
	Name            string        `json:"name,omitempty"`
	Description     string        `json:"description,omitempty"`
	Group           string        `json:"group,omitempty"`
	Location        string        `json:"location,omitempty"`
	IsBindingHost   bool          `json:"isBindingHost,omitempty"`
	Config          Config        `json:"vmConfig,omitempty"`
	VmCustomization Customization `json:"vmCustomization,omitempty"`
}

type Config struct {
	Cpu    Cpu    `json:"cpu,omitempty"`
	Memory Memory `json:"memory,omitempty"`
	Disks  []Disk `json:"disks,omitempty"`
	Nics   []Nic  `json:"nics,omitempty"`
}
type Cpu struct {
	Quantity    int `json:"quantity,omitempty"`
	Reservation int `json:"reservation,omitempty"`
	Weight      int `json:"weight,omitempty"`
	Limit       int `json:"limit,omitempty"`
}

type Memory struct {
	QuantityMB  int `json:"quantityMb,omitempty"`
	Reservation int `json:"reservation,omitempty"`
	Weight      int `json:"weight,omitempty"`
	Limit       int `json:"limit,omitempty"`
}

type Disk struct {
	SequenceNum  int    `json:"sequenceNum,omitempty"`
	QuantityGB   int    `json:"quantityGb,omitempty"`
	IsDataCopy   bool   `json:"isDataCopy,omitempty"`
	DatastoreUrn string `json:"datastoreUrn,omitempty"`
	IsThin       bool   `json:"isThin,omitempty"`
}

type Nic struct {
	Name         string `json:"name,omitempty"`
	PortGroupUrn string `json:"portGroupUrn,omitempty"`
	Mac          string `json:"mac,omitempty"`
}

type CloneVmResponse struct {
	Urn     string `json:"urn,omitempty"`
	Uri     string `json:"uri,omitempty"`
	TaskUrn string `json:"taskUrn,omitempty"`
	TaskUri string `json:"taskUri,omitempty"`
}

package network

type DVSwitch struct {
	Name           string `json:"name"`
	Uri            string `json:"uri"`
	Urn            string `json:"urn"`
	Description    string `json:"description"`
	IsIgmpSnooping bool   `json:"isIgmpSnooping"`
	Type           int    `json:"type"`
	PortGroupNum   int    `json:"portGroupNum"`
	QosType        int    `json:"qosType"`
	Mtu            int    `json:"mtu"`
}

type ListDVSwitchResponse struct {
	DVSwitchs []DVSwitch `json:"dvSwitchs"`
}

type PortGroup struct {
	Urn               string `json:"urn"`
	Uri               string `json:"uri"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	IsDhcpIsolation   bool   `json:"isDhcpIsolation"`
	VlanId            int    `json:"vlanId"`
	TxLimit           int    `json:"txLimit"`
	Priority          int    `json:"priority"`
	IsIpMacBind       bool   `json:"isIpMacBind"`
	PortType          int    `json:"portType"`
	TxPeakLimit       int    `json:"txPeakLimit"`
	TxBurstSize       int    `json:"txBurstSize"`
	RxLimit           int    `json:"rxLimit"`
	RxPeakLimit       int    `json:"rxPeakLimit"`
	RxBurstSize       int    `json:"rxBurstSize"`
	TxWeight          int    `json:"txWeight"`
	IpBcstSuppress    int    `json:"ipBcstSuppress"`
	IsCalcTCPCheckSum bool   `json:"isCalcTcpCheckSum"`
	IsQinQEnable      bool   `json:"isQinQEnable"`
}

type ListPortGroupResponse struct {
	PortGroups []PortGroup `json:"portGroups"`
}

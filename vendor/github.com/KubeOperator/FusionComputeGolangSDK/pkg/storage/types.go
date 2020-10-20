package storage

type Datastore struct {
	Urn              string   `json:"urn"`
	Uri              string   `json:"uri"`
	StorageType      string   `json:"storageType"`
	Name             string   `json:"name"`
	Status           string   `json:"status"`
	ClusterSize      int      `json:"clusterSize"`
	CapacityGB       int      `json:"capacityGB"`
	UsedSizeGB       int      `json:"usedSizeGB"`
	FreeSizeGB       int      `json:"freeSizeGB"`
	ActualCapacityGB int      `json:"actualCapacityGB"`
	ActualFreeSizeGB int      `json:"actualFreeSizeGB"`
	DsLockType       int      `json:"dsLockType"`
	RefreshTime      string   `json:"refreshTime"`
	Hosts            []string `json:"hosts"`
}

type ListDataStoreResponse struct {
	Datastores []Datastore `json:"datastores"`
}

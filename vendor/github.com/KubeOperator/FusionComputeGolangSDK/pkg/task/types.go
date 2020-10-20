package task

type Task struct {
	Urn        string `json:"urn"`
	Uri        string `json:"uri"`
	Type       string `json:"type"`
	EntityUrn  string `json:"entityUrn"`
	EntityName string `json:"entityName"`
	StartTime  string `json:"startTime"`
	FinishTime string `json:"finishTime"`
	User       string `json:"user"`
	Status     string `json:"status"`
	Progress   int    `json:"progress"`
	Reason     string `json:"reason"`
	ReasonDes  string `json:"reasonDes"`
}

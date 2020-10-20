package site

type Site struct {
	Uri         string `json:"uri"`
	Urn         string `json:"urn"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ListSiteResponse struct {
	Sites []Site `json:"sites"`
}

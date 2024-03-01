package accountlist

// AccountTag represents the data we collect with this app.
type AccountTag struct {
	ID     string      `json:"id"`
	ARN    string      `json:"arn,omitempty"`
	Name   string      `json:"name,omitempty"`
	Email  string      `json:"email,omitempty"`
	Status string      `json:"status,omitempty"`
	Tags   []TagValues `json:"tags,omitempty"`
	OUs    []OUType    `json:"organizationalUnits,omitempty"`
}

// OUType represents the data we collect with this app.
type OUType struct {
	accountID string
	ID        string `json:"id,omitempty"`
	ARN       string `json:"arn,omitempty"`
	Name      string `json:"name,omitempty"`
}

// TagValues represents tag key-value pairs.
type TagValues struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

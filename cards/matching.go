package cards

type Matching struct {
	Id        string                 `json:"id"`
	Type      string                 `json:"type"`
	Category  string                 `json:"category"`
	Questions map[string]interface{} `json:"questions"`
	Options   map[string]interface{} `json:"options"`
	Answers   map[string]interface{} `json:"answers"`
}

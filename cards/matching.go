package cards

type Matching struct {
	Id        string            `json:"id"`
	Type      string            `json:"type"`
	Category  string            `json:"category"`
	Questions map[string]string `json:"questions"`
	Options   map[string]string `json:"options"`
	Answers   map[string]string `json:"answers"`
}

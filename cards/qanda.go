package cards

type QNA struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

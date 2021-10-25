package cards

type QNA struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

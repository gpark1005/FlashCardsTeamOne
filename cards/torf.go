package cards

type TrueOrFalse struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   bool   `json:"answer"`
}

package cards

type MultipleChoice struct {
	Id       string                `json:"id"`
	Type     string                `json:"type"`
	Category string                `json:"category"`
	Question string                `json:"question"`
	Options  MultipleChoiceOptions `json:"options"`
	Answers  int                   `json:"answers"`
}

type MultipleChoiceOptions struct {
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
}

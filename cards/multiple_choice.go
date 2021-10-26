package cards

type MultipleChoice struct {
	Id       string          `json:"id" validate:"omitempty,uuid"`
	Type     string          `json:"type" validate:"required"`
	Category string          `json:"category" validate:"required"`
	Question string          `json:"question" validate:"required"`
	Options  MultipleChoiceO `json:"options" validate:"required"`
	Answers  int             `json:"answers" validate:"required"`
}

type MultipleChoiceO struct {
	One   string `json:"1" validate:"required"`
	Two   string `json:"2" validate:"required"`
	Three string `json:"3" validate:"required"`
	Four  string `json:"4" validate:"required"`
}

package cards

type QNA struct {
	Id       string `json:"id" validate:"omitempty,uuid"`
	Type     string `json:"type" validate:"required"`
	Category string `json:"category" validate:"required"`
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required"`
}

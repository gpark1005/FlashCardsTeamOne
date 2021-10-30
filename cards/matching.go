package cards

type Matching struct {
	Id        string                 `json:"id" validate:"omitempty,uuid"`
	Type      string                 `json:"type" validate:"required"`
	Category  string                 `json:"category" validate:"required"`
	Questions map[string]interface{} `json:"questions" validate:"required"`
	Options   map[string]interface{} `json:"options" validate:"required"`
	Answers   map[string]interface{} `json:"answers" validate:"required"`
}

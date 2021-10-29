package cards

type Info struct {
	Id          string `json:"id" validate:"omitempty,uuid"`
	Type        string `json:"type" validate:"required"`
	Category	string `json:"category" validate:"required"`
	Information string `json:"information" validate:"required"`
}

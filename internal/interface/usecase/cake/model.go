package cake

type Sort struct {
	Column string
	Type   string
}

type CakeRequest struct {
	Title       string   `json:"title" validate:"required"`
	Description *string  `json:"description" validate:"omitempty"`
	Rating      *float32 `json:"rating" validate:"omitempty,numeric,gte=0,lte=10"`
	Image       *string  `json:"image" validate:"omitempty,url"`
}

package cake

type Sort struct {
	Column string
	Type   string
}

type CakeRequest struct {
	Title       string   `json:"title" validate:"required,max=100"`
	Description *string  `json:"description" validate:"omitempty"`
	Rating      *float32 `json:"rating" validate:"omitempty,lte=10,gte=0,numeric"`
	Image       *string  `json:"image" validate:"omitempty,max=200"`
}

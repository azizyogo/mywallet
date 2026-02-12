package request

type TopUpRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

package request

type TransferRequest struct {
	ReceiverEmail string  `json:"receiver_email" binding:"required,email"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	Description   string  `json:"description"`
}

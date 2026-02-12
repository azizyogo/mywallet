package constant

type TransactionType string
type TransactionStatus string

const (
	TransactionTypeTopUp    TransactionType = "TOPUP"
	TransactionTypeTransfer TransactionType = "TRANSFER"
)

const (
	TransactionStatusPending TransactionStatus = "PENDING"
	TransactionStatusSuccess TransactionStatus = "SUCCESS"
	TransactionStatusFailed  TransactionStatus = "FAILED"
)

package request

import "wetodo/internal/transaction"

type SaveTransactionRequest struct {
	Transactions []transaction.Transaction
}

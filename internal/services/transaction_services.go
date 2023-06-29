package services

import (
	"github.com/KhoirulAziz99/mnc/internal/models"
	"github.com/KhoirulAziz99/mnc/internal/repository"
)

type TransactionServices interface {
	Make(transaction models.Transaction) error
}

type transactionServices struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionServices (transaction repository.TransactionRepository) *transactionServices {
	return &transactionServices{transactionRepository: transaction}
}

func (t transactionServices) Make(transaction models.Transaction) error {
	return t.transactionRepository.Create(transaction)
}
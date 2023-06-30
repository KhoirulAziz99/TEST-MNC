package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KhoirulAziz99/mnc/internal/models"
	"github.com/KhoirulAziz99/mnc/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewTransactionRepository(db)

	transaction := models.Transaction{
		CustomerEmail: "johndoe@example.com",
		MerchantName:  "examplemerchant",
		Paid:          100,
	}

	// Expect query to fetch customer balance
	mock.ExpectQuery("SELECT balance FROM customer").
		WithArgs(transaction.CustomerEmail).
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(200))

	// Expect query to fetch merchant balance
	mock.ExpectQuery("SELECT balance FROM merchant").
		WithArgs(transaction.MerchantName).
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(300))

	// Expect query to update customer balance
	mock.ExpectExec("UPDATE customer SET balance").
		WithArgs(transaction.Paid, transaction.CustomerEmail).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Expect query to update merchant balance
	mock.ExpectExec("UPDATE merchant SET balance").
		WithArgs(transaction.Paid, transaction.MerchantName).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Expect query to insert transaction history
	mock.ExpectExec("INSERT INTO history_trx").
		WithArgs(transaction.CustomerEmail, transaction.MerchantName, sqlmock.AnyArg(), transaction.Paid).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(transaction)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestTransactionRepository_History(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewTransactionRepository(db)

	email := "johndoe@example.com"

	// Expect query to fetch transaction history
	mock.ExpectQuery("SELECT c.id, c.name, c.email").
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "John Doe", email))

	histories, err := repo.History(email)
	assert.NoError(t, err)
	assert.Len(t, histories, 1)

	assert.NoError(t, mock.ExpectationsWereMet())
}

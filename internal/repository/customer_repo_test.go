package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/KhoirulAziz99/mnc/internal/models"
	"github.com/KhoirulAziz99/mnc/internal/repository"
)

func TestCustomerRepository_Created(t *testing.T) {
	
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer db.Close()

	
	customerRepo := repository.NewCustomerRepository(db)

	
	newCustomer := models.Customer{
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		Password:  "password123",
		Balance:   1000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	}

	
	mock.ExpectExec("INSERT INTO customer").
		WithArgs(newCustomer.Name, newCustomer.Email, newCustomer.Password, newCustomer.Balance, sqlmock.AnyArg(), sqlmock.AnyArg(), newCustomer.IsDeleted).
		WillReturnResult(sqlmock.NewResult(1, 1))

	
	err = customerRepo.Created(newCustomer)
	if err != nil {
		t.Errorf("Error calling Created: %v", err)
	}

	
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}

func TestCustomerRepository_FindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer db.Close()

	customerRepo := repository.NewCustomerRepository(db)

	expectedCustomers := []models.Customer{
		{
			ID:        1,
			Name:      "John Doe",
			Email:     "johndoe@example.com",
			Password:  "password123",
			Balance:   1000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsDeleted: false,
		},
		{
			ID:        2,
			Name:      "Jane Smith",
			Email:     "janesmith@example.com",
			Password:  "password456",
			Balance:   2000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsDeleted: false,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "balance", "created_at", "updated_at", "is_deleted"}).
		AddRow(expectedCustomers[0].ID, expectedCustomers[0].Name, expectedCustomers[0].Email, expectedCustomers[0].Password, expectedCustomers[0].Balance, expectedCustomers[0].CreatedAt, expectedCustomers[0].UpdatedAt, expectedCustomers[0].IsDeleted).
		AddRow(expectedCustomers[1].ID, expectedCustomers[1].Name, expectedCustomers[1].Email, expectedCustomers[1].Password, expectedCustomers[1].Balance, expectedCustomers[1].CreatedAt, expectedCustomers[1].UpdatedAt, expectedCustomers[1].IsDeleted)

	mock.ExpectQuery("SELECT id, name, email, password, balance, created_at, updated_at, is_deleted FROM customer").
		WillReturnRows(rows)

	customers, err := customerRepo.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomers, customers)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCustomerRepository_FindByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer db.Close()

	customerRepo := repository.NewCustomerRepository(db)

	email := "khoirulaziz@example.com"
	expectedCustomer := &models.Customer{
		ID:        1,
		Name:      "Khoirul Aziz",
		Email:     email,
		Password:  "aziz123",
		Balance:   1000,
		IsDeleted: false,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "balance", "is_deleted"}).
		AddRow(expectedCustomer.ID, expectedCustomer.Name, expectedCustomer.Email, expectedCustomer.Password, expectedCustomer.Balance, expectedCustomer.IsDeleted)

	mock.ExpectQuery("SELECT id, name, email, password, balance, is_deleted FROM customer WHERE email = ?").
		WithArgs(email).
		WillReturnRows(rows)

	customer, err := customerRepo.FindByEmail(email)
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customer)

	assert.NoError(t, mock.ExpectationsWereMet())
}

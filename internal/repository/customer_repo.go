package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/KhoirulAziz99/mnc/internal/models"
)



// CustomerRepository adalah sebuah ainterface yang mendefinisikan method-method untuk berinteraksi dengan data customer.
type CustomerRepository interface {
	Created(newCustomer models.Customer) error
	FindAll() ([]models.Customer, error)
	FindByEmail(email string) (*models.Customer, error)
}

// customerRepository adalah sebuah struct yang mengimplementasikan interface CustomerRepository.
type customerRepository struct {
	db *sql.DB
}

// NewCustomerRepository membuat instance baru dari customerRepository.
func NewCustomerRepository(db *sql.DB) *customerRepository {
	return &customerRepository{db: db}
}

// Created menambahkan customer baru ke dalam database.
func (c customerRepository) Created(newCustomer models.Customer) error {
	query := `INSERT INTO customer (name, email, password, balance, created_at, updated_at, is_deleted ) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	timeNow := time.Now()
	_, err := c.db.Exec(query, newCustomer.Name, newCustomer.Email, newCustomer.Password, newCustomer.Balance, timeNow, timeNow, newCustomer.IsDeleted)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Sukses menambahkan pelanggan baru")
	return err
}

// FindAll mengambil semua data customer dari database.
func (c customerRepository) FindAll() ([]models.Customer, error) {
	query := `SELECT id, name, email, password, balance, created_at, updated_at, is_deleted FROM customer`
	rows, err := c.db.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	customers := []models.Customer{}
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Password, &customer.Balance, &customer.CreatedAt, &customer.UpdatedAt, &customer.IsDeleted)
		if err != nil {
			log.Println(err)
		}
		customers = append(customers, customer)
	}
	return customers, err
}

// FindByEmail mencari data customer berdasarkan email dari database.
func (c customerRepository) FindByEmail(email string) (*models.Customer, error) {
	query := `SELECT id, name, email, password, balance, is_deleted FROM customer WHERE email = $1`
	row := c.db.QueryRow(query, email)
	var customer models.Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Password, &customer.Balance, &customer.IsDeleted)
	if err != nil {
		log.Println(err)
	}
	customer.Email = email

	return &customer, err
}

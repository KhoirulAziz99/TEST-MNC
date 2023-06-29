package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/KhoirulAziz99/mnc/internal/models"
)

type TransactionRepository interface {
	Create(transaction models.Transaction) error
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (t transactionRepository) Create(transaction models.Transaction) error {
	trx, err := t.db.Begin() //Mulai transaksi
	if err != nil {
		fmt.Println(err)
	}

	// Cek apakah customer dengan email yang diberikan ada dalam database
	query1 := `SELECT balance FROM customer WHERE email = $1`
	row := trx.QueryRow(query1, transaction.CustomerEmail)
	var balanceCustomer int64
	err = row.Scan(&balanceCustomer)
	if err != nil {
		_ = trx.Rollback() // Membatalkan transaksi
		if err == sql.ErrNoRows {
			log.Fatalf("Customer not found")
			return err
		}
		return err
	}

	// Cek apakah customer dengan nama yang diberikan ada dalam database
	query2 := `SELECT balance FROM merchant WHERE name = $1`
	row = trx.QueryRow(query2, transaction.MerchantName)
	var balanceMerchant int64
	 err = row.Scan(&balanceMerchant)
	if err != nil {
		_ = trx.Rollback() // Membatalkan transaksi
		if err == sql.ErrNoRows {
			log.Fatalf("Merchent not found")
			return err
		}
		return err
	}

	if balanceCustomer == 0 {
		err = trx.Rollback() // Batalkan transaksi
		fmt.Println("Your balence is 0, please top-up your wallet")
		return err
	}
	if balanceCustomer < transaction.Paid {
		err = trx.Rollback()
		fmt.Println("Your balance is not enough to make a transaction")
		return err
	}

	// var trxHistory models.TransactionHistory
	//Kurangi saldo customer
	updateQuery := `UPDATE customer SET balance = balance - $1 WHERE email = $2`
	_, err = trx.Exec(updateQuery, transaction.Paid, transaction.CustomerEmail)
	if err != nil {
		trx.Rollback()
		fmt.Println("error : ", err)
		log.Fatalf("Failed to update balence from customer")
		
		return err
	}

	updateQuery2 := `UPDATE merchant SET balance = balance + $1 WHERE name = $2`
	_, err = trx.Exec(updateQuery2, transaction.Paid, transaction.MerchantName)
	if err != nil {
		trx.Rollback()
		fmt.Println("error : ", err)
		log.Fatalf("Failed to update balence from merchant")
		return err
	}

	//Simpan data di tabel transaksi
	timeNow := time.Now()
	insertQuery := `INSERT INTO history_trx (customer_email, merchant_name, created_at) VALUES ($1, $2, $3)`
	_, err = trx.Exec(insertQuery, transaction.CustomerEmail, transaction.MerchantName, timeNow)
	if err != nil {
		trx.Rollback()
		log.Println("error : ", err)
		log.Fatalf("Transaction Failed")
		return err
	}
	
	err = trx.Commit() // Konfirmasi transaksi
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	fmt.Println("Succsess make a transaction")

	return err

}

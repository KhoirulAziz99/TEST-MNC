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
	History(email string) ([]models.TransactionHistory, error)
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
	insertQuery := `INSERT INTO history_trx (customer_email, merchant_name, created_at, amount) VALUES ($1, $2, $3, $4)`
	_, err = trx.Exec(insertQuery, transaction.CustomerEmail, transaction.MerchantName, timeNow, transaction.Paid)
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

func(t transactionRepository) History(email string) ([]models.TransactionHistory, error) {
	query := `SELECT 
	c.id, 
	c.name, 
	c.email, 
	c.password, 
	c.balance, 
	m.id, 
	m.name, 
	m.no_telephon, 
	m.category, 
	m.balance, 
	t.created_at, 
	t.amount 
	FROM history_trx t JOIN  customer c 
	ON t.customer_email=c.email JOIN merchant m 
	ON t.merchant_name = m.name
	WHERE t.customer_email = $1;`

	
	rows, err := t.db.Query(query, email)
	if err != nil {
		log.Println("Failed to get history, err : ",err)
	}
	defer rows.Close()

	histories := []models.TransactionHistory{}
	for rows.Next() {
		var transaction models.TransactionHistory
		err := rows.Scan(
			&transaction.CustomerId.ID,
			&transaction.CustomerId.Name,
			&transaction.CustomerId.Email,
			&transaction.CustomerId.Password,
			&transaction.CustomerId.Balance,
			&transaction.MerchantId.ID,
			&transaction.MerchantId.Name,
			&transaction.MerchantId.NoTelephon,
			&transaction.MerchantId.Category,
			&transaction.MerchantId.Balance,
			&transaction.CreatedAt,
			&transaction.Amount,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transfer row: %v", err)
		}

		if err != nil {
			log.Println(err)
		}
		histories = append(histories, transaction)

		
		
	}

	return histories, err
}
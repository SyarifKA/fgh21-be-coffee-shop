package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func CreateTransaction(data models.Transaction) (models.Transaction, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	fmt.Println(data)
	sql := `
		INSERT INTO transactions (no_order, add_full_name, add_email,
		add_address, payment, user_id, transaction_detail_id,
		order_type_id, transaction_status_id) VALUES 
		($1, $2, $3, $4, $5, $6,$7, $8, $9) RETURNING *
	`

	row, err := db.Query(context.Background(), sql, data.NoOrder, data.AddFullName, data.AddEmail, data.AddAddress, data.Payment, data.UserId, data.TransactionDetail, data.OrderTypeId, data.TransactionStatusId)

	if err != nil {
		return models.Transaction{}, err
	}

	transaction, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Transaction])

	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, err
}
func EditTransactionStatus(data models.Transaction, id int) (models.Transaction, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE transactions SET "transaction_status_id"=$1 WHERE id=$2 returning *`

	query := db.QueryRow(context.Background(), sql, data.TransactionStatusId ,id)

	var result models.Transaction
	err := query.Scan(
		&result.Id,
		&result.TransactionStatusId,
	)
	
	if err != nil {
		log.Println(err)
		return models.Transaction{}, err
	}

	return result, err
}

package service

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/AdityaByte/mysql/model"
	"github.com/AdityaByte/mysql/utils"
)

func AddUser(user *model.User, db *sql.DB, tableName string) error {
	tx, err := db.Begin() // tx is the transaction when we doing a atomic operation we need transaction.
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Rollbacked the changes done by transaction.")
		}
	}()

	if err != nil {
		return fmt.Errorf("ERROR: Failed to begin the transaction: %v", err)
	}

	encodedPassword, err := utils.EncodePassword(user.Password)
	if err != nil {
		return err
	}

	if strings.TrimSpace(encodedPassword) == "" {
		return fmt.Errorf("ERROR: Empty password")
	}

	result, err := tx.Exec(fmt.Sprintf("INSERT INTO %s(email, password) VALUES(?, ?)", tableName), user.Email, encodedPassword)
	if err != nil {
		tx.Rollback();
		return fmt.Errorf("ERROR: User insertion failed: %v", err)
	}

	// If everything goes right then don't forget to commit the transaction changes.
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("ERROR: Failed to commit the transaction: %v", err)
	}

	log.Println("Data Inserted Successfully", result)

	return nil
}

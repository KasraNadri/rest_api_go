package models

import (
	"log"
	"rest_api/db"
	"rest_api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		log.Println("preparing the query failed:", err)
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		log.Println("hashing the password failed:", err)
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		log.Println("inserting the user failed: ", err)
		return err
	}

	userId, err := result.LastInsertId()

	if err != nil {
		log.Println("retrieving the last insert ID failed:", err)
		return err
	}

	u.ID = userId
	return nil
}

func (u User) ValidateCredentials() error {
	query := "SELECT email, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedRow User
	err := row.Scan(&retrievedRow.Email, &retrievedRow.Password)

	if err != nil {
		log.Println("retrieving the record with the email failed", err)
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedRow.Password)

	if !passwordIsValid {
		log.Println("validating the password failed: ", err)
		return err
	}

	return nil
}

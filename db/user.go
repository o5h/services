package db

import (
	"database/sql"
	"time"
)

type User struct {
	UserId           uint64
	UserName         string
	Email            string
	PasswordHash     string
	FirstName        string
	LastName         string
	RegistrationDate time.Time
}

func UserInsert(tx *sql.Tx, user *User) int64 {
	var id int64
	must(tx.QueryRow(
		`INSERT INTO users(
		username,
		email,
		password_hash,
		first_name,
		last_name,
		registration_date
		) VALUES($1, $2, $3, $4, $5, $6)
		RETURNING user_id`,
		user.UserName,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.RegistrationDate,
	).Scan(&id))
	return id
}

func UserGetByUserName(tx *sql.Tx, userName string) *User {
	var user User
	err := tx.QueryRow(
		`SELECT 		
		username,
		email,
		password_hash,
		first_name,
		last_name,
		registration_date
		FROM users 
		WHERE username = $1`, userName).
		Scan(
			&user.UserName,
			&user.Email,
			&user.PasswordHash,
			&user.FirstName,
			&user.LastName,
			&user.RegistrationDate)
	return mustRow(&user, err)
}

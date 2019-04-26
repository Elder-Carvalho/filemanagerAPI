package repository

import (
	"database/sql"
	. "filemanagerAPI/models"
	"golang.org/x/crypto/bcrypt"
	// "fmt"
)

type UserRepository struct {
	DB *sql.DB
}

func (ur UserRepository) FindAll() (result []User, err error) {
	var users []User
	rows, err := ur.DB.Query("SELECT a.id, a.name, a.email, b.id FROM users as a INNER JOIN roles as b WHERE a.role = b.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			user  User
			id    int
			name  string
			email string
			role  int
		)
		if err := rows.Scan(&id, &name, &email, &role); err != nil {
			return nil, err
		}
		user.ID = id
		user.Name = name
		user.Email = email
		user.Role = role
		users = append(users, user)
	}
	return users, nil
}

func (ur UserRepository) Insert(user *User) (lastInsertedId int64, err error) {
	var hashedPassword []byte
	var insertResult sql.Result
	rowsCount := 0
	name := user.Name
	email := user.Email
	password := user.Password
	role := user.Role

	rows, err := ur.DB.Query("SELECT name, email FROM users WHERE name=? AND email=?", name, email)
	defer rows.Close()
	if err != nil {
		return -1, err
	}
	for rows.Next() {
		rowsCount++
	}
	if rowsCount != 0 {
		return 0, nil
	}

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	stmt, err := ur.DB.Prepare("INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	insertResult, err = stmt.Exec(&name, &email, &hashedPassword, &role)
	if err != nil {
		return -1, err
	}
	lastInsertID, err := insertResult.LastInsertId()
	if err != nil {
		return -1, err
	}
	return lastInsertID, nil
}

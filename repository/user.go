package repository

import (
	"database/sql"
	. "filemanagerAPI/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func (ur UserRepository) FindAll() (result []User, err error) {
	var users []User
	rows, err := ur.DB.Query("SELECT a.id, a.name, a.email, b.id as role FROM users as a INNER JOIN roles as b WHERE a.role = b.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			user  User
			id    int64
			name  string
			email string
			role  int64
		)
		if err := rows.Scan(&id, &name, &email, &role); err != nil {
			return nil, err
		}
		user.ID = id
		user.Name = name
		user.Email = email
		user.Role = role
		users = append(users, user)
		fmt.Println(users)
	}
	if err := rows.Err(); err != nil {
		return nil, err
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

func (ur UserRepository) Login(email, password string) (user User, errorNumber int64) {
	var userResult User
	var id int64
	var name string
	var emailRes string
	var passwordRes string
	var role int64
	var foundCount = 0

	rows, err := ur.DB.Query("SELECT id, name, email, password, role FROM users WHERE email = ?", email)
	defer rows.Close()
	if err != nil {
		return User{}, 1
	}

	for rows.Next() {
		foundCount++
		if err := rows.Scan(&id, &name, &emailRes, &passwordRes, &role); err != nil {
			return User{}, 2
		}
	}
	if foundCount == 0 {
		return User{}, 3
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordRes), []byte(password))
	if err != nil {
		return User{}, 4
	}
	userResult.ID = id
	userResult.Name = name
	userResult.Email = emailRes
	userResult.Role = role
	return userResult, 0
}

package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func InitializeDatabase(connStr string) (*DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) GetUsers() ([]string, error) {
	rows, err := db.Query("SELECT postgres FROM animals")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var users []string
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		users = append(users, username)
	}
	return users, nil
}

func (db *DB) AddUser(username string) error {
	_, err := db.Exec("INSERT INTO animals (postgres) VALUES ($1)", username)
	return err
}

func (db *DB) DeleteUser(username string) error {
	_, err := db.Exec("DELETE FROM animals WHERE postgres = $1", username)
	return err
}

func (db *DB) UpdateUser(oldUsername, newUsername string) error {
	_, err := db.Exec("UPDATE animals SET postgres = $1 WHERE postgres = $2", newUsername, oldUsername)
	return err
}

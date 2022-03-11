package memorydb

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:petsalone.db?mode=memory")
	if err != nil {
		return db, err
	}

	// Create pets table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS pets (id INTEGER PRIMARY KEY, name TEXT, pet_type TEXT, missing_since DATETIME)")
	if err != nil {
		return db, err
	}

	_, err = statement.Exec()
	if err != nil {
		return db, err
	}

	// Create users table
	statement, err = db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, normalized_username TEXT, email TEXT, normalized_email, password_hash TEXT, email_confirmed INTEGER, lockout_enabled INTEGER)")
	if err != nil {
		return db, err
	}

	_, err = statement.Exec()
	if err != nil {
		return db, err
	}

	err = create_pets(db)
	if err != nil {
		return db, err
	}

	err = create_users(db)
	if err != nil {
		return db, err
	}

	// Test data
	//rows, _ := db.Query("SELECT * from pets")
	//var id int
	//var name string
	//var pet_type string
	//var missing_since time.Time
	//fmt.Println("testing data is in db")
	//for rows.Next() {
	//	rows.Scan(&id, &name, &pet_type, &missing_since)
	//	fmt.Println(strconv.Itoa(id)+":", name, pet_type, missing_since)
	//}

	return db, nil
}

func create_pets(db *sql.DB) error {
	statement, err := db.Prepare("INSERT INTO pets (name, pet_type, missing_since) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec("Max", "cat", time.Now().AddDate(0, 0, -6))
	if err != nil {
		return err
	}

	_, err = statement.Exec("Fluffy", "dog", time.Now().AddDate(0, 0, -10))
	if err != nil {
		return err
	}

	_, err = statement.Exec("Snowball", "ferret", time.Now().AddDate(0, 0, -2))
	if err != nil {
		return err
	}

	return nil
}

func create_users(db *sql.DB) error {
	username := "elmyraduff"
	email := "elmyraduff@petsalone.com"
	password := "MontanaMax!!"
	password_hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))

	statement, err := db.Prepare("INSERT INTO users (username, normalized_username, email, normalized_email, password_hash, email_confirmed, lockout_enabled) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(username, strings.ToUpper(username), email, strings.ToUpper(email), password_hash, "True", "False")
	if err != nil {
		return err
	}

	return nil
}

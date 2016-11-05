package model

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"os"

	_ "github.com/lib/pq"
)

var Database *sql.DB

func init() {
	var err error
	Database, err = sql.Open("postgres", "dbname=todo user=todo password=todo sslmode=disable")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func CheckDatabase() (err error) {
	statement := "select * from users"
	stmt, err := Database.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	return err
}

// Encrypt hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

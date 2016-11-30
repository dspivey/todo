package model

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"todo/log"

	"bytes"
	"os/exec"

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

// CheckDatabase will verify that the database exists and the application can connect to it
func CheckDatabase() (err error) {
	log.Info("Testing database connection...")
	statement := "select * from users"
	stmt, err := Database.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	return err
}

func RunScript(dbname string, user string, script string) (err error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	args := []string{
		"-d",
		dbname,
		"-U",
		user,
		"-h",
		"localhost",
		"-p",
		"5432",
		"-x",
		"-W",
		"-a",
		"--single-transaction",
		"-v",
		"ON_ERROR_STOP=1",
		"--pset",
		"pager=off",
		"-L log.txt",
		"-f",
		script,
	}

	cmd := exec.Command("psql", args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		log.Danger("Error running PSQL command: ", stderr.String())
	}

	log.Info(stdout.String())

	return err
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatal("Cannot generate UUID: ", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F

	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return uuid
}

// Encrypt hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

package model

import (
    "time"
)

type User struct {
    ID          int
    Name        string
    Email       string
    Password    string
    CreatedAt   time.Time
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	statement := `insert into users (name, email, password, created_at) 
    values ($1, $2, $3, $4) returning id, name, email, created_at`

	stmt, err := Database.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow(user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	return err
}

// Delete user from database
func (user *User) Delete() (err error) {
	statement := "delete from users where id = $1"
    stmt, err := Database.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	
    return err
}

// Update user information in the database
func (user *User) Update() (err error) {
	statement := "update users set name = $2, email = $3 where id = $1"
	stmt, err := Database.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.Email)

	return err
}

// Get all users in the database and returns it
func Users() (users []User, err error) {
	rows, err := Database.Query("SELECT id, name, email, password, created_at FROM users")
    if err != nil {
		return users, err
	}
    
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	rows.Close()

	return users, err
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Database.QueryRow("SELECT id, name, email, password, created_at FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	return user, err
}
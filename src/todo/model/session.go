package model

import (
	"time"
)

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

// CreateSession creates a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	//TODO: create class to pull SQL from files
	statement := `insert into sessions (uuid, email, user_id, created_at) 
    values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at`

	stmt, err := Database.Prepare(statement)
	if err != nil {
		return session, err
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(CreateUUID(), user.Email, user.ID, time.Now()).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)

	return session, err
}

// Session get the session for an existing user
func (user *User) Session() (session Session, err error) {
	statement := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1`

	session = Session{}
	err = Database.QueryRow(statement, user.ID).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)

	return session, err
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	statement := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1`

	err = Database.QueryRow(statement, session.UUID).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)

	if err != nil {
		valid = false
		return valid, err
	}

	if session.ID != 0 {
		valid = true
	}

	return valid, err
}

// DeleteByUUID session from database
func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = $1"
	stmt, err := Database.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.UUID)
	return err
}

// Get the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Database.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = $1", session.UserID).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	return user, err
}

// Delete all sessions from database
func SessionDeleteAll() (err error) {
	statement := "delete from sessions"

	_, err = Database.Exec(statement)

	return err
}

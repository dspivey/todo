package model

import (
	"errors"
	"net/http"
	"time"
)

type Session struct {
	SessionId int
	UUID      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// CreateSession creates a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	//TODO: create class to pull SQL from files
	statement := `insert into sessions (uuid, email, user_id, created_at) 
    values ($1, $2, $3, $4) returning session_id, uuid, email, user_id, created_at`

	stmt, err := Database.Prepare(statement)
	if err != nil {
		return session, err
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(CreateUUID(), user.Email, user.UserId, time.Now()).Scan(&session.SessionId, &session.UUID, &session.Email, &session.UserId, &session.CreatedAt)

	return session, err
}

// Session get the session for an existing user
func (user *User) Session() (session Session, err error) {
	statement := `SELECT session_id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1`

	session = Session{}
	err = Database.QueryRow(statement, user.UserId).Scan(&session.SessionId, &session.UUID, &session.Email, &session.UserId, &session.CreatedAt)

	return session, err
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	statement := `SELECT session_id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1`

	err = Database.QueryRow(statement, session.UUID).Scan(&session.SessionId, &session.UUID, &session.Email, &session.UserId, &session.CreatedAt)

	if err != nil {
		valid = false
		return valid, err
	}

	if session.SessionId != 0 {
		valid = true
	}

	return valid, err
}

// IsAuthenticated checks if the user is logged in and has a session, if not err is not nil
func IsAuthenticated(rw http.ResponseWriter, req *http.Request) (sess Session, err error) {
	cookie, err := req.Cookie("_cookie")
	if err == nil {
		sess = Session{UUID: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session.")
		}
	}

	return sess, err
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

// User gets the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Database.QueryRow("SELECT user_id, name, email, created_at FROM users WHERE user_id = $1", session.UserId).Scan(&user.UserId, &user.Name, &user.Email, &user.CreatedAt)

	return user, err
}

// SessionDeleteAll delete all sessions from database
func SessionDeleteAll() (err error) {
	statement := "delete from sessions"

	_, err = Database.Exec(statement)

	return err
}

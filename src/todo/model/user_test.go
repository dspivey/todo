package model

import (
	"database/sql"
	"testing"
)

// test data
var users = []User{
	{
		Name:     "Peter Jones",
		Email:    "peter@gmail.com",
		Password: "peter_pass",
	},
	{
		Name:     "John Smith",
		Email:    "john@gmail.com",
		Password: "john_pass",
	},
}

func Test_UserCreate(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	if users[0].ID == 0 {
		t.Errorf("No id or created_at in user")
	}
	u, err := UserByEmail(users[0].Email)
	if err != nil {
		t.Error(err, "User not created.")
	}
	if users[0].Email != u.Email {
		t.Errorf("User retrieved is not the same as the one created.")
	}
}

func Test_UserDelete(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	if err := users[0].Delete(); err != nil {
		t.Error(err, "- Cannot delete user")
	}
	_, err := UserByEmail(users[0].Email)
	if err != sql.ErrNoRows {
		t.Error(err, "- User not deleted.")
	}
}

func Test_UserUpdate(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	users[0].Name = "Random User"
	if err := users[0].Update(); err != nil {
		t.Error(err, "- Cannot update user")
	}
	u, err := UserByEmail(users[0].Email)
	if err != nil {
		t.Error(err, "- Cannot get user")
	}
	if u.Name != "Random User" {
		t.Error(err, "- User not updated")
	}
}

func Test_Users(t *testing.T) {
	setup()
	for _, user := range users {
		if err := user.Create(); err != nil {
			t.Error(err, "Cannot create user.")
		}
	}
	u, err := Users()
	if err != nil {
		t.Error(err, "Cannot retrieve users.")
	}
	if len(u) != 2 {
		t.Error(err, "Wrong number of users retrieved")
	}
	if u[0].Email != users[0].Email {
		t.Error(u[0], users[0], "Wrong user retrieved")
	}
}

func Test_CreateSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "Cannot create session")
	}
	if session.UserID != users[0].ID {
		t.Error("User not linked with session")
	}
}

func Test_GetSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "Cannot create session")
	}

	s, err := users[0].Session()
	if err != nil {
		t.Error(err, "Cannot get session")
	}
	if s.ID == 0 {
		t.Error("No session retrieved")
	}
	if s.ID != session.ID {
		t.Error("Different session retrieved")
	}
}

func Test_checkValidSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "Cannot create session")
	}

	id := session.ID

	s := Session{ID: id}
	valid, err := s.Check()
	if err != nil {
		t.Error(err, "Cannot check session")
	}
	if valid != true {
		t.Error(err, "Session is not valid")
	}

}

func Test_checkInvalidSession(t *testing.T) {
	setup()
	s := Session{ID: -1}
	valid, err := s.Check()
	if err == nil {
		t.Error(err, "Session is not valid but is validated")
	}
	if valid == true {
		t.Error(err, "Session is valid")
	}

}

func Test_DeleteSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "Cannot create session")
	}

	err = session.DeleteByID()
	if err != nil {
		t.Error(err, "Cannot delete session")
	}
	s := Session{ID: session.ID}
	valid, err := s.Check()
	if err == nil {
		t.Error(err, "Session is valid even though deleted")
	}
	if valid == true {
		t.Error(err, "Session is not deleted")
	}
}
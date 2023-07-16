package models

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	Todos     []Todo
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `
	INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err = Db.Exec(
		cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now(),
		time.Now(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func GetUser(id int) (user User, err error) {
	// user = User{}
	cmd := `SELECT id, uuid, name ,email, password, created_at, updated_at FROM users WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
	}
	return
}

func (user *User) UpdateUser() (err error) {
	cmd := `UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ?`
	_, err = Db.Exec(cmd, user.Name, user.Email, time.Now(), user.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (user *User) DeleteUser() (err error) {
	cmd := `DELETE FROM users WHERE id = ?`
	_, err = Db.Exec(cmd, user.ID)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetUserByEmail(email string) (user User, err error) {
	// user = User{}
	cmd := `
	SELECT
		id,
		uuid,
		name,
		email,
		password,
		created_at,
		updated_at
	FROM
		users
	WHERE
		email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (u *User) CreateSession() (session Session, err error) {
	createCmd := `
	INSERT INTO sessions (
		uuid,
		email,
		user_id,
		created_at,
		updated_at
	) VALUES (
		?, ?, ?, ?, ?
	)`
	Db.Exec(createCmd, createUUID(), u.Email, u.ID, time.Now(), time.Now())

	getCmd := `
	SELECT
		id,
		uuid,
		email,
		user_id,
		created_at,
		updated_at
	FROM
		sessions
	WHERE
		user_id = ? AND email = ?`
	err = Db.QueryRow(getCmd, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
		&session.UpdatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (session *Session) CheckSession() (valid bool, err error) {
	cmd := `
	SELECT id, uuid, email, user_id, created_at, updated_at FROM sessions WHERE uuid = ?`
	err = Db.QueryRow(cmd, session.UUID).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
		&session.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		valid = false
		return
	}
	if session.ID != 0 {
		valid = true
	}
	return
}

func (session *Session) DeleteSessionByUUID() (err error) {
	cmd := `DELETE FROM sessions WHERE uuid = ?`
	_, err = Db.Exec(cmd, session.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (session *Session) GetUserBySession() (user User, err error) {
	cmd := `SELECT
		id,
		uuid,
		name,
		email,
		password,
		created_at,
		updated_at
	FROM
		users
	WHERE
		id = ?`
	err = Db.QueryRow(cmd, session.ID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return
}
package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) CreateTodo(content string) (err error) {
	cmd := `INSERT INTO todos (content, user_id, created_at, updated_at) VALUES (?, ?, ?, ?)`
	_, err = Db.Exec(cmd, content, user.ID, time.Now(), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}
	cmd := `SELECT * FROM todos WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func GetTodos() (todos []Todo, err error) {
	cmd := `SELECT * FROM todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return
}

func (user *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `SELECT * FROM todos WHERE user_id = ?`
	rows, err := Db.Query(cmd, user.ID)
	for rows.Next() {
		var todo Todo
		rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return
}

func (todo *Todo) UpdateTodo() (err error) {
	cmd := `UPDATE todos SET content = ?, updated_at = ? WHERE id = ?`
	_, err = Db.Exec(cmd, todo.Content, time.Now(), todo.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (todo *Todo) DeleteTodo() (err error) {
	cmd := `DELETE FROM todos WHERE id = ?`
	_, err = Db.Exec(cmd, todo.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
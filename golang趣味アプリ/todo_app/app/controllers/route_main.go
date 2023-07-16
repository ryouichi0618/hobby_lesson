package controllers

import (
	"log"
	"net/http"
	"time"
	"todo_app/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHtml(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generateHtml(w, user, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHtml(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoCreate(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Panicln(err)
		}
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		if err = user.CreateTodo(content); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := session.GetUserBySession()
		if err != nil {
			log.Panicln(err)
		}
		todo, err := models.GetTodo(id)
		generateHtml(w, todo, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		todo := &models.Todo{ID: id, Content: content, UserID: user.ID, UpdatedAt: time.Now()}
		if err := todo.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err = session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todo, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		if err := todo.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}

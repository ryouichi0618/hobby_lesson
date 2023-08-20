package main

import (
	"fmt"
	"log"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)
	log.Fatalln(controllers.StartMainServer())
	// u, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(u)

	// session, err := u.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)
	// valid, _ := session.CheckSession()
	// fmt.Println(valid)

}

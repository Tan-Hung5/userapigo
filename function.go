// main.go
package main

import (
	"log"
	"net/http"
	"usersservice/database"
	"usersservice/domain/services"
	"usersservice/middleware"
	"usersservice/repository"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	usersRepository := repository.NewUserRepository(db)

	usersService := services.NewUserService(usersRepository)

	router := mux.NewRouter()

	router.HandleFunc("/users", middleware.GetUsersHandler(usersService)).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", middleware.GetUserByIDHandler(usersService)).Methods("GET")
	router.HandleFunc("/users", middleware.AddUserHandler(usersService)).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", middleware.UpdateUserHandler(usersService)).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", middleware.DeleteUserHandler(usersService)).Methods("DELETE")

	log.Println("Server is running on :8080...")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)

}

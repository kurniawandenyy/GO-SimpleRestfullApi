package main

import (
	"fmt"
	"log"
	"net/http"
	user "./controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getUsers", user.ReturnAllUsers).Methods("GET")
	router.HandleFunc("/addUser", user.InsertUser).Methods("POST")
	router.HandleFunc("/updateUser/{id}", user.UpdateUser).Methods("PUT")
	router.HandleFunc("/deleteUser/{id}", user.DeleteUser).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 1000")
	log.Fatal(http.ListenAndServe(":1000", router))
}

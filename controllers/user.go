package controller

import (
	"encoding/json"
	"log"
	"../models"
	"../configs"
	"net/http"

	"github.com/gorilla/mux"
)

func ReturnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users model.Users
	var arr_user []model.Users
	var response model.Response

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * from person")
	
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_user = append(arr_user, users)
		}
	}
	response.Status = 1
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	db := config.Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	first_name := r.Form.Get("first_name")
	last_name := r.Form.Get("last_name")

	_, err = db.Exec("INSERT INTO person (first_name, last_name) values (?, ?)", first_name, last_name)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Data Added Successfully"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	db := config.Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	params := mux.Vars(r)
	id_person := params["id"]

	first_name := r.Form.Get("first_name")
	last_name := r.Form.Get("last_name")

	_, err = db.Exec("UPDATE person set first_name = ?, last_name = ? where id = ?", first_name, last_name, id_person)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Data Updated Successfully"
	log.Print("Update Data")

	w.Header().Set("Content-Type", "appllication/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	db := config.Connect()
	defer db.Close()

	params := mux.Vars(r)
	id_person := params["id"]
	_, err := db.Exec("Delete from person where id = ?", id_person)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Data Deleted Successfully"
	log.Print("Delete data from database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

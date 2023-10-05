package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	database "program-rme-be/Database"
	models "program-rme-be/Models"
)

func ListUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := database.DB.Query("SELECT id, nama, username, email, alamat, no_telp FROM users")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.Users

	for rows.Next() {
		var user models.Users
		if err = rows.Scan(&user.Id, &user.Nama, &user.Username, &user.Email, &user.Alamat, &user.NoTelp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stmt, err := database.DB.Prepare("INSERT INTO users(username,email,password) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	username := keyVal["username"]
	email := keyVal["email"]
	password := keyVal["password"]

	_, err = stmt.Exec(username, email, password)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println("User sudah ditambahkan")
}

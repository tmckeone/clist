package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Input struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Message struct {
	Message string `json:"message"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Panic(err)
	}

	db, err := sql.Open("postgres", Pgsqlconn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	var message string

	fmt.Println(input)
	var hash string

	row := db.QueryRow("SELECT hash from public.users WHERE username = $1", input.Username)
	err = row.Scan(&hash)
	switch err {
	case nil:
		message = hash
	default:
		message = "Invalid"
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(input.Password))
	if err != nil {
		message = "Invalid"
	}

	json.NewEncoder(w).Encode(Message{message})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Panic(err)
	}

	db, err := sql.Open("postgres", Pgsqlconn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	var message string
	var exists bool

	row := db.QueryRow("SELECT username FROM public.users WHERE username = $1", input.Username)
	err = row.Scan()
	switch err {
	case sql.ErrNoRows:
		exists = false
		message = "Registration Successful"
	default:
		exists = true
		message = "Username Already Exists"
	}

	if !exists {
		password, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
		if err != nil {
			log.Panic(err)
		}

		_, err = db.Exec("INSERT INTO public.users (username, hash) VALUES ($1, $2)", input.Username, password)
		if err != nil {
			log.Panic(err)
		}
	}

	json.NewEncoder(w).Encode(Message{message})
}

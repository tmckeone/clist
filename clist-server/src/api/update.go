package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type UpdateStatusStruct struct {
	Id   string `json:"id"`
	Open bool   `json:"open"`
}

type ChangeTagStruct struct {
	Id  string `json:"id"`
	Tag string `json:"tag"`
}

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var update UpdateStatusStruct

	db, err := sql.Open("postgres", Pgsqlconn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Panic(err)
	}

	_, err = db.Exec("UPDATE public.tickets SET open = $1 WHERE id=$2", update.Open, update.Id)
	if err != nil {
		log.Panic(err)
	}

}

func ChangeTag(w http.ResponseWriter, r *http.Request) {
	var update ChangeTagStruct

	db, err := sql.Open("postgres", Pgsqlconn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Panic(err)
	}

	_, err = db.Exec("UPDATE public.tickets SET tag = $1 WHERE id=$2", update.Tag, update.Id)
	if err != nil {
		log.Panic(err)
	}
}

func Assign(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		return
	}

	users, ok := r.URL.Query()["user"]
	if !ok || len(users[0]) < 1 {
		return
	}

	db, err := sql.Open("postgres", Pgsqlconn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE public.tickets SET assigned_to = $1 WHERE id=$2", users[0], ids[0])
	if err != nil {
		log.Panic(err)
	}
}

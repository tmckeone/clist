package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

type NewTicket struct {
	Subject    string `json:"subject"`
	Body       string `json:"body"`
	Tag        string `json:"tag"`
	AssignedTo string `json:"assigned_to"`
}

type ViewTicketStruct struct {
	Id         string   `json:"id"`
	Subject    string   `json:"subject"`
	Body       string   `json:"body"`
	Tag        string   `json:"tag"`
	Replies    []string `json:"replies"`
	AssignedTo string   `json:"assigned_to"`
	Open       bool     `json:"open"`
}

type ViewListStruct struct {
	Id      string `json:"id"`
	Subject string `json:"subject"`
}

type ReplyStruct struct {
	Id    string `json:"id"`
	Reply string `json:"reply"`
	User  string `json:"user"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket NewTicket

	db, err := sql.Open("postgres", Pgsqlconn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		log.Panic(err)
	}

	_, err = db.Exec("INSERT INTO public.tickets (subject, body, tag, assigned_to) VALUES ($1, $2, $3, $4)", ticket.Subject, ticket.Body, ticket.Tag, ticket.AssignedTo)
	if err != nil {
		log.Panic(err)
	}
}

func ViewTicket(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		return
	}

	id := ids[0]

	db, err := sql.Open("postgres", Pgsqlconn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	var ticket ViewTicketStruct

	row := db.QueryRow("SELECT id, subject, body, tag, replies, assigned_to, open FROM public.tickets WHERE id=$1", id)
	err = row.Scan(&ticket.Id, &ticket.Subject, &ticket.Body, &ticket.Tag, pq.Array(&ticket.Replies), &ticket.AssignedTo, &ticket.Open)
	switch err {
	case sql.ErrNoRows:
		json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("Ticket ID: %s not found", id)})
	case nil:
		json.NewEncoder(w).Encode(ticket)
	default:
		log.Panic(err)
	}

}

func ReplyToTicket(w http.ResponseWriter, r *http.Request) {
	var ticket ReplyStruct

	db, err := sql.Open("postgres", Pgsqlconn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		log.Panic(err)
	}

	ticket.Id = strings.Trim(ticket.Id, "\n")

	id, err := strconv.Atoi(ticket.Id)
	if err != nil {
		http.NotFound(w, r)
	}

	now := time.Now()

	reply := fmt.Sprintf("Reply from %s at %s:\n\n%s", ticket.User, now.Format("Mon Jan _2 15:04:05 2006"), ticket.Reply)

	db.Exec("UPDATE public.tickets SET replies = array_append(replies, $1) where id=$2", reply, id)
}

func ViewMyTickets(w http.ResponseWriter, r *http.Request) {
	users, ok := r.URL.Query()["user"]
	if !ok || len(users[0]) < 1 {
		return
	}

	user := users[0]

	var tickets []ViewListStruct
	var ticket ViewListStruct

	db, err := sql.Open("postgres", Pgsqlconn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, subject FROM public.tickets WHERE assigned_to=$1 AND open=true ORDER BY id ASC", user)
	if err != nil {
		log.Panic(err)
	}

	for rows.Next() {
		rows.Scan(&ticket.Id, &ticket.Subject)

		tickets = append(tickets, ticket)
	}

	json.NewEncoder(w).Encode(tickets)
}

func SearchTicket(w http.ResponseWriter, r *http.Request) {
	var searchType string
	var queryArg string

	users, ok := r.URL.Query()["user"]
	if ok && len(users[0]) > 0 {
		searchType = "user"
		queryArg = users[0]
	}

	matches, ok := r.URL.Query()["match"]
	if ok && len(matches[0]) > 0 {
		searchType = "match"
		queryArg = "%" + matches[0] + "%"
	}

	tags, ok := r.URL.Query()["tag"]
	if ok && len(tags[0]) > 0 {
		searchType = "tag"
		queryArg = tags[0]
	}

	fmt.Println(searchType)
	fmt.Println(queryArg)

	var tickets []ViewListStruct
	var ticket ViewListStruct

	db, err := sql.Open("postgres", Pgsqlconn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	var rows *sql.Rows

	switch searchType {
	case "match":
		rows, err = db.Query("SELECT id, subject FROM public.tickets WHERE subject ILIKE $1 AND open=true ORDER BY id ASC", queryArg)
		if err != nil {
			log.Panic(err)
		}
	case "user":
		rows, err = db.Query("SELECT id, subject FROM public.tickets WHERE assigned_to=$1 AND open=true ORDER BY id ASC", queryArg)
		if err != nil {
			log.Panic(err)
		}
	case "tag":
		rows, err = db.Query("SELECT id, subject FROM public.tickets WHERE tag=$1 AND open=true ORDER BY id ASC", queryArg)
		if err != nil {
			log.Panic(err)
		}
	default:
		fmt.Println("No search type")
		return
	}

	/*if err == sql.ErrNoRows {
		fmt.Println("No rows found")
		json.NewEncoder(w).Encode([]ViewListStruct{{Id: "error", Subject: "No tickets found"}})
		return
	}*/

	for rows.Next() {
		rows.Scan(&ticket.Id, &ticket.Subject)

		tickets = append(tickets, ticket)
	}

	json.NewEncoder(w).Encode(tickets)
}

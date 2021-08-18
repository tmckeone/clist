package api

import (
	"net/http"
)

type TestResp struct {
	Content string `json:"content"`
}

const Pgsqlconn = "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable"

func MasterHandler() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/login", Login)
	mux.HandleFunc("/api/register", Register)
	mux.HandleFunc("/api/create", CreateTicket)
	mux.HandleFunc("/api/reply", ReplyToTicket)
	mux.HandleFunc("/api/view/", ViewTicket)
	mux.HandleFunc("/api/mytickets/", ViewMyTickets)
	mux.HandleFunc("/api/search/", SearchTicket)
	mux.HandleFunc("/api/update/status", UpdateStatus)
	mux.HandleFunc("/api/update/tag", ChangeTag)
	mux.HandleFunc("/api/assign/", Assign)

	return mux
}

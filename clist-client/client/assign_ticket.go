package client

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// calls the corresponding endpoint on the server to assign a ticket to yourself.
func Take(id string) {
	client := http.Client{Timeout: time.Second}

	_, err := client.Get(fmt.Sprintf("http://localhost:8080/api/assign/?id=%s&user=%s", id, "carl.stucky"))
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Ticket %s is now assigned to you.", id)
}

// calls the corresponding endpoint on the server to assign a ticket to another user.
func Assign(id string, user string) {
	client := http.Client{Timeout: time.Second}

	_, err := client.Get(fmt.Sprintf("http://localhost:8080/api/assign/?id=%s&user=%s", id, user))
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Ticket %s is now assigned to %s.", id, user)
}

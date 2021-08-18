package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

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

func ViewTicket(id string) {

	// send GET request to the server with id as a URL parameter
	client := http.Client{Timeout: time.Second}

	resp, err := client.Get(fmt.Sprintf("http://localhost:8080/api/view/?id=%s", id))
	if err != nil {
		log.Panic(err)
	}

	// save response to a struct
	var ticket ViewTicketStruct

	err = json.NewDecoder(resp.Body).Decode(&ticket)
	if err != nil {
		log.Panic(err)
	}

	// format / print struct
	printViewTicket(ticket)
}

func ViewMyTickets() {

	// send a GET request from the server with the current username as a URL parameter
	client := http.Client{Timeout: time.Second}

	resp, err := client.Get(fmt.Sprintf("http://localhost:8080/api/mytickets/?user=%s", "troy.mckeone"))
	if err != nil {
		log.Panic(err)
	}

	// decode response into a slice of struct
	var tickets []ViewListStruct

	err = json.NewDecoder(resp.Body).Decode(&tickets)
	if err != nil {
		log.Panic(err)
	}

	// print the slice of struct
	fmt.Println("*********************************")
	fmt.Println("Open tickets that are assigned to me:")
	for _, ticket := range tickets {
		printTicketList(ticket)
	}
	fmt.Println("*********************************")

}

// print a ticket ID and Subject
func printTicketList(ticket ViewListStruct) {
	sep := "---------------------------------"
	fmt.Println(sep)
	fmt.Println("ID:", ticket.Id, "|", "Subject:", ticket.Subject)
}

// print the full view of a ticket
func printViewTicket(ticket ViewTicketStruct) {
	sep := "---------------------------------"

	if ticket.Id == "" {

		fmt.Println("No ticket found")
		return
	}

	var status string
	if ticket.Open {
		status = "Status: Open"
	} else {
		status = "Status: Closed"
	}

	fmt.Println("*********************************")
	fmt.Println("Ticket ID: ", ticket.Id)
	fmt.Println(sep)
	fmt.Println(status)
	fmt.Println(sep)
	fmt.Printf("Subject:\n\n%s\n", ticket.Subject)
	fmt.Println(sep)
	fmt.Println("Tag:", ticket.Tag)
	fmt.Println(sep)
	fmt.Println("Assigned To:", ticket.AssignedTo)
	fmt.Println(sep)
	fmt.Printf("Body: \n\n%s\n", ticket.Body)

	if len(ticket.Replies) != 0 {
		fmt.Println(sep)
		for _, reply := range ticket.Replies {
			fmt.Println(reply)
			fmt.Println(sep)
		}
	}

	fmt.Println("*********************************")
}

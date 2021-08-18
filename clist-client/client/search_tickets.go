package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func SearchQuery() {
	// read user input and save it as a variable
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the query string and press enter")
	query, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}

	// clean unwanted characters from the input
	query = cleanString(query)

	// send a GET request to the server, passing the input as a URL parameter
	client := http.Client{Timeout: time.Second}

	resp, err := client.Get(fmt.Sprintf("http://localhost:8080/api/search/?match=%s", query))
	if err != nil {
		log.Panic(err)
	}

	// save the response to a struct from view_tickets.go
	var tickets []ViewListStruct

	err = json.NewDecoder(resp.Body).Decode(&tickets)
	if err != nil {
		log.Panic(err)
	}

	// format / print the struct
	fmt.Println("*********************************")
	fmt.Println("Tickets Found:")
	for _, ticket := range tickets {
		printTicketList(ticket) // view_ticket.go
	}
	fmt.Println("*********************************")
}

func SearchTag(tag string) {
	// clean unwanted characters from the user input
	tag = cleanString(tag)

	// send a GET request to the server with the input as a URL parameter
	client := http.Client{Timeout: time.Second}

	resp, err := client.Get(fmt.Sprintf("http://localhost:8080/api/search/?tag=%s", tag))
	if err != nil {
		log.Panic(err)
	}

	// save the response to a struct from view_tickets.go
	var tickets []ViewListStruct

	err = json.NewDecoder(resp.Body).Decode(&tickets)
	if err != nil {
		log.Panic(err)
	}

	// format / print the struct
	fmt.Println("*********************************")
	fmt.Println("Tickets Found:")
	for _, ticket := range tickets {
		printTicketList(ticket) // view_ticket.go
	}
	fmt.Println("*********************************")
}

func SearchUser(user string) {
	// clean unwanted characters from user input
	user = cleanString(user)

	// send a GET request to the server with the input as a URL parameter
	client := http.Client{Timeout: time.Second}

	resp, err := client.Get(fmt.Sprintf("http://localhost:8080/api/search/?user=%s", user))
	if err != nil {
		log.Panic(err)
	}

	// save the response to a struct from view_tickets.go
	var tickets []ViewListStruct

	err = json.NewDecoder(resp.Body).Decode(&tickets)
	if err != nil {
		log.Panic(err)
	}

	// format / print the output
	fmt.Println("*********************************")
	fmt.Println("Tickets Found:")
	for _, ticket := range tickets {
		printTicketList(ticket) // view_ticket.go
	}
	fmt.Println("*********************************")
}

// clean unwanted characters from the input string, return the clean string.
func cleanString(s string) string {
	s = strings.Trim(s, "\n")
	s = strings.Trim(s, "\r")
	s = strings.ReplaceAll(s, " ", "%20")

	return s
}

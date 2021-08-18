package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type NewTicket struct {
	Subject    string `json:"subject"`
	Body       string `json:"body"`
	Tag        string `json:"tag"`
	AssignedTo string `json:"assigned_to"`
}

func CreateTicket() {
	// read user input and save it as variables
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the ticket subject and then press enter")
	subject, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Please enter the ticket tag and then press enter")
	tag, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Enter the content of your body, type ~ and enter")
	body, err := reader.ReadString(byte('~'))
	if err != nil {
		log.Panic(err)
	}

	// clean unwated characters from the input
	subject = strings.Trim(subject, "\n")
	tag = strings.Trim(tag, "\n")
	body = strings.Trim(body, "~")

	// pass the input to a JSON object
	ticket := NewTicket{Subject: subject, Tag: tag, Body: body, AssignedTo: "troy.mckeone"}

	postData, err := json.Marshal(&ticket)
	if err != nil {
		log.Panic(err)
	}

	// send the JSON object to the server
	client := http.Client{Timeout: time.Second}

	_, err = client.Post("http://localhost:8080/api/create", "application/json", bytes.NewBuffer([]byte(postData)))
	if err != nil {
		log.Panic(err)
	}

}

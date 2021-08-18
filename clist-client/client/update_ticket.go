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

type ReplyStruct struct {
	Id    string `json:"id"`
	Reply string `json:"reply"`
	User  string `json:"user"`
}

type UpdateStatusStruct struct {
	Id   string `json:"id"`
	Open bool   `json:"open"`
}

type ChangeTagStruct struct {
	Id  string `json:"id"`
	Tag string `json:"tag"`
}

func Reply(id string) {

	// read user input and save it as a variable
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your reply and when done, type ~ then enter to exit.")
	reply, err := reader.ReadString('~')
	if err != nil {
		log.Panic(err)
	}

	// clean unwanted characters from the input
	reply = strings.Trim(reply, "~")

	// create a struct with the input
	ticket := ReplyStruct{Id: id, Reply: reply, User: "troy.mckeone"}

	// pass the struct to a JSON object
	postData, err := json.Marshal(&ticket)
	if err != nil {
		log.Panic(err)
	}

	// send a POST request to the server with the JSON object
	client := http.Client{Timeout: time.Second}

	_, err = client.Post("http://localhost:8080/api/reply", "application/json", bytes.NewBuffer([]byte(postData)))
	if err != nil {
		log.Panic(err)
	}
}

func UpdateStatus(id string, open bool) {

	// save input as a struct
	update := UpdateStatusStruct{Id: id, Open: open}

	// pass input struct to a JSON object
	postData, err := json.Marshal(&update)
	if err != nil {
		log.Panic(err)
	}

	// send a POST request to the server with the JSON object
	client := http.Client{Timeout: time.Second}

	_, err = client.Post("http://localhost:8080/api/update/status", "application/json", bytes.NewBuffer([]byte(postData)))
	if err != nil {
		log.Panic(err)
	}

	// print what the status was changed to
	if open {
		fmt.Println("Changed status of", id, "to Open")
	} else {
		fmt.Println("Changed status of", id, "to Closed")
	}
}

func ChangeTag(id string, tag string) {

	// save input to a struct
	update := ChangeTagStruct{Id: id, Tag: tag}

	// pass input to a JSON object
	postData, err := json.Marshal(&update)
	if err != nil {
		log.Panic(err)
	}

	// send a POST request to the server with the JSON object
	client := http.Client{Timeout: time.Second}

	_, err = client.Post("http://localhost:8080/api/update/tag", "application/json", bytes.NewBuffer([]byte(postData)))
	if err != nil {
		log.Panic(err)
	}

	// print the update
	fmt.Println("Updated tag of", id, "to", tag)
}

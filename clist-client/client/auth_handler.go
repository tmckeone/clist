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

type Input struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Message struct {
	Message string `json:"message"`
}

//Token is the client's representation of the password hash.
var Token Message

//AttachToken is used to pass Token to the server in order to verify which user is accessing an endpoint
func AttachToken(req *http.Request) {
	if Token.Message != "" {
		req.Header.Add("Authorization", "Bearer "+Token.Message)
	}
}

func Login() {

	// read user input and save it as variables.
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Username:")

	username, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Password:")

	password, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}

	// clean any unwanted characters from the input strings
	username = cleanString(username)
	password = cleanString(password)

	// pass the input into a JSON object
	input := Input{Username: username, Password: password}

	postData, err := json.Marshal(&input)
	if err != nil {
		log.Panic(err)
	}

	// send the JSON object to the server
	client := http.Client{Timeout: time.Second}

	resp, err := client.Post("http://localhost:8080/api/login", "application/json", bytes.NewBuffer([]byte(postData)))
	if err != nil {
		log.Panic(err)
	}

	// decode the server response into Token
	err = json.NewDecoder(resp.Body).Decode(&Token)
	if err != nil {
		log.Panic(err)
	}

	// check if the token is valid and tell the user
	message := Token.Message

	if strings.Contains(Token.Message, "Invalid") {
		fmt.Println(message)
		message = ""
	} else {
		message = "Login Successful!"
		fmt.Println(message)
	}

}

func Register() {

	// read user input and save it as variables
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your desired username:")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Enter your desired password:")
	password, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}

	// clean any unwanted characters
	username = cleanString(username)
	password = cleanString(password)

	// pass input into a JSON object
	input := Input{Username: username, Password: password}

	postData, err := json.Marshal(&input)
	if err != nil {
		log.Panic(err)
	}

	// send the JSON object to the server
	client := http.Client{Timeout: time.Second}

	resp, err := client.Post("http://localhost:8080/api/register", "application/json", bytes.NewBuffer([]byte(postData)))
	if err != nil {
		log.Panic(err)
	}

	// print the response from the server
	var message Message
	err = json.NewDecoder(resp.Body).Decode(&message)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(message.Message)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	FilePath := "user.json"
	JsonFileHandler, err := os.Open(FilePath)
	if err != nil {
		log.Fatal("File not found!")
	}
	defer JsonFileHandler.Close()

	fmt.Println("Successfully open file:", FilePath)

	byteVal, _ := ioutil.ReadAll(JsonFileHandler)

	var users Users
	json.Unmarshal(byteVal, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Name:", users.Users[i].Name)
		fmt.Println("User Age:", users.Users[i].Age)
		fmt.Println("Facebook:", users.Users[i].Social.Facebook)
		fmt.Println("Twitter:", users.Users[i].Social.Twitter)
	}

}

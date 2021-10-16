package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name"`
	Type    string   `xml:"type,attr"`
	Social  Social   `xml:"social"`
}

type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func main() {
	FilePath := "user.xml"
	XMLFileHandler, err := os.Open(FilePath)
	if err != nil {
		log.Fatal("file not found:", err)
	}
	defer XMLFileHandler.Close()
	fmt.Println("successfully open:", FilePath)

	byteVal, _ := ioutil.ReadAll(XMLFileHandler)

	var users Users
	xml.Unmarshal(byteVal, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Name:", users.Users[i].Name)
		//fmt.Println("User Age:", users.Users[i].Age)
		fmt.Println("Facebook:", users.Users[i].Social.Facebook)
		fmt.Println("Twitter:", users.Users[i].Social.Twitter)
	}
}

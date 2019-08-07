package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Users is the complete array of all users in the file
type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

// User struct, this contains our
// Type attribute, our user's name and
// a social struct which will contain all
// our social links
type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

// Social , a simple struct which contains all our
// social links
type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func main() {

	// Open XML file
	xmlFile, err := os.Open("test.xml")
	// if os.Open returns an error than handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened test.xml")
	defer xmlFile.Close()

	// read the opened xmlFile as a byte array
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// initialize Users array
	var users Users
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		fmt.Println("Twitter Url: " + users.Users[i].Social.Twitter)
		fmt.Println("YouTube Url: " + users.Users[i].Social.Youtube)
	}

}

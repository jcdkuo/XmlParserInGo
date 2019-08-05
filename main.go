package main

import (
	"encoding/xml"
	"fmt"
	//"io/ioutil"
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

}

/*
Objective: log file contents a data of customer with email addresses. Program should find out all unique addresses and displays duplicate values of email id.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	emails          []string
	uniqueEmails    []string
	duplicateEmails []string
)

func findDuplicateValues(id []string) {
	m := make(map[string]bool)
	for _, id := range id {
		_, ok := m[id]
		if ok == true {
			// email id is present in map , means duplicate
			// fmt.Println("Duplicate id is", id)
			duplicateEmails = append(duplicateEmails, id)

		} else {
			m[id] = true
		}
	}
	// map m contains all unique values
	for item, _ := range m {
		uniqueEmails = append(uniqueEmails, item)
	}
	fmt.Println("Number of duplicate id's is -->", len(duplicateEmails), "and those are are -->", uniqueEmails)
	fmt.Println("Number od uniques email id's is -->", len(uniqueEmails), "and those are are -->", duplicateEmails)
}

func main() {
	// open a file emailsorterdata.log
	file, err := os.Open("emailsorterdata.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		emailField := strings.Fields(lineText)
		// remove header line for email
		if emailField[3] != "email" {
			//fmt.Println(emailField[3]) // all email addresses will be printed
			// get all email addresses stored in a slice by trimming double quotes
			emails = append(emails, strings.Trim(emailField[3], "\""))

		}

	}
	// fmt.Println(emails) // prints a slice of all email addresses
	findDuplicateValues(emails)
}

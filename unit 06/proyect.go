package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Save contacts in Json file
func saveContactToFile(contact []Contact) error {
	file, err := os.Create("contact.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	//Serialize data
	err = encoder.Encode(contact)
	if err != nil {
		return err
	}
	return nil
}

// Load contacts from a Json file
func loadContactFromFile(contacts *[]Contact) error {
	file, err := os.Open("contact.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	//Slice for contacts
	var contacts []Contact
	//Load contacts that exits in the file selected
	err := loadContactFromFile(&contacts)
	if err != nil {
		fmt.Println("Error to load contacts ", err)
	}
	// Create instance for bufio
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		fmt.Print("====Management Contact =====\n",
			"1.Add New Contact\n",
			"2.Show Contacts\n",
			"3.Exit \n",
			"4.Choose one option: ")

		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch option {
		case 1:
			//Enter a new contact
			var contact Contact
			fmt.Print("Name : ")
			contact.Name, _ = reader.ReadString('\n')
			fmt.Print("Email : ")
			contact.Email, _ = reader.ReadString('\n')
			fmt.Print("Phone : ")
			contact.Phone, _ = reader.ReadString('\n')
			//Add contact to slice
			contacts = append(contacts, contact)
			//Save in a json file
			if err := saveContactToFile(contacts); err != nil {
				fmt.Println("Error to save a contact", err)
			}
		case 2:
			fmt.Println("=========================================================")
			for index, contact := range contacts {
				fmt.Println(index+1, ". Contact")
				fmt.Print("Name : ", contact.Name)
				fmt.Print("Email : ", contact.Email)
				fmt.Print("Phone : ", contact.Phone)
			}
			fmt.Println("=========================================================")
		case 3:
			return
		default:
			fmt.Println("Invalid option")
		}

	}
}

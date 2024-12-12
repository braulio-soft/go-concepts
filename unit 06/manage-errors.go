package main

import (
	"log"
	"os"
)

func main() {
	log.SetPrefix("main(): ")
	file, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("Hey im your first log ")

}

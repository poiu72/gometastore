package main

import (
	"fmt"
	"log"

	"github.com/poiu72/gometastore/hmsclient"
)

func main() {
	client, err := hmsclient.Open("localhost", 9083)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client.GetAllDatabases())
}

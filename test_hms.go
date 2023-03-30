package main

import (
	"fmt"
	"log"

	"github.com/poiu72/gometastore/hmsclient"
)

func main() {
	client, err := hmsclient.Open("xxxx", 9083, true, "/Users/gfjia/Downloads/cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client.GetAllDatabases())
}

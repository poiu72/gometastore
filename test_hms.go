package main

import (
	"fmt"
	"log"

	"github.com/poiu72/gometastore/hmsclient"
)

func main() {
	client, err := hmsclient.Open("xxx", 9083, true, false, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client.GetAllDatabases())
}

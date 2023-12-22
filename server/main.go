package main

import (
	"creatures/paths"
	"fmt"
	"net/http"
)

const SRV_PORT string = ":8190"
const SRV_IP string = "localhost"
const SRV_STRING string = "http://" + SRV_IP + SRV_PORT

func main() {
	//open server

	http.Handle("/", http.FileServer(http.Dir("../static")))
	http.HandleFunc("/random_creature", paths.GetRandomCreature)
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("../fonts/"))))
	fmt.Printf("server open at %s", SRV_STRING)

	err := http.ListenAndServe(SRV_PORT, nil)
	if err != nil {
		fmt.Printf("cant open server %v", err)
	}
	fmt.Printf("server open on %s", SRV_PORT)
}

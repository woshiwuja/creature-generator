package main

import (
	//	"anarchy2036/models"

	//	"anarchy2036/queries"
	"fmt"
	//	"io"
	"net/http"
	//	"os"
	//	"strings"
	//	"text/template"
	//"gorm.io/driver/postgres"
)

const SRV_PORT string = "8190"
const SRV_STRING string = "http://localhost:" + SRV_PORT

func main() {
	//open server

	http.Handle("/", http.FileServer(http.Dir("../static")))

	fmt.Printf("server open at %s", SRV_STRING)

	err := http.ListenAndServe(":"+SRV_PORT, nil)
	if err != nil {
		fmt.Printf("cant open server %v", err)
	}
	fmt.Printf("server open on %s", SRV_PORT)
}

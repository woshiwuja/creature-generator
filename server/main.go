package main

import (
	"anarchy2036/db_conn"
//	"anarchy2036/models"
	"anarchy2036/paths"
//	"anarchy2036/queries"
	"fmt"
//	"io"
	"net/http"
//	"os"
//	"strings"
//	"text/template"

	"gorm.io/gorm"
	//"gorm.io/driver/postgres"
)

const SRV_PORT string = ":2036"

var DB *gorm.DB



func main() {
	//open server

	var error error
	DB, error = db_conn.DBConnect("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")
	if error != nil {
		fmt.Printf("problem connecting to db %v", error)
	}
	fmt.Printf("connected to db %v", &DB)

	http.Handle("/", http.FileServer(http.Dir("../static")))
	http.HandleFunc("/hello", paths.GetHello)
	//http.HandleFunc("/clock", paths.GetClock)
	http.HandleFunc("/stack", paths.GetStack)
	http.HandleFunc("/blog/post1", paths.GetBlogPost1)
	http.HandleFunc("/blog/post2", paths.GetBlogPost2)
	http.HandleFunc("/blog", paths.GetBlog)
	http.HandleFunc("/chart", paths.GetUsers(DB))
	http.HandleFunc("/minigame", paths.GetMinigame)
	http.HandleFunc("/minigame1", paths.GetMinigame1)
	http.HandleFunc("/minigame2", paths.GetMinigame2)

	fmt.Printf("server open on %s", SRV_PORT)

	err := http.ListenAndServe(SRV_PORT, nil)
	if err != nil {
		fmt.Printf("cant open server %v", err)
	}
	//DB CONNECTION
	fmt.Printf("server open on %s", SRV_PORT)
}

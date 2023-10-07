package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

const NAME string = "anarchy2036"
const SRV_PORT string = ":3333"

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/hello.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	data := map[string]interface{}{
		"Name": NAME,
	}
	builder := &strings.Builder{}
	htmlTemplate := string(htmxFile)
	template := template.Must(template.New("hello").Parse(htmlTemplate))
	if err := template.Execute(builder, data); err != nil {
		panic(err)
	}
	s := builder.String()
	io.WriteString(w, s)
} //THIS WORKS
func getStack(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/stack.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	data := map[string]interface{}{
		"Name": NAME,
	}
	builder := &strings.Builder{}
	htmlTemplate := string(htmxFile)
	template := template.Must(template.New("hello").Parse(htmlTemplate))
	if err := template.Execute(builder, data); err != nil {
		panic(err)
	}
	s := builder.String()
	io.WriteString(w, s)
}
func getClock(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/clock request received\n")
	var clock = time.Now()
	htmxFile, err := os.ReadFile("../static/components/clock.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	data := map[string]interface{}{
		"Clock": clock.Format("15:04:05"),
	}
	builder := &strings.Builder{}
	htmlTemplate := string(htmxFile)
	template := template.Must(template.New("hello").Parse(htmlTemplate))
	if err := template.Execute(builder, data); err != nil {
		panic(err)
	}
	s := builder.String()
	io.WriteString(w, s)
} //THIS WORKS
func main() {
	fmt.Printf("server open on %s", SRV_PORT)
	http.Handle("/", http.FileServer(http.Dir("../static")))
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/clock", getClock)
	http.HandleFunc("/stack", getStack)

	err := http.ListenAndServe(SRV_PORT, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed")
	} else if err != nil {
		fmt.Printf("error %s", err)
		os.Exit(1)
	}
}

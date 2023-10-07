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
const SRV_PORT string = ":2036"

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
func getRant(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/rant.html")
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

func getMinigame1(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/minigame/minigame1.html")
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
func getMinigame2(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/minigame/minigame2.html")
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
func getMinigame(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/minigame.html")
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
	http.HandleFunc("/rant", getRant)
	http.HandleFunc("/minigame", getMinigame)
	http.HandleFunc("/minigame1", getMinigame1)
	http.HandleFunc("/minigame2", getMinigame2)
	err := http.ListenAndServe(SRV_PORT, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed")
	} else if err != nil {
		fmt.Printf("error %s", err)
		os.Exit(1)
	}
}

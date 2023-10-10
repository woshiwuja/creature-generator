package main

import (
	"encoding/json"
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
type dataStruct struct {
	Username string `json:"username"`
	EventCount int `json:"event_count"`
}
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
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
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
func getApiChart(w http.ResponseWriter, r *http.Request){
	chartData := []dataStruct {{Username: "root",EventCount: 24,},{Username:"admin",EventCount:12}}
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(chartData)
}
func getBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/blog.html")
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
func getChart(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/chart request received\n")
	data,err := http.Get("http://localhost:2036/api/chartdata")
	if err != nil {
		fmt.Printf("cant load data :(")
	}
	htmxFile, err := os.ReadFile("../static/components/chart.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	b, err := io.ReadAll(data.Body)
	
	if err != nil{
		fmt.Printf("cant read body")
	}
	fmt.Printf("%s",string(b))
	var chart []dataStruct
	json.Unmarshal(b, &chart)
	/*chartDataMap := map[string]interface{}{
		"username":
		"event_count":chart.EventCount,
	}*/
	builder := &strings.Builder{}
	htmlTemplate := string(htmxFile)
	template := template.Must(template.New("chart").Parse(htmlTemplate))
	if err := template.Execute(builder, chart); err != nil {
		panic(err)
	}
	s := builder.String()
	io.WriteString(w, s)
}
func getBlogPost1(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/blog/post1.html")
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
func getBlogPost2(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/blog/post2.html")
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
	http.HandleFunc("/blog/post1", getBlogPost1)
	http.HandleFunc("/blog/post2", getBlogPost2)
	http.HandleFunc("/blog", getBlog)
	http.HandleFunc("/chart", getChart)
	http.HandleFunc("/api/chartdata", getApiChart)
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

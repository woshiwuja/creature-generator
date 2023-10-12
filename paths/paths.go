package paths

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
	"io"
	"strings"
	"gorm.io/gorm"
	"anarchy2036/queries"
	"anarchy2036/models"
)


func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/hello.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
} //THIS WORKS
func GetStack(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/stack.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("./static/components/blog.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}

func GetBlogPost1(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/blog/post1.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}
func GetBlogPost2(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("./static/components/blog/post2.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}
func GetMinigame1(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/minigame/minigame1.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))

}
func GetMinigame2(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/minigame/minigame2.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}
func GetMinigame(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("/hello request received\n")
	htmxFile, err := os.ReadFile("../static/components/minigame.html")
	if err != nil {
		fmt.Printf("error reading file")
	}
	io.WriteString(w, string(htmxFile))
}

func GetUsers(DB *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Users []models.User
		Users = queries.GetUsers(DB)
		htmxFile, err := os.ReadFile("../static/components/chart.html")
		if err != nil {
			fmt.Printf("error reading file %v", err)
		}
		htmlTemplate := string(htmxFile)
		template := template.Must(template.New("hello").Parse(htmlTemplate))
		builder := &strings.Builder{}
		template.Execute(builder, Users)
		s := builder.String()
		io.WriteString(w, s)
	}
}

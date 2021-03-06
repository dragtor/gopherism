package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type arc string

type apiHandler struct {
	Book map[arc]StoryArc
}

func (a apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlList := strings.Split(r.URL.Path, "/")
	fmt.Fprintln(os.Stdout, urlList, len(urlList))
	if len(urlList) != 3 {
		fmt.Fprintf(w, fmt.Sprintf("%s %s", "Invalid url", string(r.URL.Path)))
	}
	if _, ok := a.Book[arc(urlList[2])]; !ok {
		fmt.Fprintf(w, fmt.Sprintf("%s", "page not found"))
		return
	}
	storyArc := a.Book[arc(urlList[2])]
	title := fmt.Sprintf("<h3>%s</h3>\n", storyArc.Title)
	paras := ""
	for _, para := range storyArc.Story {
		paras = paras + "<div>&nbsp;&nbsp;&nbsp;&nbsp;" + para + "</div>"
	}
	options := ""
	for _, option := range storyArc.Options {
		opt := fmt.Sprintf(`<div><a href="http://localhost:8080/api/%s">%s</a></div>`, option.Arc, option.Text)
		options += opt
	}

	fmt.Fprintf(w, "<html><head><title></title></head><body>%s %s %s</body></html>", title, paras, options)

}

func main() {
	file, err := os.Open("samples/gopher.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	filesize := fileinfo.Size()

	buffer := make([]byte, filesize)
	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}
	//var js JsonStruct
	var result map[string]StoryArc
	err = json.Unmarshal(buffer, &result)
	if err != nil {
		panic(err)
	}
	arcSize := len(result)
	b := make(map[arc]StoryArc, arcSize)
	for k, v := range result {
		b[arc(k)] = v
	}
	t, err := template.New("book").Parse(`{{define "Book"}}{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "Book", "<script>alert('you have been pwned')</script>")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{Book: b})
	http.ListenAndServe(":8080", mux)
}

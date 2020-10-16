package main

import (
	"flag"
	"fmt"
	"github.com/dragtor/gopherism/urlshortner/pkg"
	"net/http"
	"os"
)

func mapHandler(pathToUrls map[string]string, mux *http.ServeMux) http.HandlerFunc {
	redirect := func(w http.ResponseWriter, r *http.Request) {
		if _, ok := pathToUrls[r.RequestURI]; !ok {
			fmt.Fprint(w, "<html><head><title>Page Unregister</title></head><body>Redirection Not available</body></html>")
			return
		}
		http.Redirect(w, r, pathToUrls[r.RequestURI], 301)
	}
	return redirect
}

var (
	enableYamlInput *bool
	path            *string
)

func init() {
	enableYamlInput = flag.Bool("ep", false, "Enable path to yaml")
	path = flag.String("p", "samples/redirection.yaml", "path location")
	flag.Parse()
}

func readRedirectionYaml(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return buffer, nil

}

func handlerLogic(method string, mux *http.ServeMux) http.HandlerFunc {
	switch method {
	case "MAP":
		pathsToUrls := map[string]string{
			"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
			"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		}
		handler := pkg.MapHandler(pathsToUrls, mux)
		return handler

	case "YAML":
		yaml, err := readRedirectionYaml(*path)
		if err != nil {
			panic(err)
		}
		yamlHandler, err := pkg.YAMLHandler([]byte(yaml), mux)
		if err != nil {
			panic(err)
		}
		return yamlHandler
	case "DB":
		break
	}
	return nil
}

func configSelectionPolicy() string {
	if *enableYamlInput {
		return "YAML"
	}
	return "MAP"
}

func main() {
	mux := defaultMux()
	policy := configSelectionPolicy()
	fmt.Printf("selection policy %s  enableYamlInput : %v\n", policy, *enableYamlInput)
	handler := handlerLogic(policy, mux)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pong")
}

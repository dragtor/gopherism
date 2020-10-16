package main

import (
	"flag"
	"fmt"
	"github.com/dragtor/gopherism/urlshortner/pkg"
	"net/http"
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
	enableYamlInput = flag.Bool("-ep", false, "Enable path to yaml")
	path = flag.String("-p", "../samples/redirect.yaml", "path location")
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
	fmt.Printf("selection policy %s\n", policy)
	//var handler func(http.ResponseWriter, *http.Request)
	handler := handlerLogic(policy, mux)

	/*
		yaml := `
			- path: /urlshort
			  url: https://github.com/gophercises/urlshort
			- path: /urlshort-final
			  url: https://github.com/gophercises/urlshort/tree/solution
			`
		yamlHandler, err := pkg.YAMLHandler([]byte(yaml), mapHandler)
		if err != nil {
			panic(err)
		}*/
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

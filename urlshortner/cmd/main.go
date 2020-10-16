package main

import (
	"fmt"
	"net/http"
	//"github.com/dragtor/gopherism/urlshortner/pkg"
)

//func mapHandler(pathToUrls map[string]string, mux *http.ServeMux) http.HandlerFunc {	return http.HandlerFunc(http.ReirectHandler("www.google.com", 200))

//}

func mapHandler(pathToUrls map[string]string, mux *http.ServeMux) http.HandlerFunc {
	url := "https://www.facebook.com"
	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, 301)
	}

	return redirect
}
func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	//mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	mapHandler := mapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	/*
			yaml := `
		- path: /urlshort
		  url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		  url: https://github.com/gophercises/urlshort/tree/solution
		`
		   yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
			if err != nil {
				panic(err)
			}
	*/
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

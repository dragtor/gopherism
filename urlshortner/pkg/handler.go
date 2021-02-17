package pkg

import (
	"net/http"
    "fmt"
    yaml "gopkg.in/yaml.v2"
    "log"
)

func MapHandler(pathToUrls map[string]string, mux *http.ServeMux) http.HandlerFunc {
     redirect := func(w http.ResponseWriter, r *http.Request) {
         if _, ok := pathToUrls[r.RequestURI]; !ok {
             fmt.Fprint(w, "<html><head><title>Page Unregister</title></head><body>Redirection Not    available</body></html>")
             return
         }
         http.Redirect(w, r, pathToUrls[r.RequestURI], 301)
     }
     return redirect
 }

func YAMLHandler(yml []byte, mux *http.ServeMux) (http.HandlerFunc, error) {
 fmt.Println(string(yml))
  parsedYaml, err := parseYAML(yml)
  if err != nil {
    log.Print(err)
    return nil, err
  }
  pathMap := buildMap(parsedYaml)
  fmt.Printf("%+v",pathMap)
  return MapHandler(pathMap, mux), nil
}

func buildMap(parsedYaml *Yaml) map[string]string {
    redirectionMap := make(map[string]string,len(*parsedYaml))
    for _ , data := range *parsedYaml {
        redirectionMap[data.Path] = data.URL
    }
    return redirectionMap 
}

type Yaml []struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYAML(yamldata []byte) (*Yaml, error){
    var yml Yaml
    err := yaml.Unmarshal(yamldata, &yml)
    if err != nil {
        return nil , nil 
    }
    return &yml, nil
}

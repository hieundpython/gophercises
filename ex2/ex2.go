package ex2

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	return mux
}

func Exec() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/google":    "https://google.com.vn",
		"/thanhnien": "https://thanhnien.vn",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	yaml := `
- path: /google
  url: https://google.com.vn
- path: /thanhnien
  url: https://thanhnien.vn/
`
	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8000", yamlHandler))
}

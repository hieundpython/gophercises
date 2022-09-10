package ex3

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func handleContent(store Store) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		pathUrl := strings.Trim(r.URL.Path, "/")

		// load template
		bookStoreTp, err := template.ParseFiles("./ex3/template.html")
		if err != nil {
			log.Fatalln(err)
		}
		// return with template
		chapter := store[pathUrl]
		err = bookStoreTp.Execute(w, chapter)
		if err != nil {
			log.Fatalln(err)
		}

	}
}

func Exec() {
	fileFlag := flag.String("filePath", "./ex3/book.json", "json file")

	flag.Parse()

	file, err := os.Open(*fileFlag)

	if err != nil {
		log.Fatalln(err)
	}

	store, err := DecodeJson(file)
	if err != nil {
		log.Fatalln(err)
	}

	mux := http.NewServeMux()

	for key := range store {
		mux.HandleFunc("/"+key, handleContent(store))
	}

	log.Fatal(http.ListenAndServe("localhost:3000", mux))

}

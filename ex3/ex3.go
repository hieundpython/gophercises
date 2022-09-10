package ex3

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World")
}

func Exec() {
	file, err := os.Open("./ex3/book.json")

	if err != nil {
		log.Fatalln(err)
	}

	store, err := DecodeJson(file)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(store)

	// Create http handle for handle all routing here

	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe("localhost:3000", mux))

}

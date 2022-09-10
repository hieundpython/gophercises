package ex3

import (
	"encoding/json"
	"os"
)

func DecodeJson(f *os.File) (Store, error) {

	d := json.NewDecoder(f)

	var store Store
	err := d.Decode(&store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

type Store map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

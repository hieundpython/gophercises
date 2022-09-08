package ex2

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(ymlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathUrl

	err := yaml.Unmarshal(ymlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}

	pathToUrls := make(map[string]string)

	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.Url
	}

	return MapHandler(pathToUrls, fallback), nil
}

type pathUrl struct {
	Path string
	Url  string
}

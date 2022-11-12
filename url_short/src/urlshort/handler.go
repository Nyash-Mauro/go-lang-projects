package urlshort

import (
	"net/http"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func NewBaseUrlMapper(urls map[string]string) func(string) (string, bool) {
	return func(path string) (string, bool) {
		url, ok := urls[path]
		return url, ok
	}
}
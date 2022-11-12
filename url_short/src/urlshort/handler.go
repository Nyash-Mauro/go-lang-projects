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


func NewYamlUrlMapper(filename string) (func(string) (string, bool), error) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	yml := []map[string]string{}
	err = yaml.Unmarshal(content, &yml)

	if err != nil {
		return nil, err
	}

	mapping := make(map[string]string)

	for _, m := range yml {
		mapping[m["path"]] = m["url"]
	}

	return NewBaseUrlMapper(mapping), nil
}


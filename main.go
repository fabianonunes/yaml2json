package main

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	v, err := yaml.YAMLToJSON(input)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	_, _ = os.Stdout.WriteString(string(v))
}

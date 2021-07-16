package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

type JSONObject map[string]interface{}

func readFile(filename string) ([]byte, error) { //TODO: Also add a param to parse JSON or YAML
	return ioutil.ReadFile(filename)
}

type JQStep interface {

}

type ObjStep struct {
	isSet bool
	index int
}
type ArrStep struct {
	isSet bool
	index int
}

type JQPath []JQStep

func parsePath(path string) (JQPath, error) {
	var res JQPath
	pathLen := len(path)
	i := 0
	for i < pathLen {

	}
	return res, nil
}

func main() {
	inputFile := *flag.String("f", "sample.json", "File to query")
	queryPath := *flag.String("p", ".", "JSON/YAML query string")
	isYAML := *flag.Bool("yaml", false, "Is the file YAML? (Rather than JSON)")

	flag.Parse()

	fmt.Println("You said...")
	fmt.Printf("'%s'\n", inputFile)
	fmt.Printf("'%t'\n", isYAML)
	fmt.Printf("'%s'\n", queryPath)
	fmt.Println()

	data, err := readFile(inputFile)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Your file:")
	fmt.Println(string(data))

	var res JSONObject
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Unmarshalled JSON:")
	fmt.Printf("%+v\n", res)
}

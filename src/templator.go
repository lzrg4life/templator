package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	if len(os.Args) != 4 {
		fmt.Println("You must provide exactly 3 arguments: <source> <base> <output>")
	}

	source := os.Args[1]
	base := os.Args[2]
	out := os.Args[3]

	sourceData, err := ioutil.ReadFile(source)
	check(err)

	var data interface{}
	err = json.Unmarshal(sourceData, &data)

	baseData, err := ioutil.ReadFile(base)
	check(err)

	templ, err := template.New("Base").Parse(string(baseData))
	check(err)

	f, err := os.Create(out)
	check(err)
	defer f.Close()

	err = templ.Execute(f, data)
	check(err)
}

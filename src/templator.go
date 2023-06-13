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
		fmt.Println("You must provide exactly 3 arguments: <sourceFileName> <templateFileName> <outFileName>")
	}

	sourceFileName := os.Args[1]
	templateFileName := os.Args[2]
	outFileName := os.Args[3]

	template := prepareTemplate(templateFileName)
	data := readSource(sourceFileName)
	executeTemplate(outFileName, template, data)
}

func readSource(sourceFileName string) interface{} {
	sourceData, err := ioutil.ReadFile(sourceFileName)
	check(err)

	var data interface{}
	err = json.Unmarshal(sourceData, &data)
	check(err)

	return data
}

func prepareTemplate(templateFileName string) *template.Template {
	baseData, err := ioutil.ReadFile(templateFileName)
	check(err)

	template, err := template.New("Base").Parse(string(baseData))
	check(err)

	return template
}

func executeTemplate(outFileName string, template *template.Template, data interface{}) {
	f, err := os.Create(outFileName)
	check(err)
	defer f.Close()

	err = template.Execute(f, data)
	check(err)
}

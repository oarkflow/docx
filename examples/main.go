package main

import (
	"github.com/oarkflow/docx"
)

func main() {
	input := "Employment Contract - {employee_name}.docx"
	doc, err := docx.Open(input, docx.PlaceholderMap{})
	if err != nil {
		panic(err)
	}
	err = doc.ReplaceAll(docx.PlaceholderMap{})
	if err != nil {
		panic(err)
	}
	doc.WriteToFile(input)
}

package main

import (
	"fmt"
	"github.com/cdornsife/dashf"
	"github.com/spf13/pflag"
	"os"
)

// TestStruct represents the data in the file
type TestStruct struct {
	One   int    `json:"one"`
	Two   string `json:"two"`
	Three []struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	} `json:"three"`
}

func main() {

	var inputFile string
	var out = new(TestStruct)

	pflag.StringVarP(&inputFile, "file", "f", "", "Input for create/update: json|yaml|-")
	pflag.Parse()

	err := dashf.Unmarshal(inputFile, &out)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", out)
}

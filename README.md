# DashF

[![Build Status](https://travis-ci.org/cdornsife/dashf.svg?branch=master)](https://travis-ci.org/cdornsife/dashf)

Unmarshall A file, STDIN or a URL into a golang struct. The supported formats are JSON and YAML. 

## What is DashF?

DashF means the flag -f.  When a CLI wishes to accept a resource (filename/directory/url) they typically use the -f flag. DashF provides your cli a consistent way to unmarshall user-defined resources into golang structs. Most CLI tools process the -f flag in the same manner, but the code is typically buried and not easily re-usable.

## Doesn't Viper/envconfig/pflag do this?
No. Viper and envconfig provide a great libraries for configuring your application. Viper and pflag can accept the input but provide no mechanism for processing the resource. 

## [Example](example/)

Here is some example input JSON. This can be located in a local file, a URL or even piped via STDIN.
```
{
  "one": 1,
  "two": "two",
  "three": [
    {
      "a": "a",
      "b": "b",
      "c": "c"
    }
  ]
}
```
Example code to handle the input.
```
package main

import (
        "fmt"
        "github.com/cdornsife/dashf"
        "github.com/spf13/pflag"
        "os"
)

// TestStruct represents the data in the json/yaml file
// Example files are in the tests directory
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
```

Command line usage:
```
$ cd example

# Pass in a - to use STDIN
$ cat ../tests/test.json | go run main.go -f -

# Specify a local file
$ go run main.go -f ../tests/test.yaml

# Specify a URL
go run main.go --file https://raw.githubusercontent.com/cdornsife/dashf/master/tests/test.yaml

```




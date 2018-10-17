package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"./browserUtil"
	"./requestHandler"
	"github.com/aymerick/douceur/parser"
	"github.com/napsy/go-css"
)

var parsedCSS = ""

func main() {
	if argsOK() {
		b, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			fmt.Print(err)
		}
		stylesheet, err := parser.Parse(string(b))
		if err != nil {
			fmt.Println(err)
			panic("Please fill a bug :)")
		}
		parsed, err := css.Unmarshal([]byte(stylesheet.String()))
		if err != nil {
			panic(err)
		}
		fmt.Printf("Defined rules:\n")
		parsedCSS = stylesheet.String()

		for k, v := range parsed {
			fmt.Println(k)
			for k, v := range v {
				// Inspect(k, v)
				fmt.Println(k)
				fmt.Println(v)
			}
		}
		http.HandleFunc("/dashboard", requestHandler.HandleRequest(parsedCSS))

		fmt.Println("Server is listening on port 8080")
		go browserUtil.LaunchBrowser("http://127.0.0.1:8080/dashboard")
		http.ListenAndServe(":8080", nil)
	} else {
		fmt.Println("Wrong arguments")
	}
}

/*
	Indicates whether CLI args are ok or not.
*/
func argsOK() bool {
	if len(os.Args) > 1 {
		if os.Args[1] == "--file" && len(os.Args[2]) > 0 {
			return true
		}
		return false
	}
	return false
}

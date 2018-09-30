package main

import (
	"fmt"
	"github.com/aymerick/douceur/parser"
	"github.com/napsy/go-css"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var parsedCSS = ""

func main() {
	if argsOK() {
		http.HandleFunc("/dashboard", displayInspections)
		b, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			fmt.Print(err)
		}
		stylesheet, err := parser.Parse(string(b))
		if err != nil {
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
				Inspect(k, v)
				fmt.Println(k)
				fmt.Println(v)
			}
		}

		fmt.Println("Server is listening on port 8080")
		go launchBrowser()
		http.ListenAndServe(":8080", nil)
	}
}

func launchBrowser() {
	url := "http://127.0.0.1:8080/dashboard"
	time.Sleep(1000 * time.Millisecond)
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
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

func displayInspections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, parsedCSS)
	return
}

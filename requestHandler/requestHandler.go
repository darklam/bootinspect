package requestHandler

import (
	"fmt"
	"net/http"
)

func HandleRequest(parsedCSS string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, parsedCSS)
	}
}

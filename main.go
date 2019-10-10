// main go

package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	romannumerals "github.com/maestre3d/romanserver/utils/mock"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")
	if urlPathElements[1] == "roman_number" {
		number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
		if number == 0 || number > 10 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not Found"))
		} else {
			fmt.Fprintf(w, "%q", html.EscapeString(romannumerals.Numerals[number]))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Listening on PORT: %q", s.Addr)
	s.ListenAndServe()
}

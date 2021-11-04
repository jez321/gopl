// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jez321/gopl/ch01/1.6/lissajous"
)

func main() {
	http.HandleFunc("/lissajous", lissajousHandler) // each request calls handler
	http.HandleFunc("/", echoHandler)               // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	cycles := r.URL.Query().Get("cycles")
	if cycles == "" {
		cycles = "5"
	}

	fmt.Println(cycles)

	cyclesInt, err := strconv.Atoi(cycles)
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
		return
	}

	lissajous.WriteCustom(w, cyclesInt)
}

package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func displayHTTPHeaders(w http.ResponseWriter, r *http.Request) {

	// Method
	fmt.Println("> Method: ", r.Method)

	// URL
	fmt.Println("> URL: ", r.URL)

	// HTTP version
	fmt.Println("> Proto: ", r.Proto)

	// Header
	fmt.Println("> Header: ")
	for k, v := range r.Header {
		fmt.Println(k, strings.Join(v, ""))
	}

	// Body
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	fmt.Println("> Body:")
	fmt.Println(buf.String())

	// ContentLength
	fmt.Println("> Content-Length: ", r.ContentLength)

	// Host
	fmt.Println("> Host: ", r.Host)

	// Remoet Address
	fmt.Println("> Remote Address: ", r.RemoteAddr)

	// Just send a 500

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - This server is just for debugging"))

}

func main() {
        fmt.Println("Starting HTTP listener on :9999")
	http.HandleFunc("/", displayHTTPHeaders)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

// The HelloWorld server
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			defer logResponse(r, http.StatusNotFound)
			http.NotFound(w, r)
			return
		}
		defer logResponse(r, http.StatusOK)
		fmt.Fprintf(w, "Hello, World!")
	})
	log.Fatal(http.ListenAndServe(config.Web.Host+":"+strconv.Itoa(config.Web.Port), nil))
}

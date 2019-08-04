// The HelloWorld server
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/antik9/microservice-go/internal/hello"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			defer hello.LogResponse(r, http.StatusNotFound)
			http.NotFound(w, r)
			return
		}
		defer hello.LogResponse(r, http.StatusOK)
		fmt.Fprintf(w, "Hello, World!")
	})
	log.Fatal(http.ListenAndServe(
		hello.ServerConfig.Web.Host+":"+strconv.Itoa(hello.ServerConfig.Web.Port), nil))
}

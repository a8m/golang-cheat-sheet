// Full example can be found at https://play.golang.org/p/pwWxdrQSrYv
// Go programs can embed static files using the "embed" package as follows:

package main

import (
	"embed"
	"log"
	"net/http"
)

// content holds the static content (2 files) for the web server.
// go:embed a.txt b.txt
var content embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(content)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

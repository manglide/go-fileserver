// main
package main

import (
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

func main() {
	portS := os.Getenv("PORT")

	if portS == "" {
		portS = "8080"
	}

	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+portS, nil)
}

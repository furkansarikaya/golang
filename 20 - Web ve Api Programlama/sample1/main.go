package main

import (
	"fmt"
	"net/http"
)

//https://github.com/cihanozhan/golang-webapi-samples/blob/master/00_webApi/main.go

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba %s", r.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)

	fmt.Println("Web Server")
}

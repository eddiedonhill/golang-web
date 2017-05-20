package main

import (
        "net/http"
        "io"
        "fmt"
)

func main() {
        http.HandleFunc("/", index)
        http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="images/basketball.jpg">	`)
}

func basketballPic (w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "basketball.jpg")
}

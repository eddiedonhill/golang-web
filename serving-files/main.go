package main

import (
        "net/http"
        "io"
        "fmt"
)

func main() {
        http.HandleFunc("/", index)
        http.Handle("/home", home)
        http.Handle("/about", about)
        http.Handle("/contact", contact)
        http.Handle("/faq", faq)
        http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="dog1.jpg">	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog1.jpg")
}

func home (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="dog2.jpg">	`)
}

func dogPic2(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog2.jpg")
}

func about (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="dog3.jpg">	`)
}

func dogPic3(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog3.jpg")
}

func contact (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="dog4.jpg">	`)
}

func dogPic4(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog4.jpg")
}

func faq (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="dog5.jpg">	`)
}

func dogPic5(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog5.jpg")
}

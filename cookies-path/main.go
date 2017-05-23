package main

import (
    "net/http"
    "html/template"
)

var tpl *template.Template

func init() {
            tpl = template.Must(template.ParseGlob("templates/*.gothtml"))
}

func main() {
        http.HandleFunc("/", index)
        http.HandleFunc("/dog/bowzer", bowzer)
        http.HandleFunc("/dog/bowzer/pictures", bowzerpics)
        http.HandleFunc("/cat", cat)
        http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
            tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func bowzer(w http.ResponseWriter, r *http.Request) {
            tpl.ExecuteTemplate(w, "bowzer.gohtml", nil) {
                c := &http.Cookie {
                    Name: "user-cookie",
                    Value: "this would be the value",
                    Path: "/dog/bowzer",
                }
                tpl.ExecuteTemplate(w, "bowzer.gohtml", c)
            }
}

func bowzerpics(w http.ResponseWriter, r *http.Request) {
            car c *http.Cookie
            c, err := r.Cookie("user-cookie")
            if err != nil {
                    fmt.Println(err)
                    fmt.Println("%T\n", C)
                    fmt.Printf("%T\n", C)
            }
            fmt.Println("c", c)
            fmt.Printf("%T\n", c)
            tpl.ExecuteTemplate(w, "bowzerpics.gohtml", C)
}

func cat(w http.ResponseWriter, r *http.Request) {
            var c *http.Cookie
            c, err := r.Cookie("user-cookie")
            if err != nil {
                    fmt.Println(err)
                    fmt.Printf("%T\n", C)
            }
            tpl.ExecuteTemplate(w, "cat.gohtml", nil)
}

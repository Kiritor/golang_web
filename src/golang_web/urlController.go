package main

import (
    "log"
    "net/http"
    "html/template"
)

type User struct {
    UserName string
}

type urlController struct {
}

func (this *urlController)IndexView(w http.ResponseWriter, r *http.Request, user string) {
    t, err := template.ParseFiles("template/html/index.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, &User{user})
}

func (this *urlController)LoginView(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("template/html/login.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}

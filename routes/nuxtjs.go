package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"html/template"
)


func NuxtjsRoutes(r *mux.Router) {
    
    r.PathPrefix("/_nuxt").Handler(http.StripPrefix("/_nuxt", http.FileServer(http.Dir("./nuxtjs/dist/_nuxt"))))

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/{two}", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/{two}/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/{two}/{three}", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/{two}/{three}/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/{two}/{three}/{four}", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/{two}/{three}/{four}/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/{two}/{three}/{four}/{five}", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/{two}/{three}/{four}/{five}/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/{two}/{three}/{four}/{five}/{six}", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/{two}/{three}/{four}/{five}/{six}/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/{two}/{three}/{four}/{five}/{six}/{seven}", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })
    r.HandleFunc("/admiin/{one}/{two}/{three}/{four}/{five}/{six}/{seven}/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("nuxtjs/dist/admiin/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, false)
    })                               
}
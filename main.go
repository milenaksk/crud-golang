package main

import (
    "database/sql"
    "fmt"
    "net/http"
)

func Read(w http.ResponseWriter, r *http.Request) {
    // Sua implementação aqui
}

func Create(w http.ResponseWriter, r *http.Request) {
    // Sua implementação aqui
}

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", "postgres://root:root@postgres/crud?sslmode=disable")
    if err != nil {
        panic(err)
    }

    if err = db.Ping();  err != nil {
        panic(err)
    }

    fmt.Println("You are connected to your database.")
}

func main() {
    http.HandleFunc("/users/read", Read)
    http.HandleFunc("/users/create", Create)
    http.ListenAndServe(":8080", nil)
}

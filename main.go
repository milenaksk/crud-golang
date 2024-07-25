package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct { //entidade e atributos
    Id int 
    Name  string
    Email string 
    Age int
}

//métodos

func Read(w http.ResponseWriter, r *http.Request) {
    // Sua implementação aqui
}

func Create(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed) //405 não permitido
        return
    }

    u := User{}
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
        fmt.Println("server failed to handle ", err)
        return
    }

    _,err = db.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)", u.Name, u.Email, u.Age)
    if err != nil {
        fmt.Println("server failed to handle ", err)
        return
    }

    w.WriteHeader(http.StatusCreated)
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

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
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed) //405 não permitido
        return
    }

    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        fmt.Println("server failed to handle ", err)
        return
    }

    defer rows.Close()

    data := make([]User, 0)

    for rows.Next() {
        user := User{}
        err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age)
        if err != nil {
            fmt.Println("server failed to handle ", err)
            return
        }
        data = append(data, user)
    }

    if err = rows.Err(); err != nil {
        fmt.Println("server failed to handle ", err)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)
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

func Update(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {
    if r.Method != "DELETE" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed) //405 não permitido
        return
    }

    id := r.URL.Query().Get("id") // uga uga: adicionar declaração da variável `id`

    _,err := db.Exec("DELETE FROM users WHERE id=$1;", id)
    if err != nil {
        fmt.Println("server failed to handle ", err)
        return
    }

    w.WriteHeader(http.StatusOK)
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
    http.HandleFunc("/users/update", Update)
    http.HandleFunc("/users/delete", Delete)
    http.HandleFunc("/users/read", Read)
    http.HandleFunc("/users/create", Create)
    http.ListenAndServe(":8080", nil)
}

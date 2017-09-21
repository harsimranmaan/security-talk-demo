package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html"
	"html/template"
	"log"
	"net/http"
)

func insecureHandler(w http.ResponseWriter, r *http.Request) {
	var username = r.FormValue("username")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Hi %s! <br />", username)
	if username != "" {
		_, err := db.Exec(fmt.Sprintf("INSERT INTO users(username) values('%s')", username))
		log.Println(err)
	}
	users := getUsers()
	for _, user := range users {
		fmt.Fprintf(w, "| %d | %s | <br/>", user.id, user.username)
	}

}
func relativelySecureHandler(w http.ResponseWriter, r *http.Request) {
	var username = html.EscapeString(r.FormValue("username"))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Hi %s! <br />", username)
	if username != "" {
		stmt, err := db.Prepare("INSERT INTO users(username) values(?)")
		log.Println(err)
		_, err = stmt.Exec(username)
		log.Println(err)
	}
	users := getUsers()
	for _, user := range users {
		//double escaping but that's beside the point
		fmt.Fprintf(w, "| %d | %s | <br/>", user.id, html.EscapeString(user.username))
	}
}



func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func getUsers() []User {
	var users []User
	var id int
	var username string
	rows, err := db.Query("SELECT uid,username FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err = rows.Scan(&id, &username); err != nil {
			log.Fatal(err)
		}
		users = append(users, User{id, username})
	}
	return users
}

type User struct {
	id       int
	username string
}

var db *sql.DB

func main() {
	var err error
	if db, err = sql.Open("sqlite3", "./securitydemo"); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/secure", relativelySecureHandler)
	http.HandleFunc("/insecure", insecureHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8500", nil)
}

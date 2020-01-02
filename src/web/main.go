package main

import (
	"fmt"
	"net/http"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type Post struct {
	Id int
	Title string
	Body string
}

var db, err = sql.Open("mysql", "root:@/go_lang?charset=utf8")

func main() {

	// stmt, err := db.Prepare("INSERT INTO posts(title, body) values(?,?)")
	// checkErr(err)

	// _, err = stmt.Exec("My First Post", "My first content")
	// checkErr(err)

	rows, err := db.Query("SELECT * FROM posts")
	checkErr(err)

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.id, &post.Title, &post.Body)
		items = append(items, post)
	}

	println(itens)

	db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		post := Post{Id: 1, Title: "Undefined", Body: "No content..."}

		if title := r.FormValue("title"); title != "" {
			post.Title = title
		}

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "index.html", post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
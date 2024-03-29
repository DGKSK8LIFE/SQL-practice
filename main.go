package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ShoppingCart struct {
	gorm.Model
	Item string
}

var (
	tpl    *template.Template
	tplTwo *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("main.html"))
	tplTwo = template.Must(template.ParseGlob("searchsite.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/adding", processor)
	http.HandleFunc("/adding/searchsite", parseThenQuery)
	http.HandleFunc("/searchsite", parseThenQuery)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "main.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "shoppinglist.sqlite")
	if err != nil {
		panic("failed to connect to db")
	}
	defer db.Close()
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	item := r.FormValue("item")

	db.AutoMigrate(&ShoppingCart{})
	db.Create(&ShoppingCart{Item: item})
	db.Model(&ShoppingCart{Item: item})

	tpl.ExecuteTemplate(w, "main.html", item)
}

func parseThenQuery(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "availableItems.sqlite")
	if err != nil {
		panic("failed to connect to db")
	}
	searchedItem := r.FormValue("item")
	defer db.Close()
	tplTwo.ExecuteTemplate(w, "searchsite.html", nil)

	rows, _ := db.Query("SELECT * FROM itemsInStock;")
	var searched string
	for rows.Next() {
		rows.Scan(&searched)
		if searchedItem == searched {
			fmt.Printf("%s is in stock", searched)
			break
		}
	}
}

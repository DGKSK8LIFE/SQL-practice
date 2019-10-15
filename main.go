package main

import (
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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("main.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/adding", processor)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "main.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "storedata.db")
	if err != nil {
		panic("failed to connect to db")
	}
	defer db.Close()
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	item := r.FormValue("item")

	fmt.Println(item)

	db.AutoMigrate(&ShoppingCart{})
	db.Create(&ShoppingCart{Item: item})
	db.Model(&ShoppingCart{Item: item})

	tpl.ExecuteTemplate(w, "main.html", item)
}

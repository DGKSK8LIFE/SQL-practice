package main

import (
	"net/http"
	"html/template"
) 

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("main.html"))
}


func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/adding", processor)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "main.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	item := r.FormValue("item")	

	d := struct{
		Item string
	}{
		Item: item,
	}

	tpl.ExecuteTemplate(w, "main.html", d)
} 

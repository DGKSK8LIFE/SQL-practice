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
	
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "main.html", nil)
}
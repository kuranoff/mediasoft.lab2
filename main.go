package main

import (
	"html/template"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"./packages/cosine"
	"./packages/dataset"
)

type PageData struct {
	Title string
}

func root(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "@Куранов",
	}
	indx, _ := template.ParseFiles("templates/index.html")
	indx.Execute(w, data)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", root)
	http.HandleFunc("/upload", dataset.Upload)
	http.HandleFunc("/cosinesimilarity", cosine.Similarity)
	http.Handle("/metrics", promhttp.Handler())

	server.ListenAndServe()
}

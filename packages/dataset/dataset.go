package dataset

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"./ai"
	"./db"
	"./preproc"
)

// PageUploadData ...
type PageUploadData struct {
	Title          string
	UploadFileName string
	Label          string
}

// Upload ...
func Upload(w http.ResponseWriter, r *http.Request) {

	dl := r.FormValue("datasetLabel")

	err := r.ParseMultipartForm(200000)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	file, handler, err := r.FormFile("originalFile")
	if err != nil {
		fmt.Println("Ошибка получения файла")
		fmt.Println(err)
		return
	}
	defer file.Close()

	resFile, err := os.Create("./upload/" + handler.Filename)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	defer resFile.Close()

	io.Copy(resFile, file)
	defer resFile.Close()

	data := PageUploadData{
		Title:          "Dataset",
		UploadFileName: handler.Filename,
		Label:          dl,
	}

	// Нормализация dataset
	preproc.TextPreprocessing(data.UploadFileName)

	// Создание векторной модели
	ai.FastText(data.UploadFileName)

	// Запись в БД
	db.NewDataset(data.UploadFileName)

	upld, _ := template.ParseFiles("templates/upload.html")
	upld.Execute(w, data)

}

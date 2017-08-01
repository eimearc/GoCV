package main

import (
	"fmt"
	"net/http"
	"os"
	"encoding/gob"
	"log"

	"github.com/eimearc/latex"
)

var page latex.Page
var filePath string = "/root/go/src/github.com/eimearc/latex.txt"

func upload() latex.Page {
	fmt.Println("Upload file.")

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		log.Fatal("Error opening file", filePath, err)
	}

	var result latex.Page
	enc := gob.NewDecoder(f)
	err = enc.Decode(&result)
	if err != nil {
		log.Fatal("Error decoding page from", f, err)
	}

	fmt.Println("File has been successfuly gob decoded.")
	return result
}

func download() {
	fmt.Println("Download file.")

	f, err := os.Create(filePath)
	defer f.Close()
	if err != nil {
		log.Fatal("Error opening file", filePath, err)
	}

	enc := gob.NewEncoder(f)
	err = enc.Encode(page)
	if err != nil {
		log.Fatal("Error encoding page", page, err)
	}

	fmt.Println("File has been successfuly gob encoded.")
}

func createPDF() {
	fmt.Println("Create PDF.")
	latex.CreatePDF()
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handle upload.")
	fmt.Fprintf(w, fmt.Sprintf("%#v\n", upload()))
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handle download.")
	download()
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handle create.")
	createPDF()
}

func main() {
	page = latex.Page{
		Name:      "Elmer Fudd",
		Sections:  []latex.Section{latex.Section{Title: "Hello", Body: "Hi there world."}},
		Dimension: latex.Dimension{20, 20, 20, 20},
	}

	http.HandleFunc("/create/", createHandler)
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/download/", downloadHandler)
	http.ListenAndServe(":80", nil)
}

package main

import (
	"encoding/gob"
	"os/exec"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/eimearc/latex"
)

var page latex.CV
var filePath string = "/root/go/src/github.com/eimearc/latex.txt"

func upload() latex.CV {
	fmt.Println("Upload file.")

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		log.Fatal("Error opening file", filePath, err)
	}

	var result latex.CV
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
	convertToPNG("tmp/tmp.pdf")
}

func viewPDF() {
	fmt.Println("View PDF.")
}

func convertToPNG(fileName string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Fatal("File", fileName, "does not exist\n", err)
	}
	pngName := strings.TrimSuffix(fileName, ".pdf") + ".png"
	cmd := exec.Command("convert", fileName, pngName)
	err := cmd.Run()
	if err != nil {
		log.Fatal("Cannot convert", fileName, "to", pngName, "\n", err)
	}
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	viewPDF()
	fmt.Fprint(w, "<!DOCTYPE html><html><body><p>Hello</p><img src=\"/tmp/tmp.png\" width=\"100%\" height=\"100%\"/></body></html>")
}

func gobRegister() {
	gob.Register(latex.Contact{})
	gob.Register(latex.Education{})
	gob.Register(latex.Experience{})
	gob.Register(latex.PersonalDetail{})
	gob.Register(latex.Skill{})
}

func main() {
	gobRegister()

	page = latex.CV{
		Sections:  []latex.Section{
			latex.PersonalDetail{"Elmer Fudd"},
			latex.Education{"Acme University"}},
		Dimension: latex.Dimension{20, 20, 20, 20},
	}

	http.HandleFunc("/create/", createHandler)
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/download/", downloadHandler)
	http.HandleFunc("/view/", viewHandler)
	http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("tmp"))))
	http.ListenAndServe(":80", nil)
}

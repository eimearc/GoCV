package main

import(
	"io/ioutil"
	"fmt"
	"net/http"
	"html/template"
	"github.com/eimearc/latex"
)

var page latex.Page

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	latex.SaveFile(page)
	fmt.Println("Saved file.")
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	p := latex.LoadFile()
	fmt.Println("Loaded file.")
	fmt.Fprintf(w, fmt.Sprintf("%#v\n",p))
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, latex.WriteFile("Latex function working"))
}

func main() {
	page = latex.Page{
		Name: "Elmer Fudd",
		Sections: []latex.Section{latex.Section{Title: "Hello", Body:"Hi there world."}},
		Dimension: latex.Dimension{ 20, 20, 20, 20},
	}
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/load/", loadHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

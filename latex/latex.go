package latex

import(
	"fmt"
	"os"
	"log"
	"encoding/gob"
)

// Page
	// Name
	// Section[]
		// Title
		// Body
	// Dimensions
		// Left
		// Right
		// Top
		// Bottom

var filePath string = "/root/go/src/github.com/eimearc/latex.txt" 

type Page struct {
	Name string
	Sections []Section
	Dimension Dimension
}

type Section struct {
	Title string
	Body string
}

type Dimension struct {
	Left, Right, Top, Bottom int
}

func WriteFile(w string) string  {
	return w
}

func SaveFile(page Page) {
	fmt.Println("Saving file.")
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

func LoadFile() Page {
	fmt.Println("Loading file.")
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil{
		log.Fatal("Error loading file", filePath, err)
	}

	var result Page
	enc := gob.NewDecoder(f)
	err = enc.Decode(&result)
	if err != nil {
		log.Fatal("Error decoding page from", f, err)
	}

	fmt.Println("File has been successfuly gob decoded.")

	return result
}

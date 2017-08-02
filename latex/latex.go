package latex

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

const filePath string = "/root/go/src/github.com/eimearc/server/tmp/"

type Page struct {
	Name      string
	Sections  []Section
	Dimension Dimension
}

type CV struct {
	Dimension Dimension
	Sections  []Section
}

type Dimension struct {
	Left, Right, Top, Bottom int
}

func CreatePDF() {
	cv := CV{
		Dimension: Dimension{20, 20, 20, 20},
		Sections: []Section{
			PersonalDetail{"Bugs Bunny"},
			Education{"Acme University"},
		},
	}

	fmt.Println("Create PDF.")
	cv.createLatex()
	cv.compileLatex()
}

func (cv CV) createLatex() {
	fmt.Println("Create Latex.")
	// Write to file tmp/tmp.tex
	fileName := filePath + "tmp.tex"
	code := "\\documentclass[a4paper,12pt]{article}\n" +
		"\\usepackage{array}\n" +
		"\\usepackage{color}\n" +
		"\\usepackage{geometry}\n" +
		"\\usepackage{lipsum}\n" +
		"\\usepackage{multirow}\n" +
		"\\usepackage{tabu}\n" +
		"\\usepackage[table]{xcolor}\n" +
		"\\begin{document}\n" +
		"Hello world.\n"
	/*
		for _, s := range cv.Sections {
			code += s.GetLatex() + "\n"
		}
	*/
	t := Table{
		[]Row{
			Row{
				[]Cell{
					Cell{
						[]Word{
							Word{
								"Hello",
								Black}},
						Gray},
					Cell{
						[]Word{
							Word{
								"World",
								Black}},
						Gray}},
				Gray},
			Row{
				[]Cell{
					Cell{
						[]Word{
							Word{
								"Golang",
								Black}},
						White},
					Cell{
						[]Word{
							Word{
								"Gopher",
								Black}},
						White}},
				White},
		}
	}
	fmt.Println(t.getLatex())
	code += t.getLatex()
	code += "\\end{document}\n"
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		log.Fatal("Error creating Latex.\n", err)
	}
	_, err = f.WriteString(code)
	if err != nil {
		log.Fatal("Error writing Latex.\n", err)
	}
}

func (cv CV) compileLatex() {
	fmt.Println("Compile Latex.")
	// write to file tmp/tmp.pdf
	fileName := filePath + "tmp.tex"
	cmd := exec.Command("pdflatex", "-output-directory", "tmp", fileName)
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error compiling Latex\n", err)
	}
}

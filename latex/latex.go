package latex

import (
	"fmt"
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

const filePath string = "tmp/"

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
}

func (cv CV) compileLatex() {
	fmt.Println("Compile Latex.")
	// write to file tmp/tmp.pdf
}

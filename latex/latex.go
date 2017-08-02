package latex

import(
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

type Page struct {
	Name string
	Sections []Section
	Dimension Dimension
}

type CV struct {
	Dimension Dimension
	Sections []Section
}

type Dimension struct {
	Left, Right, Top, Bottom int
}

func CreatePDF() {
	fmt.Println("Create PDF.")
	createLatex()
	compileLatex()
}

func createLatex() {
	fmt.Println("Create Latex.")
}

func compileLatex() {
	fmt.Println("Compile Latex.")
}

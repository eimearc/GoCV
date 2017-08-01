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

type Section struct {
	Title string
	Body string
}

type Dimension struct {
	Left, Right, Top, Bottom int
}

func CreatePDF() {
	fmt.Println("Create PDF.")
}

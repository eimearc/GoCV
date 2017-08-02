package latex

import (
	"fmt"
)

type Format interface {
	getLatex() string
}

type Color int

const (
	Black Color = iota
	Gray
	Blue
	Red
	White
)

type Font int

const (
	Arial Font = iota
	Helvetica
)

type Emphasis struct {
	Italic, Bold, SmallCaps bool
}

type Word struct {
	Content  string
	Color    Color
	//Font     Font
	//Emphasis Emphasis
}

type Cell struct {
	Words []Word
	Color Color
	//Multicolumn int
}

type Row struct {
	Cells []Cell
	Color Color
}

type Table struct {
	Rows []Row
}

func (c Color) getLatex() string {
	switch c {
	case Black:
		return "black"
	case Gray:
		return "gray"
	case Blue:
		return "blue"
	case Red:
		return "red"
	case White:
		return "white"
	}
	return ""
}

func (w Word) getLatex() string {
	l := fmt.Sprintf("\\color{%s} %s", w.Color.getLatex(), w.Content)
	return l
}

func (c Cell) getLatex() string {
	var l string
	for _, w := range c.Words {
		l += w.getLatex()
	}
	l = fmt.Sprintf("\\cellcolor{%s} %s", c.Color.getLatex(), l)
	return l
}

func (r Row) getLatex() string {
	var l string
	for i, c := range r.Cells {
		l += c.getLatex()
		if i < len(r.Cells)-1 {
			l += " & "
		}
	}
	l = fmt.Sprintf("\\rowcolor{%s} %s \\\\\n", r.Color.getLatex(), l)
	return l
}

func (t Table) getLatex() string {
	var l string
	l += "\\begin{tabu} to \\textwidth {X[1,l]X[1,l]}\n"
	for _, c := range t.Rows {
		l += c.getLatex()
	}
	l += "\\end{tabu}\n"
	return l
}

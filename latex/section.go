package latex

import (
	"fmt"
)

type Section interface {
	GetName() string
	GetLatex() string
}

type Contact struct {
	Name string
}

type Education struct {
	Name string
}

type Experience struct {
	Name string
}

type PersonalDetail struct {
	Name string
}

type Skill struct {
	Name string
}

func (c Contact) GetName() string {
	return c.Name
}

func (e Education) GetName() string {
	return e.Name
}

func (e Experience) GetName() string {
	return e.Name
}

func (p PersonalDetail) GetName() string {
	return p.Name
}

func (s Skill) GetName() string {
	return s.Name
}

func (p PersonalDetail) GetLatex() string {
	name := p.Name
	return fmt.Sprintf("\\centerline{\\Huge{%s}}", name)
}

func (e Education) GetLatex() string {
	name := e.Name
	return fmt.Sprintf("\\centerline{\\Huge{%s}}", name)
}

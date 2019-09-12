package main

import (
	"fmt"
	"strconv"
)

type Planet struct {
	Name string
	Mass float64
}

func (p Planet) String() string {
	return fmt.Sprintf("%s (%.3f)", p.Name, p.Mass)
}

func (p Planet) GoString() string {
	return fmt.Sprintf(`main.Planet{Name:%s, Mass:%.3f}`, strconv.Quote(p.Name), p.Mass)
}

func (p Planet) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case s.Flag('#'):
			s.Write([]byte(p.GoString()))
		case s.Flag('+'):
			s.Write([]byte(fmt.Sprintf(`{"Name":%s,"Mass":%.3f}`, strconv.Quote(p.Name), p.Mass)))
		default:
			s.Write([]byte(p.String()))
		}
	case 's':
		s.Write([]byte(p.String()))
	}
}

var planets = []Planet{
	{Name: "Mercury", Mass: 0.055},
	{Name: "Venus", Mass: 0.815},
	{Name: "Earth", Mass: 1.0},
	{Name: "Mars", Mass: 0.107},
}

func main() {
	fmt.Printf("%v\n", planets)
	fmt.Printf("%#v\n", planets)
	fmt.Printf("%+v\n", planets)
}

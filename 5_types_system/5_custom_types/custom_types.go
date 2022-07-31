package main

import "fmt"

type Score float64

func (s Score) between(start, end float64) bool {
	return float64(s) >= start && float64(s) <= end
}

/*func scoreToConcept(s Score) string {
	if s.between(9.0, 10.0) {
		return "A"
	} else if s.between(7.0, 8.99) {
		return "B"
	} else if s.between(5.0, 6.99) {
		return "C"
	} else if s.between(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}*/

func scoreToConcept(s Score) string {
	switch {
	case s.between(9.0, 10.0):
		return "A"
	case s.between(7.0, 8.99):
		return "B"
	case s.between(5.0, 6.99):
		return "C"
	case s.between(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(scoreToConcept(9.8))
	fmt.Println(scoreToConcept(6.9))
	fmt.Println(scoreToConcept(2.1))
}

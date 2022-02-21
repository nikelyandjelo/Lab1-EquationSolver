package main

import (
 
	"math"

)
func main(){

}

func solutionEquation (param []float64) []float64 {
	var roots []float64

	var a = param[0]
	var b = param[1]
	var c = param[2]

	var Discriminant = (b*b) - (4*a*c)
	
	if Discriminant == 0 {
		var oneRoot = -(b) / (2 * a)
		roots = append(roots, oneRoot)
	} else if Discriminant > 0 {
		var firstRoot = (-b + math.Sqrt(Discriminant)) / (2*a)
		var secondRoot = (-b - math.Sqrt(Discriminant)) / (2*a)
		roots = append(roots, firstRoot, secondRoot)
	}

	return roots
}


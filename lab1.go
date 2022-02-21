package main

import (
	"fmt"
  "os"
	"math"
	"bufio"
	"strconv"
	"strings"
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

func interactive () []float64{

	var param []float64
	var arguments = []string{"a", "b", "c"} 
 
	var readInput = bufio.NewReader(os.Stdin)
 
	for  i := 0; i != 3; i++{
	 fmt.Printf("%s: ", arguments[i])
	 var parameters, flaw  = readInput.ReadString('\n')
	 if flaw != nil {
		fmt.Println("Произошла ошибка при вводе данных. Повторите попытку", flaw)
		continue
	 }
 
	 parameters = strings.TrimSuffix(parameters, "\r\n")
	 parameters1, flaw := strconv.ParseFloat(parameters, 64)
	 if flaw != nil {
		fmt.Printf("Ошибка. Введите правильные значения пж, получили %s \n", parameters)
		continue
	 }
 
	 if parameters1 == 0 && len(param) == 0 {
		fmt.Print("Ошибка. а не может быть 0\n")
		continue
	 }
	 param = append(param, parameters1)
	}
 
	return param
 
 }
 
 

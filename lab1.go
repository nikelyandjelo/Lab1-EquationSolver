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
	var inputParameters []float64
 	var fileName string 
 	var result []float64

 	if len(os.Args) > 1 {
  fileName = os.Args[1]
 	}

 	if fileName != "" {
  inputParameters = noInteractive(fileName)
 	} else {
  inputParameters = interactive()
	}

 	result = solutionEquation(inputParameters)

 	fmt.Printf("Уравнение: (%.2f) x^2 + (%.2f) x + (%.2f) = 0\n", 
 	inputParameters[0], inputParameters[1], inputParameters[2])

 	if len(result) == 1 {
 	 fmt.Printf("x1 = %f\n", result[0])
 	} else if len(result) == 2 {
  	fmt.Printf("x1 = %f\nx2 = %f\n", result[0], result[1])
	}

	 fmt.Printf("Имеем %d корня", len(result))

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

 func noInteractive(fileName string) []float64 {
	var value []float64
 
	file, err := os.ReadFile(fileName) 
	if err != nil {
	 fmt.Println("Произошла ошибка, открывая файл", err)
	 os.Exit(1)
	}
 
	str :=strings.Fields(string(file))
 
	if len(str) != 3 {
	 fmt.Fprint(os.Stderr, "Ошибка: недопустимый формат файла, жду 3 агрумента \n")
	 os.Exit(1)
	}
 
	for index, values := range str {
	 params, err := strconv.ParseFloat(values, 64)
	 if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %s\n", err)
		os.Exit(1)
	 }
	 if index == 0 && params == 0 {
		fmt.Print("Ошибка . Не может быть  0\n")
		os.Exit(1)
	 }
	 value = append(value,params)
	}
	return value
 }
 
 

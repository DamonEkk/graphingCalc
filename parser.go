package main

import "fmt"
import "math"
import "strconv"
// Parser for a graphing calculator

// Operations enum

const (
	ADD = iota
	SUBTRACT
	MULTIPLY
	DIVIDE
	POWER
	MODULO
)

// Special function table
type specialFunc struct{
	Inputs int
	Special func(float64) float64
}

var specialFunctions = map[string] specialFunc{
	// Trig
	"sin": {Inputs: 1, Special: math.Sin},
	"cos": {Inputs: 1, Special: math.Cos},
	"tan": {Inputs: 1, Special: math.Tan},
	"asin": {Inputs: 1, Special: math.Asin},
	"acos": {Inputs: 1, Special: math.Acos},
	"atan": {Inputs: 1, Special: math.Atan},

	// Hyperblic
	"sinh": {Inputs: 1, Special: math.Sinh},
	"cosh": {Inputs: 1, Special: math.Cosh},
	"tanh": {Inputs: 1, Special: math.Tanh},
	"asinh": {Inputs: 1, Special: math.Asinh},
	"acosh": {Inputs: 1, Special: math.Acosh},
	"atanh": {Inputs: 1, Special: math.Atanh},

	// Exponentials NOTE pow handled as operator
	"e": {Inputs: 2, Special: math.Exp},
	"log": {Inputs: 1, Special: math.Log},
	"log10": {Inputs: 1, Special: math.Log10},
	"log2": {Inputs: 1, Special: math.Log2},
	"sqrt": {Inputs: 1, Special: math.Sqrt},
	"cbrt": {Inputs: 1, Special: math.Cbrt},
	"abs": {Inputs: 1, Special: math.Abs}, 
}

// Find expression

type Multiply struct{}

type Division struct{}

type Addition struct{}

type Subtract struct{}

type Power struct{}

type modulo struct{}

// Solve any expression with solve
type Calculator interface {
	Solve(a, b string) string
}

// Takes values and multiplies
func (expr Multiply) Solve(a string, b string) string{
	// Num convert
	numA, errA := strconv.Atoi(a)
	numB, errB := strconv.Atoi(b)

	// Simple a and b are both nums
	if (errA == nil && errB == nil){
		return string(numA * numB)
	}



	return ""
}

// Takes values and divides
func (expr Division) Solve(a string, b string){
	return
}

// Takes values and adds
func (exp Addition) Solve(a string, b string){
	return
}

// Takes values and subtrcts
func (exp Subtract) Solve(a string, b string){
	return
}

// Used to find substrings with parentheses ()
func ParenthesesString(expression string, startIndex int)string{
	var subExpression string = string(expression[startIndex])
	paramCount := 0

	for i := startIndex; i < len(expression); i++{
		if (expression[i] == ')'){
			subExpression += string(expression[i])
			
			if (paramCount == 0){
				break
			} else{
				paramCount--
			}
		}
		
		// If more parantheses are opened we need to make sure the for loop doesnt close prematurely
		if (expression[i] == '('){
			paramCount++
		}

		subExpression += string(expression[i])
	}
	
	// Create a _test later to check conditions
	if (paramCount != 0){
		fmt.Printf("Incomplete expression = %s || Missing closing parantheses\n", expression)
	}

	return subExpression
}

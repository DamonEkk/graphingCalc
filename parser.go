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

type Multiply struct{
	Left string
	Right string
}

type Division struct{
	Left string
	Right string
}

type Addition struct{
	Left string
	Right string
}

type Subtract struct{
	Left string
	Right string
}

type Power struct{
	Left string
	Right string
}

type Modulo struct{
	Left string
	Right string
}

// Solve any expression with solve
type Calculator interface {
	Solve(a, b string) string
}


// Takes values and multiplies
func (expr Multiply) Solve() string{
	// Num convert
	numA, errA := strconv.Atoi(expr.Left)
	numB, errB := strconv.Atoi(expr.Left)

	// Simple a and b are both nums
	if (errA == nil && errB == nil){
		return strconv.Itoa(numA * numB)
	}

	return ""
}

// Takes values and divides
func (expr Division) Solve() string{
	// Num convert
	numA, errA := strconv.Atoi(expr.Left)
	numB, errB := strconv.Atoi(expr.Left)

	// Simple a and b are both nums
	if (errA == nil && errB == nil){
		return strconv.Itoa(numA / numB)
	}

	return ""
}

// Takes values and adds
func (expr Addition) Solve() string{
	// Num convert
	numA, errA := strconv.Atoi(expr.Left)
	numB, errB := strconv.Atoi(expr.Left)

	// Simple a and b are both nums
	if (errA == nil && errB == nil){
		return strconv.Itoa(numA + numB)
	}

	return ""
}

// Takes values and subtrcts
func (expr Subtract) Solve() string{
	// Num convert
	numA, errA := strconv.Atoi(expr.Left)
	numB, errB := strconv.Atoi(expr.Left)

	// Simple a and b are both nums
	if (errA == nil && errB == nil){
		return strconv.Itoa(numA - numB)
	}

	return ""
}

func (expr Power) Solve() string{
	// Num convert
	numA, errA := strconv.Atoi(expr.Left)
	numB, errB := strconv.Atoi(expr.Left)

	// Simple a and b are both nums
	if (errA == nil && errB == nil){
		return strconv.Itoa(numA ^ numB)
	}

	return ""
}

func (expr Modulo) Solve() string{
	// Num convert
	numA, errA := strconv.Atoi(expr.Left)
	numB, errB := strconv.Atoi(expr.Left)

	// Simple a and b are both nums
	if (errA == nil && errB == nil){
		return strconv.Itoa(numA % numB)
	}


	return ""
}


// Main call function
func Parser(expression string)string{
	expression = Checker(expression)
	var subExpression string = ""
	var newExpression string = ""
	
	
	// Looks for nested () and recursively parses through them.
	for i := 0; i < len(expression);i++{
		if (expression[i] == '('){
			subExpression = ParenthesesString(expression, i) // Get substring
			subRange := len(subExpression) // Get size of substring to subtract from main.
			subExpression = Parser(subExpression) // Solve substring
			newExpression += subExpression // Add solved substring to current string
			i += subRange - 1 // Skip to after substring
		}
	}

	// BOMDAS
	
	// POWER / LOG / SQR / MODULO
	for i := 0; i < len(newExpression); i++{
	
		switch newExpression[i]{
		case '^':
			var pow Power
			l,r := GetSurrounding(i, newExpression)
			pow.Left = l
			pow.Right = r
			pow.Solve()

		case '%':
			var mod Modulo
			l,r := GetSurrounding(i, newExpression)
			mod.Left = l
			mod.Right = r
			mod.Solve()

		// sqr cases, should implement cbr cases also but ill do it later.
		case 's':
			if (i+2 <= len(newExpression)){
				if (newExpression[i:3] == "qr("){
					sqrParams := ParenthesesString(newExpression, i+3) // Need to check if there are any variables inside before converting to int.
					value, err := strconv.Atoi(sqrParams)

					if (err == nil ){
						return strconv.FormatFloat(math.Sqrt(float64(value)), 'f', 2, 64)
					} else{
						fmt.Printf("SQRT Variables found: %s\n", value)
						// Handle variables 
					}
				}
			}
		
		// Log cases. 
		case 'l':
			if (len(expression) < i+5){
				if (expression[i:3] == "og("){
					logParam := ParenthesesString(newExpression, i+3)
					value, err := strconv.Atoi(logParam)

					if (err == nil){
						return strconv.FormatFloat(math.Log(float64(value)), 'f', 2, 64)
					} else{ // Handle variable output (not just nums found)
						fmt.Printf("LOG Variables found: %s\n", value)
					}

				} else if (expression[i:4] == "og2("){
					logParam := ParenthesesString(newExpression, i+4)
					value, err := strconv.Atoi(logParam)

					if (err == nil){
						return strconv.FormatFloat(math.Log(float64(value)), 'f', 2, 64)
					} else{ // Handle variable output (not just nums found)
						fmt.Printf("LOG2 Variables found: %s\n", value)
					}

				} else if (expression[i:5] == "og10("){
					logParam := ParenthesesString(newExpression, i+5)
					value, err := strconv.Atoi(logParam)

					if (err == nil){
						return strconv.FormatFloat(math.Log(float64(value)), 'f', 2, 64)
					} else{ // Handle variable output (not just nums found)
						fmt.Printf("LOG10 Variables found: %s\n", value)
					}
				}
			}
		}
	}

	// MULTIPLICATION / DIVIDISION
	for i := 0; i < len(newExpression); i++{
		switch newExpression[i]{
		case '*':
			var mod Multiply
			l,r := GetSurrounding(i, newExpression)
			mod.Left = l
			mod.Right = r
			mod.Solve()
			
		case '/':
			var mod Division
			l,r := GetSurrounding(i, newExpression)
			mod.Left = l
			mod.Right = r
			mod.Solve()	
		}
		
	}

	// ADDITON / SUBTRACTION
	for i := 0; i < len(newExpression); i++{	

		switch newExpression[i]{
		case '+':
			var mod Addition
			l,r := GetSurrounding(i, newExpression)
			mod.Left = l
			mod.Right = r
			mod.Solve()

		case '-':
			var mod Subtract 
			l,r := GetSurrounding(i, newExpression)
			mod.Left = l
			mod.Right = r
			mod.Solve()

		}	
		
	}

	
	return ""	
}

// Id like to somehow pass in any of the operator structs and set them directly, but couldnt figure it out. Here is my hack.
func GetSurrounding(modIndex int, expr string)(string,string){
	a := ""
	aReverse := ""
	b := ""
	// Left side A

	index := modIndex - 1
	for {
		if (index <= 0){
			break
		}

		if (expr[index] != '/' && expr[index] != '*' && 
			expr[index] != '+' && expr[index] != '-' &&
			expr[index] != '(' && expr[index] != ')'){
			aReverse += string(expr[index])

		} else{
			break
		}

		index--
	}

	for i := len(aReverse); i > 0; i--{
		a += string(aReverse[i])
	}

	// Right side B

	index = modIndex + 1
	for {
		if (index >= len(expr)){
			break
		}

		if (expr[index] != '/' && expr[index] != '*' && 
			expr[index] != '+' && expr[index] != '-' &&
			expr[index] != '(' && expr[index] != ')'){
			b += string(expr[index])

		} else{
			break
		}

		index++
	}

	return a, b
}


// Used to translate short form into computer readable. Such as 3(4*3) -> 3*(4*3)
func Checker(expression string)string{
	var newExpression string = ""

	// For every '(' found check to make sure its not the first index and add a *
	for i := 0; i < len(expression); i++{
		if (expression[i] == '(' && i != 0 && 
		expression[i - 1] != '/' && expression[i - 1] != '*' && 
		expression[i - 1] != '+' && expression[i - 1] != '-'){
			newExpression += string("*")
		}
	
		newExpression += string(expression[i])
	}

	return newExpression
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

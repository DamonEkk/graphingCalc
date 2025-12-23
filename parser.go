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
	numB, errB := strconv.Atoi(expr.Right)

	// Simple a and b are both nums
	if (errA == nil && errB == nil){ // Num * Num
		return strconv.Itoa(numA * numB)
	} else if (errA == nil && errB != nil){ // Num * Expression
		return MultAdder(numA, expr.Right)	
	} else if (errA != nil && errB == nil){ // Expression * Num
		return MultAdder(numB, expr.Left)
	} else { // Expression * Expression

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

// Used to mult one sided variable to a num.
func MultAdder(num int, expression string)string{
	expressionNum := ""
	variableName := ""
	specialInside := ""
	result := ""
	stopIndex := 0
	flag := 0

	// Check each char for num values that will be multiplied.
	for i := 0; i < len(expression); i++{
		_, err := strconv.Atoi(string(expression[i]))
		if (err != nil){
			stopIndex = i
			break
		}
		expressionNum += string(expression[i])
	}

	// Used to check special function, second loop used for readability
	for i := stopIndex; i < len(expression); i++{

		if (flag == 1){
			
			if (expression[i] == ')'){
				flag = 0
				continue
			}

			specialInside += string(expression[i])
			continue
		}

		if (expression[i] != '('){
			variableName += string(expression[i])
		}  else {
			flag = 1 
			continue
		}
	}

	special, ok := specialFunctions[variableName]

	if (!ok && len(expressionNum) == 0){
		result = strconv.Itoa(num)
		result += variableName
	} else if (!ok && len(expressionNum) > 0){
		eNum, _ := strconv.Atoi(expressionNum)
		result = strconv.Itoa(eNum * num)
		result += variableName
	} else{
		result = SpecialFunctionHandler(special, specialInside)

		// 2 cases here. 1 the spcialFunction inside is returned back due to having a variable inside. 2. We get a num back and can finish.

	}

	return result
}

func SpecialFunctionHandler(special specialFunc, params string)string{
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
			newExpression = ReplaceExpression(i, newExpression, pow.Solve(), pow.Left, pow.Right) // Solve and replace in expression string


		case '%':
			var mod Modulo
			l,r := GetSurrounding(i, newExpression)
			mod.Left = l
			mod.Right = r
			newExpression = ReplaceExpression(i, newExpression, mod.Solve(), mod.Left, mod.Right) // Solve and replace in expression string


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
			// Case Log()
			if (len(expression) < i+5){
				if (expression[i:3] == "og("){
					logParam := ParenthesesString(newExpression, i+3)
					value, err := strconv.Atoi(logParam)

					if (err == nil){
						return strconv.FormatFloat(math.Log(float64(value)), 'f', 2, 64)
					} else{ // Handle variable output (not just nums found)
						fmt.Printf("LOG Variables found: %s\n", value)
					}

				// Case Log2()
				} else if (expression[i:4] == "og2("){
					logParam := ParenthesesString(newExpression, i+4)
					value, err := strconv.Atoi(logParam)

					if (err == nil){
						return strconv.FormatFloat(math.Log(float64(value)), 'f', 2, 64)
					} else{ // Handle variable output (not just nums found)
						fmt.Printf("LOG2 Variables found: %s\n", value)
					}

				// Case Log10()
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
			var mult Multiply
			l,r := GetSurrounding(i, newExpression) // Get the left and right variables and consts on operator
			mult.Left = l
			mult.Right = r
			newExpression = ReplaceExpression(i, newExpression, mult.Solve(), mult.Left, mult.Right) // Solve and replace in expression string
			
		case '/':
			var div Division
			l,r := GetSurrounding(i, newExpression)
			div.Left = l
			div.Right = r
			newExpression = ReplaceExpression(i, newExpression, div.Solve(), div.Left, div.Right)
		}
		
	}

	// ADDITON / SUBTRACTION
	for i := 0; i < len(newExpression); i++{	
		
		// Same logic as above with different operators
		switch newExpression[i]{
		case '+':
			var add Addition
			l,r := GetSurrounding(i, newExpression)
			add.Left = l
			add.Right = r
			newExpression = ReplaceExpression(i, newExpression, add.Solve(), add.Left, add.Right)

		case '-':
			var sub Subtract 
			l,r := GetSurrounding(i, newExpression)
			sub.Left = l
			sub.Right = r
			newExpression = ReplaceExpression(i, newExpression, sub.Solve(), sub.Left, sub.Right)

		}	
		
	}
	return newExpression	
}

// Replace expression allows us to delete the original expression that was solved and replace it with the solved.
func ReplaceExpression(index int, originalExpression string, solvedExpression string, left string, right string)string{
	var replacedExpresssion string = ""

	// Parse through originalExpression till we get to the range of the solved operation. Replace with solvedExpression
	for i := 0; i < len(originalExpression); i++{
		if (i == index) {
			replacedExpresssion += solvedExpression
		} else if (i >= len(left) && i <= len(right)){
			continue
		}

		replacedExpresssion += string(originalExpression[i])
	}

	return replacedExpresssion
}

// Id like to somehow pass in any of the operator structs and set them directly, but couldnt figure it out. Here is my hack.
func GetSurrounding(modIndex int, expr string)(string,string){
	a := ""
	aReverse := ""
	b := ""
	// Left side A

	index := modIndex - 1
	for {
		// If index == 0 we have found all possible variables.
		if (index < 0){
			break
		}

		// Check left side for operator. End of variable essentially
		if (expr[index] != '/' && expr[index] != '*' && 
			expr[index] != '+' && expr[index] != '-' &&
			expr[index] != '(' && expr[index] != ')'){
			aReverse += string(expr[index])

		} else{
			break
		}

		index--
	}

	// Left is added backwards so small loop to reverse it.
	for i := len(aReverse); i > 0; i--{
		a += string(aReverse[i])
	}

	// Right side B

	index = modIndex + 1
	for {
		// Find end of expression.
		if (index > len(expr)){
			break
		}

		// Check for operators as a boundary for variable.
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

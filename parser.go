package main

// Parser for a graphing calculator

// Find expression

type Multiply struct{}

type Division struct{}

type Addition struct{}

type Subtract struct{}

// Solve any expression with solve
type Calculator interface {
	Solve(a, b string) string
}

// Takes values and multiplies
func (expr Multiply) Solve(a string, b string){
	return
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

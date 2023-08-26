// https://www.udemy.com/course/go-language/learn/lecture/35216168#notes
//
package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
)

var pl = fmt.Println

func main() {
	pl("5 + 4 = ", 5+4)
	pl("5 - 4 = ", 5-4)
	pl("5 * 4 = ", 5*4)
	pl("5 / 4 = ", 5/4)
	pl("6 % 4 = ", 6%4) // Get remainder after division


	// Shorthand increment
	// Instead of mInt = mInt + 1 (mInt += 1)
	// -= *= /=
	mInt := 1
	mInt += 1  
	// Also increment or decrement with ++ and --
	mInt++
	pl("mInt = ", mInt)


	// Float precision increases with the size of your values
	pl("Float Precision =", 0.11111111111111111111+
		0.11111111111111111111)


	// Create a random value between 0 and 50
	// Get a seed value for our random number generator based on
	// seconds since 1/1/70 to make our random number more random	
	pl("---------------- Random numbers -------------")
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	// Between 0 and up to, but not including
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)

	// There are many math functions
	pl("Abs(-10) = ", math.Abs(-10))
	pl("Pow(4,2) = ", math.Pow(4,2))
	pl("Sqrt(16) = ", math.Sqrt(16))
	pl("Cbrt(8) = ", math.Cbrt(8))
	pl("Ceil(4.4) = ", math.Ceil(4.4))
	pl("Floor(4.4) = ", math.Floor(4.4))
	pl("Roumd(4.4) = ", math.Round(4.4))
	pl("Log2(8) = ", math.Log2(8))
	pl("Log10(100) = ", math.Log10(100))
	pl("Log(7.389) = ", math.Log(7.389))
	pl("Max(5,4) = ", math.Max(5,4))
	pl("Min(5,4) = ", math.Min(5,4))

	// See: https://pkg.go.dev/math  and  
	// https://golangdocs.com/math-package-in-golang

	// Convert radians to degrees and back
	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	// Convert 1.5708 radians to degrees
	d90 := r90 * (180 /math.Pi)
	fmt.Printf("%f radians = %f degrees\n", r90, d90)

	pl("Sin(90) = ", math.Sin(r90))
	// There are also functions for Cos, Tan, Acos, Asin
	// Atan, Asinh, Acosh, Atanh, Atan2, Cosh, Sinh, Sincos
	// Htpot
}

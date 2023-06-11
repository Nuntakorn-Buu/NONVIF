// `package main` is the declaration of the package name.
// In Go, the `main` package is a special package that is used to build executable programs.
// It must contain a `main` function as the entry point of the program.
package main

// The `import` statement is used to import packages that are necessary for the program to run.
// In this case, the program is importing the following packages:
import (
	"crypto/rand" // Used to generate mathematical random values, it helps to generate highly secure random numbers for use in key generation, encryption, random number generation, URL generation, and more.
	"fmt"         // Tools for manipulating the screen display. and reading data from the user You can use the Printf() function.
	"log"         // Use the message log (log) and don't forget the message. (error reporting)
	"math/big"    // Used to work with large numbers, which are arithmetic operations that can handle numbers whose length is greater than the ones stored in the underlying data type, such as large integers.
)

// Random 6 digits number
// The function generates a random six-digit integer.
func sixDigits() int64 {
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatal(err)
	}
	return n.Int64()
}

// Check Error Handling
// The function prints an error message if the error parameter.
func FetchError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

package piscine

import (
	"fmt"
	"os"
	"strconv"
)

func Fprime() {
	// must have exactly one argument, else output nothing
	if len(os.Args) != 2 {
		return
	}

	// convert argument to integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 { // if invalid or <=1 -> print nothing
		return
	}

	d := 2        // start checking factors from 2
	first := true // used to avoid leading '*'

	// factorize while divisor squared <= number
	for d*d <= n {
		// if divisible, print factor and divide n
		for n%d == 0 {
			if !first { // add "*" only after first number
				fmt.Print("*")
			}
			fmt.Print(d) // print the factor
			first = false
			n /= d // reduce number
		}
		d++ // move to next possible factor
	}

	// if n is still >1, it is prime itself
	if n > 1 {
		if !first {
			fmt.Print("*")
		}
		fmt.Print(n)
	}

	fmt.Println()
}

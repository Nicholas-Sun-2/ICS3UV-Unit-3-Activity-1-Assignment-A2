/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-18
 * @fileoverview This program calculates the length, width and height of a cube that has a volume of 1000 mm^3.
 */

package main

import "fmt"

func main() {
	// Declaring variables
	var volume int = 1000

	var side_length int = 10 // We don't have cube root yet.

	// Printing
	fmt.Println("A cube with volume of " + fmt.Sprint(volume) + " mm^3 has a length, width, and height of " + fmt.Sprint(side_length) + " mm.")
}

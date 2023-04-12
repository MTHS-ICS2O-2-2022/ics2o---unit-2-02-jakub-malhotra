// Created by: Jakub Malhotra
// Created on: April 2023
//
// This program displays, "area and perimeter of a rectangle"
package main
import "fmt"
func main() {
	// This function displays, "area and perimeter of a rectangle"
	// the Go standard is to indent using "tabs" NOT spaces!
var length int
	fmt.Print("Enter the length: ")
	fmt.Scan(&length)
	var width int
	fmt.Print("Enter the width: ")
	fmt.Scan(&width)
area := length * width
perimeter := 2 * length + 2 * width
fmt.Println("The area is:", area, "cm²")
fmt.Println("The perimeter is:", perimeter, "cm")
	fmt.Println("\nDone.")
}

package data_structures

import (
	"fmt"
)

func MultiDimensional() {

	fmt.Println("===== MultiDimensional Arrays === \n ")

	const description string = `
	Multi Dimensional Array declaration in GO

	Syntax : var arr [2][2]int
	Description : this declares an empty array of 2 rows, 2 column capacity
	output -> arr = [[0, 0], [0, 0]]`

	fmt.Println(description)

	var arr [2][2]int
	fmt.Println("\n	Example : \n\n	Declaring Array \n	var arr [2][2]int \n ")
	fmt.Println("	Output : ", arr, "\n ")

	fmt.Println("\n === Multi Dimensional Array Initialisation === \n ")

	const description1 string = `
	Multi Dimensional Array Initialisation in GO

	Syntax : var arr1 = [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 0},
	} 
	Description : this initialises an array with 2 rows, 5 columns
	output -> arr = [[1, 2, 3, 4, 5], [6, 7, 8, 9, 0]]`

	fmt.Println(description1)

	var arr1 = [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 0},
	}
	fmt.Println("\n	Example : \n\n	Declaring Array \n	var arr [2][5]int \n ")
	fmt.Println("	Output : ", arr1, "\n ")

}

func Append() {

	fmt.Println("===== Append Arrays === \n ")
	const appendDescription = `
	Arrays cannot be appended using any inbuilt function 
	As there no function in go to concat/append/merge array
	Because append creates new space with dynamic size which
	array cannot have instead slices can be appended.

	Alternative Way (Not recommended instead use slice):
	- Convert Array to slice
	- Append Slice to another
	- Convert Slice to Array
	`

	fmt.Println(appendDescription, "\n ")

}

func Copy() {

	fmt.Println("===== Copy Array === ")
	const copyDescription = `
	Syntax : soureArray = destinationArray
	Note: Copying array is immutable i.e, it does not affect original array if any modification
	is made to copy of the original array

	Example:

	`
	fmt.Println(copyDescription)

	var src = [...]string{"apple", "lemon", "banana", "grapes", "strawberry"}
	var dest = src

	fmt.Println("\n	Original Array : ", src)
	fmt.Println("	Copied   Array : ", dest)
	dest[1] = "mango"
	fmt.Println("\n	After modification to copied array : ")
	fmt.Println("\n	Original Array : ", src)
	fmt.Println("	Copied   Array : ", dest, " \n ")
}

func Assign() {

	fmt.Println("\n === ARRAY : Assign Value === \n ")

	const description string = `
	Assign value to array index

	Syntax : array[index] = value
	Description : this assign values to specific index
	
	Example:
	
	Before
		arr = [0, 0, 0, 0, 0]

	Assign 5 at index 2 in array arr
		arr[2] = 5

	After
		arr = [0, 0, 5, 0, 0]
	`
	fmt.Println(description, "\n ")
}

// InitArrayInGo - explains how to initialize the array in go
func Init() {

	fmt.Println("\n === ARRAY DECLARATION === \n ")

	const description string = `
	Array Declaration in GO

	Syntax : var arr [5]int
	Description : this declares an empty array of 5 elements capacity
	output -> arr = [0, 0, 0, 0, 0]`

	fmt.Println(description)

	var arr [5]int
	fmt.Println("\n	Example : \n\n	Declaring Array \n	var arr [5]int \n ")
	fmt.Println("	Output : ", arr, "\n ")

	fmt.Println("\n === ARRAY INTIALISATION === \n ")

	const description1 string = `
	Array Initialisation in GO

	Syntax : var arr1 = [5]int{1, 2, 3, 4, 5} 
	Description : this initialises an array with 5 elements
	output -> arr = [1, 2, 3, 4, 5]`

	fmt.Println(description1)

	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("\n	Example : \n\n	Declaring Array \n	var arr [5]int \n ")
	fmt.Println("	Output : ", arr1, "\n ")
}

package main

import "fmt"

const row int = 5

func buildRow(array *[]int) {
	aux			:= 0;
	coefficient := 1;
	xExponent	:= row;
	yExponent	:= 1;

	for xExponent > 0 {
		aux = (coefficient * xExponent) / yExponent
		(*array)[yExponent] = aux
		coefficient = aux
		xExponent--
		yExponent++
	}
}

func main() {
	array := make([]int, row + 1)
	array[0] = 1

	if row == 0 {
		fmt.Println(array)
	} else if row == 1 {
		array[1] = 1
		fmt.Println(array)
	} else if row > 1 {
		buildRow(&array)
		fmt.Println(array)
	}else{
		fmt.Println("Only positive integer number.")
	}
}

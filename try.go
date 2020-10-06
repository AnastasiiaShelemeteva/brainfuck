package main

import (
	"fmt"
	"strings"
)
func main() {
	
	memory := []int {0}
	position := 0
	inputVar := ("+++++++++++++++++++++++++++++++++++--.>>+++++++++++++++++++++++++++++++++++--.<+++++++++++++++++++++++++++++++++++--.<++++.")
	toArray := strings.Split(inputVar, "")	


		for i, element := range toArray {
			switch element {
			case "+": 
			memory[position] += 1
				
			case "-":
			memory[position] -= 1
				
			case ">":
				position += 1
				if len(memory) <= position {
					memory = append(memory, 0)
				}
			case "<":
				position -= 1
				
			case ".":
				fmt.Print(position, ": ", string(memory[position]), "\n")	
			/* 
			case "[":
				if memory[position] == 0 {
					countOpened := 0

				} */

			
							
			}
		i+=1
	} 
	fmt.Printf("%d: %v", position, memory)
}
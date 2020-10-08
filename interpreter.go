package main

import (
	"fmt"
/* 	"os" */
	//"io/ioutil"
	//"bufio"
	"strings"
	
)







func main() {
	prgPos := 0
	mem := []int{0}
	memPos := 0

	inputVar := ("++++++++++[>+++++++++>++++++++++>+++>+++++++++++>+++++++>++++++++++++<<<<<<-]>---.>+++++.+++..---.--------.>>-.<++.>>++.>+.<<<<+++.+. ")
	toArray := strings.Split(inputVar, "")	
	for prgPos < len(toArray) {
		if toArray[prgPos] == ">" {
			memPos += 1
			if len(mem) <= memPos {
				mem = append(mem, 0)
			}
			
		} else if toArray[prgPos] == "<" {
			memPos -= 1
			
		} else if toArray[prgPos] == "+" {
			mem[memPos] += 1
			
		} else if toArray[prgPos] == "-" {
			mem[memPos] -= 1
			
		} else if toArray[prgPos] == "." {
			fmt.Println("OUR ELEMENT IS ", string(mem[memPos]), "our memPos is", memPos, "symbol is ", toArray[prgPos])
		} else if toArray[prgPos] == "[" {
			
			if mem[memPos] == 0 {
				counterOpen := 0
				prgPos += 1
				
				for prgPos < len(toArray) {
					if toArray[prgPos] == "]" && counterOpen == 0 {
						
						break
					} else if  toArray[prgPos] == "[" {
						counterOpen += 1
						
					} else if toArray[prgPos] == "]" {
						
						counterOpen -= 1
					} 
					prgPos += 1
				} 
				
			} 
		} else if toArray[prgPos] == "]" {
			
			if mem[memPos] != 0 {
				for toArray[prgPos] != "[" {
					prgPos -= 1
					
			/* 	} else {
					maybe here should be something or not
				} */
			}
		} 		
		
	}
	prgPos+=1

	}

	fmt.Println("Here is an array: ", mem)


}
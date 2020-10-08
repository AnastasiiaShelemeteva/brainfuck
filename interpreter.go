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

	inputVar := ("+++++[>+>++>+++>++++>+++++<<<<<-]>-.>+++.+++..---.--.>>-.<++.>>++.>+.<<<<+++.+.")
	toArray := strings.Split(inputVar, "")	
	for prgPos < len(toArray) {
		if toArray[prgPos] == ">" {
			memPos += 1
			if len(mem) <= prgPos {
				mem = append(mem, 0)
			}
			fmt.Println("our memPos is", memPos, "symbol is ", toArray[prgPos], "memory is ", mem[memPos])
		} else if toArray[prgPos] == "<" {
			memPos -= 1
			fmt.Println("our memPos is", memPos, "symbol is ", toArray[prgPos], "memory is ", mem[memPos])
		} else if toArray[prgPos] == "+" {
			mem[memPos] += 1
			fmt.Println("our memPos is", memPos, "symbol is ", toArray[prgPos], "memory is ", mem[memPos])
		} else if toArray[prgPos] == "-" {
			mem[memPos] -= 1
			fmt.Println("our memPos is", memPos, "symbol is ", toArray[prgPos], "memory is ", mem[memPos])
		} else if toArray[prgPos] == "[" {
			fmt.Println("our memPos is", memPos, "symbol is ", toArray[prgPos], "memory is ", mem[memPos])
			if mem[memPos] == 0 {
				counterOpen := 0
				prgPos += 1
				fmt.Println("our memPos is", memPos, "symbol is ", toArray[prgPos], "memory is ", mem[memPos], "counterOpen is ", counterOpen)
				for prgPos < len(toArray) {
					if toArray[prgPos] == "]" && counterOpen == 0 {
						fmt.Println("our memPos is", memPos, "symbol is ", toArray[prgPos], "memory is ", mem[memPos], "counterOpen is ", counterOpen)
						break
					} else if  toArray[prgPos] == "[" {
						counterOpen += 1
						fmt.Println("our memPos is", memPos, "symbol is ", toArray[prgPos], "memory is ", mem[memPos], "counterOpen is ", counterOpen)
					} else if toArray[prgPos] == "]" {
						fmt.Println("our memPos is", memPos, "symbol is ", toArray[prgPos], "memory is ", mem[memPos], "counterOpen is ", counterOpen)
						counterOpen -= 1
					} 
					prgPos += 1
				} 
				
			} 
		} else if toArray[prgPos] == "]" {
			fmt.Println("our memPos is", memPos, "symbol is ", toArray[prgPos], "memory is ", mem[memPos])
			if mem[memPos] != 0 {
				for toArray[prgPos] != "[" {
					prgPos -= 1
					fmt.Println("our memPos is", memPos, "symbol is ", toArray[prgPos], "memory is ", mem[memPos])
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
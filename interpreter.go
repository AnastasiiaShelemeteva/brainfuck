package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	prgPos := 0
	mem := []int{0}
	memPos := 0

	inputVar := ("+++++++++++++++++++++++++++++++++.>,.")
	toArray := strings.Split(inputVar, "")
	for prgPos < len(toArray) {
		if toArray[prgPos] == ">" {
			memPos++
			if len(mem) <= memPos {
				mem = append(mem, 0)
			}

		} else if toArray[prgPos] == "<" {
			memPos--

		} else if toArray[prgPos] == "+" {
			mem[memPos]++

		} else if toArray[prgPos] == "-" {
			mem[memPos]--

		} else if toArray[prgPos] == "." {
			fmt.Print(string(mem[memPos]))
		} else if toArray[prgPos] == "," {
			fmt.Print("\nNo you need to write one letter. If you write more, only the first will count. Please, put here your letter: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()
			byteInput := []byte(input)
			mem[memPos] = int(byteInput[0])
		} else if toArray[prgPos] == "[" {
			if mem[memPos] == 0 {
				counterOpen := 0
				prgPos++
				for prgPos < len(toArray) {
					if toArray[prgPos] == "]" && counterOpen == 0 {
						break
					} else if toArray[prgPos] == "[" {
						counterOpen++
					} else if toArray[prgPos] == "]" {
						counterOpen--
					}
					prgPos++
				}

			}
		} else if toArray[prgPos] == "]" {

			if mem[memPos] != 0 {
				for toArray[prgPos] != "[" {
					prgPos--
				}
			}

		}
		prgPos++

	}

	fmt.Println("\nHere is an array: ", mem)

}

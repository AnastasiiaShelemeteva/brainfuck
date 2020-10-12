package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	prgPos := 0
	mem := []int{0}
	memPos := 0

	inputVar, err := ioutil.ReadFile("test.bf")
	if err != nil {
		log.Fatal(err)
	}

	toArray := inputVar
	for prgPos < len(toArray) {
		if toArray[prgPos] == 62 { // if command is ">"
			memPos++
			if len(mem) <= memPos {
				mem = append(mem, 0)
			}

		} else if toArray[prgPos] == 60 { // if command is "<"
			memPos--

		} else if toArray[prgPos] == 43 { // if command is "+"
			mem[memPos]++
			if mem[memPos] > 255 {
				mem[memPos] = 0
			}

		} else if toArray[prgPos] == 45 { // if command is "-"
			mem[memPos]--
			if mem[memPos] < 0 {
				mem[memPos] = 255
			}

		} else if toArray[prgPos] == 46 { // if command is ":"
			fmt.Print(string(mem[memPos]))
		} else if toArray[prgPos] == 44 { // if command is ","
			fmt.Print("\nNo you need to write one letter. If you write more, only the first will count. Please, put here your letter: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()
			byteInput := []byte(input)
			mem[memPos] = int(byteInput[0])
		} else if toArray[prgPos] == 91 { // if command is "["

			if mem[memPos] == 0 {
				counterOpen := 0
				prgPos++
				for prgPos < len(toArray) {
					if toArray[prgPos] == 93 && counterOpen == 0 {
						break
					} else if toArray[prgPos] == 91 {
						counterOpen++
					} else if toArray[prgPos] == 93 {
						counterOpen--
					}
					prgPos++
				}
			}
		} else if toArray[prgPos] == 93 { // if command is "]"
			if mem[memPos] != 0 {
				counterClosed := 0
				prgPos--
				for prgPos < len(toArray) {
					if toArray[prgPos] == 91 && counterClosed == 0 {
						break
					} else if toArray[prgPos] == 93 {
						counterClosed++
					} else if toArray[prgPos] == 91 {
						counterClosed--
					}
					prgPos--
				}
			}
		}

		prgPos++
	}

}

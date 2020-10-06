package main

import (
	"fmt"
	"os"
	//"io/ioutil"
	"bufio"
	
)

//prgPos := 0 // position of pointer
//var memory int = 0 // the value 
// var memPos int = 0 // the value of exactly point (what is actually wroten, like +,- etc )


func main() {

memory := 0
memPos := 0
testArr := ("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++..>++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.>+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.>+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.>+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.>++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.>+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.")
	for i, element := range testArr {
		switch element {
		case 62: //>
		i += 1
			/* if memory <= memPos {
				memory = append(memory, 0)
		}  */
		case 60: //<
		i -= 1
			if memPos <= 0 {
				fmt.Println("\nERROR. Proramm will be exited")
				os.Exit(0)
			}
		case 43://+
		memory += 1
		case 45: //-
		memory -= 1
		case 46:// .
		fmt.Println(string(memory))
		memory = 0 // here we need to say, that if next character is punkt - we have to repeat printing memory. If not, we need to set memory to zero
		case 44: 
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		memory := scanner.Text() 
		fmt.Println(memory)
		case 91: 
			if memory == 0 {
				countOpened := 0
				i += 1
				for i < len(testArr) {
					if (element == 93) && (countOpened == 0) {
					break 
				} else if element == 91 {
					countOpened += 1
				} else if element == 93 {
					countOpened -= 1
				i += 1
				}
			}
		}	
		default: 
		fmt.Println("You are doing something wrong")
	}
	
	i += 1
}
	
	
	
	
	
	
	
	
	//argsWithoutProg := os.Args[1: ] 
	
	/* fmt.Println(os.Args)
	fmt.Println(len(os.Args)) */
	

/* 	var test string = "Hey there"
	fmt.Println(test[]) 
 */

/* 	data, err := ioutil.ReadFile("try.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data)) */


}


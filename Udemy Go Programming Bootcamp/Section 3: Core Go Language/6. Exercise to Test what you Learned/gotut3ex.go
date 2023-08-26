package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("What is your age? :")

	age, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Trims whitespace from age
	age = strings.TrimSpace(age)
	iAge, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	} else {	
		grade := iAge - 5
		if iAge < 5 {
			fmt.Println(" Too young for school")
		} else if iAge == 5 {
			fmt.Println("Go to Kindergarten")
		} else if (iAge > 5) && (iAge <= 17) {
			fmt.Printf("Go to grade %d\n", grade)
		} else {
			fmt.Println("Go to college")
		}
	} 
}

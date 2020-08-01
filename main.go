package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter something:")
	for scanner.Scan() {
		fmt.Println("You said: ", scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

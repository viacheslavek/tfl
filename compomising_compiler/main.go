package main

import (
	"fmt"
	"log"
)

func main() {
	if true || false { // { // забыл закрыть фигурную скобку -> true || false evaluated but not used

		p, err := anotherFunction() // Assignment count mismatch: 1 = 2
		fmt.Println("Hello, World!")
		if err != nil { // тут ошибка Cannot convert 'nil' to type 'int'
			log.Fatalf(`")))`)
		}
		fmt.Println(p) // Unresolved reference 'p'
	}
} //

func anotherFunction() (int, error) {
	fmt.Println("Another function")
	return 0, nil
}

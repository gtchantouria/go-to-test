package main

import "fmt"

func main() {
	fmt.Println(sayHi("Guille"))
}

func sayHi(person string) string {
	return fmt.Sprintf("Hi %s", person)
}

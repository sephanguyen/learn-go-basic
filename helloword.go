package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {

	fmt.Println(CreateMessage(salutation.name, salutation.greeting))
}

func CreateMessage(name string, greeting string) string {
	return greeting + " " + name
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}

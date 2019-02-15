package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation, do func(string)) {

	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")
	do(message)
	do(alternate)

}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	message = greeting[0] + " " + name
	alternate = "HEY! " + name
	return
}

func PrintLine(message string) {
	fmt.Println(message)
	return
}

func Print(message string) {
	fmt.Print(message)
	return
}

func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s, Print)
}

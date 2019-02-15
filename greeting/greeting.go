package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {

	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if isFormal {
		do(message)
	}
	do(alternate)

}

func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
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

func CreatePrintFunction(custom string) Printer {
	return func(message string) {
		fmt.Println(message + custom)
	}

}

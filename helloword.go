package main

import "./greeting"

func main() {
	saluations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What 's up?"},
	}

	saluations = append(saluations, greeting.Salutation{"Frank", "Hi"})

	saluations.Greet(greeting.CreatePrintFunction("!!!"), true, 5)
	greeting.TypeSwitchTest("asdas")
}

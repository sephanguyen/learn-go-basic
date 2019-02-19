package main

import (
	"fmt"

	"./greeting"
)

func RenameToFrog(r greeting.Renamable) {
	r.Rename("Frog")
}

func main() {
	saluations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What 's up?"},
	}

	// saluations[0].Rename("Mom")
	// RenameToFrog(&saluations[0])
	fmt.Fprintf(&saluations[0], "the count is %d", 10)

	saluations = append(saluations, greeting.Salutation{"Frank", "Hi"})

	c := make(chan greeting.Salutation)

	for s := range c {
		fmt.Println(s.Name)
	}
}

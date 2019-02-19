package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Renamable interface {
	Rename(newName string)
}

func (salutation *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}

func (salutation Salutations) ChannelGreeter(c chan Salutation) {
	for _, s := range salutation {
		c <- s
	}
	close(c)
}

func (salsutation *Salutation) Rename(newName string) {
	salsutation.Name = newName
}

type Salutations []Salutation

type Printer func(string)

func (salutations Salutations) Greet(do Printer, isFormal bool, times int) {
	for _, s := range salutations {
		message, alternate := CreateMessage(s.Name, s.Greeting)
		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}

}

func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string{
		"Bob":  "Mr ",
		"Joe":  "Dr ",
		"Amy":  "Dr ",
		"Mary": "Mrs ",
	}

	delete(prefixMap, "Mary")

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude "
}

func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("Salutation")
	default:
		fmt.Println("unknow")

	}
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

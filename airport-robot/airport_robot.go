package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(visitor string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(visitor))
}

type Italian struct {}
func (it Italian) LanguageName() string {
	return "Italian"
}
func(it Italian) Greet(visitor string) string {
	return fmt.Sprintf("Ciao %s!", visitor)
}

type Portuguese struct {}
func (por Portuguese) LanguageName() string {
	return "Portuguese"
}
func(por Portuguese) Greet(visitor string) string {
	return fmt.Sprintf("Olá %s!", visitor)
}
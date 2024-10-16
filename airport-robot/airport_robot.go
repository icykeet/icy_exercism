package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

type Italian struct{}
type Portuguese struct{}

func (i Italian) LanguageName() string {
	return "I can speak Italian"
}
func (i Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

func (i Portuguese) LanguageName() string {
	return "I can speak Portuguese"
}
func (i Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}

func SayHello(visitorName string, greeter Greeter) string {
	return fmt.Sprintf("%s: %s", greeter.LanguageName(), greeter.Greet(visitorName))
}

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

package main

import (
	"fmt"
	"log"

	"zeisberg.net/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Ldate + log.Ltime + log.Llongfile)

	// A slice of names.
	names := []string{"Vera", "Matthias", "Julia"}
	// Request a greeting message.
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.

	fmt.Println(messages)

	for i, name := range names {
		ausgabe := fmt.Sprintf("%d: Name: '%v' \t -> wird begrüßt mit '%v'", i, name, messages[name])
		fmt.Println(ausgabe)
	}
}

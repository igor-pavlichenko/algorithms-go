package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"

	"rsc.io/quote/v4"
)

func main()  {
	fmt.Println(quote.Go())

	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

	names := []string{
		"Igor",
		"Gladys",
		"Samantha",
		"Darrin",
	}

	// Request a greeting message.
    messages, err := Hellos(names)
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

	// If no error was returned, print the returned message
    // to the console.
    fmt.Println(messages)
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for i := 0; i < len(names); i++ {
		name := names[i]
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! 🖖",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func SingleGreeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name given")
	}
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func MultipleGreetings(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := SingleGreeting(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func randomGreeting() string {
	greetings := []string{
		"Hello, %v. Welcome!",
		"Nice to meet you %v",
		"All hail our overlord %v",
	}

	return greetings[rand.Intn(len(greetings))]
}

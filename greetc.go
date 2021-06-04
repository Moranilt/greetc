package greetc

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randGreet() string {
	var length int

	formats := []string{
		"Hello, %v!",
		"Wellcome, %v! Nice job!",
		"Hey, jedi. How are you, %v?",
	}

	length = len(formats)

	return formats[rand.Intn(length)]
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name is emty!")
	}

	msg := fmt.Sprintf(randGreet(), name)

	return msg, nil
}

package greetings

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	math_rand "math/rand"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[math_rand.Intn(len(formats))]
}

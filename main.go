package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var chars [6]string
	rand.Seed(time.Now().UnixNano())
	char := strings.ToLower(randomString(6))
	chars[0] = strings.ToTitle(char[0:1])
	chars[1] = char[1:2]
	chars[2] = char[2:3]
	chars[3] = char[3:4]
	chars[4] = char[4:5]
	chars[5] = char[5:6]

	fmt.Println("A Palindrome")
	fmt.Println("************")
	var buffer bytes.Buffer
	for i := 0; i < len(chars); i++ {
		buffer.WriteString(chars[i])
	}
	for i := len(chars) - 2; i >= 0; i-- {
		buffer.WriteString(chars[i])
	}
	fmt.Println(buffer.String())
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

//covering pointers, slices, string and runes, structs, methods
//covering comments, casting

package main

import (
	"fmt"
	"strings"
)

type messagePackage struct {
	message, commentary string
	seed                int
}

func main() {
	test1()

	test2()
}

func (messagePackage *messagePackage) encryptMessagePackage() {
	newMessage := make([]rune, 0)
	for _, value := range messagePackage.message {
		//fmt.Printf("the value is: %c\n", messagePackage.message[idx])
		newMessage = append(newMessage, value+rune(messagePackage.seed))
	}

	messagePackage.message = string(newMessage)
}

func (messagePackage *messagePackage) decryptMessagePackage() {
	newMessage := make([]rune, 0)
	for _, value := range messagePackage.message {
		newMessage = append(newMessage, value-rune(messagePackage.seed))
	}

	messagePackage.message = string(newMessage)
}

func test1() {
	testMessage := messagePackage{
		message:    "this is a message",
		seed:       1,
		commentary: "the seed of this message package is one"}

	fmt.Println("the original message is: ", testMessage.message)

	testMessage.encryptMessagePackage()
	fmt.Println("after encrypting the message, the message is: ", testMessage.message)

	testMessage.decryptMessagePackage()
	fmt.Println("after decrypting the message, the message is: ", testMessage.message)
}

/*
func (messagePackage *messagePackage) verify(input ...rune) float32 {
	amountCorrect := 0
	for _, character := range input {
		if strings.Contains(string(messagePackage.message), string(character)) {
			amountCorrect++
		}

	}
	return len(messagePackage.message) / amountCorrect * 100
}

func test2() {
	testPackage := messagePackage{message: "abcdefg"}

	fmt.Println(testPackage.verify("e", "f", "g", "h", "i", "j"))
}
*/

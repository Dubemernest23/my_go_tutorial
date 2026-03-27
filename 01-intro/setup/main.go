package main

import "fmt"

// x := "kikk"
// var x int  := 8
// const p string = "h"

func main() {
	sendsofar := 45
	const sendToAdd int = 2
	sendsofar = increamentSend(sendsofar, sendToAdd)
	sendMessage()
	fmt.Println("You have sent", sendsofar, "messages")
}

func increamentSend(sendSoFar, sendToAdd int) int {
	sendSoFar = sendSoFar + sendToAdd
	return sendSoFar
}

func sendMessage() (x, y string) {
	fmt.Println("Message sent")
	return y, x // explicitly return y and x
}

func sendMessage2() (x, y string) {
	fmt.Println("Message sent")
	return // implicitly return x and y
}

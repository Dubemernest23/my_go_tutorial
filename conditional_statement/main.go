package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Conditional statement")
	x := calcStudentsScore(20)
	fmt.Printf("Result statement: %v \n", x)
	y := shortStatement(40)
	fmt.Printf("info: %v \n", y)
	loopNum()
	checkDay(6)
}

func calcStudentsScore(score int) string {
	if score >= 70 {
		return "Student scored A"
	} else if score >= 60 {
		return "Student scored B"
	} else if score >= 50 {
		return "student scored C"
	} else {
		return "student should resit this is exam"
	}
}

func shortStatement(bal int) string { //
	items := 3
	price := 400
	walletBal := bal
	if total := items * price; total >= walletBal {
		// in short if-else-conditionals, you can declare a variable in the if scopee and variales declared there
		// will only be used in there.
		// strconv.itoa() is used in converting num to string
		return "Insufficient amount when the total is " + strconv.Itoa(total)
	}
	return "Payment made successfully"
}

func loopNum() {
	// for loop
	n := 30
	sum := 2
	for i := 0; i <= n; i++ {
		// init part, condition, increament
		sum = sum * i
		fmt.Printf("Sum prints %v \n", sum)
		fmt.Printf("I prints %v \n", i)
	}
}

func checkDay(val int) {

	switch val {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednestday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Print("Unknow day of the week")
	}
	// case
}

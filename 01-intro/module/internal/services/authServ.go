package services

import (
	"fmt"
	// "module/pkg/logger"
)

// you can only export functions that started wih capital letter

func validateDetails() {

	fmt.Print("Validation completed")
}

func GetUser(user string) string {
	validateDetails()
	return user
}

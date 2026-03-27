package main

import (
	"fmt"
	"strconv"
)

func main() {
	// run() returns an error, so we handle it
	if err := run(); err != nil {
		fmt.Println("run failed:", err)
		return
	}
	fmt.Println("program finished successfully")
}

func run() error {
	fmt.Println("run() started")
	defer fmt.Println("run() finished") // always prints when run() exits

	val, err := parseVal("50330")
	if err != nil {
		return fmt.Errorf("run: could not parse value: %w", err)
	}

	fmt.Println("parsed value:", val)
	return nil // nil means no error
}

func parseVal(s string) (int, error) {
	val, err := strconv.Atoi(s) // Atoi = "ASCII to integer", like parseInt() in TS
	if err != nil {
		return 0, fmt.Errorf("parseVal: %q is not a valid number", s)
	}
	return val, nil
}

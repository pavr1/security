package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	expressions = []string{
		"^[a-zA-Z0-9]+$",
		"^[a-zA-Z0-9]:[a-zA-Z]{3}[a-zA-Z0-9]",
	}
)

func main() {
	var (
		err          error
		fileContents string
	)

	if len(os.Args) < 3 {
		fmt.Printf("no arguments supplied in pre-commit check\n")
		os.Exit(1)
	}

	fileName := os.Args[2]

	if fileContents, err = readFile(fileName); err != nil {
		fmt.Printf("Error reading file for security check: %s\n", err.Error())
		os.Exit(1)
	}

	if isvalid, info := isValid(fileName, fileContents); !isvalid {
		fmt.Printf("invalid! => %s\n", info)
		os.Exit(1)

		return
	} else {
		fmt.Println("valid!")
	}
	os.Exit(0)
}

func readFile(fileName string) (string, error) {
	var (
		err error
		b   []byte
	)

	if b, err = os.ReadFile(fileName); err != nil {
		return "", err
	}

	return string(b), nil
}

func isValid(fileName, fileContents string) (bool, string) {
	if strings.Contains(fileName, ".env") {
		return false, "You are not allowed to commit .env files. This poses a security risk.\n"
	}

	fmt.Println("Not an .env file. Proceeding...")

	//avoid validating pre-commit file
	if strings.Contains(fileName, "pre-commit") {
		return true, ""
	}

	fmt.Println("Not a pre-commit file. Proceeding...")

	for i, exp := range expressions {
		r, _ := regexp.Compile(exp)

		match := r.MatchString(fileContents)

		if match {
			return false, fmt.Sprintf("It looks like you are attemting to set token in file '%s'. This is not allowed\n", fileName)
		} else {
			fmt.Printf("Expression #%d valid\n", i+1)
		}
	}

	return true, ""
}

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

//greater than 15 chars
var (
	expressions = []string{
		//"[a-zA-Z0-9]+ := \"\\w+.{15,}\"",
		"\"\\w+.{15,}\"",
		"`\\w+.{15,}`",
	}
)

func main() {
	var (
		err          error
		fileContents string
	)

	if len(os.Args) < 3 {
		fmt.Print("	no arguments supplied in pre-commit check.")
		os.Exit(1)
	}

	fileName := os.Args[2]

	if fileContents, err = readFile(fileName); err != nil {
		fmt.Printf("	error reading file for security check: %s.", err.Error())
		os.Exit(1)
	}

	if isvalid, info := isValid(fileName, fileContents); !isvalid {
		fmt.Printf("	invalid! => %s\n", info)
		os.Exit(1)

		return
	} else {
		fmt.Print("	valid.")
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
		return false, "	You are not allowed to commit .env files. This poses a security risk.\n"
	}

	//avoid validating pre-commit file
	if strings.Contains(fileName, "pre-commit") {
		return true, ""
	}

	for i, exp := range expressions {
		r, err := regexp.Compile(exp)
		if err != nil {
			return false, fmt.Sprintf("	there was an error compiling regex '%s'. Error: %v", exp, err)
		}

		match := r.Match([]byte(fileContents))

		fmt.Printf("	Exp: %d, Match: %t\n", i+1, match)

		if match {
			return false, fmt.Sprintf("looks like you left a token or key in file '%s'. This is not allowed!\n", fileName)
		}
	}

	return true, ""
}

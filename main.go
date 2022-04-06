package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	expression = "(?i)%s\\s*(=|:=|==|===)\\s*[\"'`]{1}.*?[\"'`]"
)

func main() {
	var (
		err          error
		fileContents string
		keys         = []string{"api_key", "shppa_", "github_token"}
	)

	fmt.Printf("Argument Length: %d\n", len(os.Args))
	fmt.Printf("Argument1: %s\n", os.Args[0])

	if len(os.Args) < 2 {
		fmt.Printf("no arguments supplied in pre-commit check\n")
		os.Exit(1)
	}

	fmt.Printf("Argument2: %s\n", os.Args[1])

	fileName := os.Args[0]

	if fileContents, err = readFile(fileName); err != nil {
		fmt.Printf("Error reading file for security check: %s\n", err.Error())
		os.Exit(1)
	}

	for _, key := range keys {
		if isvalid, info := isValid(fileName, fileContents, key); !isvalid {
			fmt.Println(info)
			os.Exit(1)

			return
		}
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

func isValid(fileName, fileContents, key string) (bool, string) {
	if strings.HasSuffix(fileName, ".env") {
		return false, "You are not allowed to commit .env files. This poses a security risk.\n"
	}

	//avoid these checks in pre-commit file
	if strings.Contains(fileName, "pre-commit") {
		return true, ""
	}

	exp := fmt.Sprintf(expression, key)
	r, _ := regexp.Compile(exp)

	valid := r.MatchString(fileContents)

	if !valid {
		return false, fmt.Sprintf("It looks like you are attemting to set a '%s' in %s, this is not allowed.\n", key, fileName)
	} else {
		return true, ""
	}
}

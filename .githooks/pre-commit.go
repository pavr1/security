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
		keys         = []string{"api_key", "shppa_", "github_token", "aws_access_key_id", "aws_secret_access_key", "aws_session_token", "github_personal_access_token", "github_oauth_access_token",
			"github_refresh_token", "github_app_installation_access_token", "github_ssh_private_key", "shippo_live_api_token", "shopify_app_shared_secret", "shopify_access_token"}
	)

	fmt.Printf("Argument Length: %d\n", len(os.Args))
	fmt.Printf("Argument1: %s\n", os.Args[0])

	if len(os.Args) < 3 {
		fmt.Printf("no arguments supplied in pre-commit check\n")
		os.Exit(1)
	}

	fmt.Printf("Argument2: %s\n", os.Args[2])

	fileName := os.Args[2]

	if fileContents, err = readFile(fileName); err != nil {
		fmt.Printf("Error reading file for security check: %s\n", err.Error())
		os.Exit(1)
	}

	for _, key := range keys {
		fmt.Printf(" - validating '%s'...", key)

		if isvalid, info := isValid(fileName, fileContents, key); !isvalid {
			fmt.Printf("invalid! File Name: '%s'  => "+info+"\n", fileName)
			os.Exit(1)

			return
		} else {
			fmt.Println("valid!")
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

	fmt.Print("not an .env file...")

	//avoid validating pre-commit file
	if strings.Contains(fileName, "pre-commit") {
		return true, ""
	}

	fmt.Print("not pre-commit file...")

	exp := fmt.Sprintf(expression, key)
	fmt.Printf("expression: %s...", exp)

	r, _ := regexp.Compile(exp)

	valid := r.MatchString(fileContents)
	fmt.Printf("is valid: %t", valid)

	if !valid {
		return false, fmt.Sprintf("It looks like you are attemting to set a '%s' in %s, this is not allowed.\n", key, fileName)
	} else {
		return true, ""
	}
}

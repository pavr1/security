package testchild

import "fmt"

const (
	AECCancelCustomerNotFound = "3"
	AECCancelOrderNotFound    = "4"
)

/**
Tests
pavr1:ghp_Gi3thu5bTo8ke9n0
1. constants
2. variables
3. hardcoded values
4. double quotes and ``
**/

func testchild() {
	var selectedKeyToValudate string
	selectedKeyToValudate = ""

	fmt.Println(selectedKeyToValudate)

	anotherTest("12")
	anotherTest(`56`)
}

func anotherTest(val string) string {
	return val
}

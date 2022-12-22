package github_test

import (
	"fmt"

	"github.com/cherryramatisdev/s/github"
)

func ExampleGetCurrentOrg() {
	fmt.Println(github.GetCurrentOrg())
	// Output:
	// cherryramatisdev
}

func ExampleGetCurrentRepo() {
	fmt.Println(github.GetCurrentRepo())
	// Output:
	// s
}

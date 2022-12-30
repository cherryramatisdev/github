package github

import (
	_ "embed"
	"fmt"
	"regexp"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

//go:embed doc.md
var githubdoc string

var Cmd = &Z.Cmd{
	Name:        `github`,
	Aliases:     []string{`gh`},
	Usage:       `[help]`,
	Description: githubdoc,
	Commands:    []*Z.Cmd{help.Cmd, prCmd, issueCmd},
}

//go:embed issue.md
var issuedoc string

var issueCmd = &Z.Cmd{
	Name:        `issue`,
	Description: issuedoc,
	Commands:    []*Z.Cmd{help.Cmd, listIssueCmd},
}

var listIssueCmd = &Z.Cmd{
	Name:        `list`,
	Description: enhancementdoc,
	Call: func(x *Z.Cmd, args ...string) error {
		return Z.Exec("gh", "issue", "list")
	},
}

//go:embed pullrequest.md
var pullrequestdoc string

var prCmd = &Z.Cmd{
	Name:        `pr`,
	Description: pullrequestdoc,
	Commands:    []*Z.Cmd{help.Cmd, createEnhancementPrCmd},
}

//go:embed enhancement.md
var enhancementdoc string

var REVIEWERS = []string{
	"hueldera",
	"guilhermeocosta",
	"HMilao",
	"victorakioz",
	"pedrohmp",
	"jorcelinojunior",
}

var createEnhancementPrCmd = &Z.Cmd{
	Name:        `enhancement`,
	Aliases:     []string{`en`},
	Description: enhancementdoc,
	Call: func(x *Z.Cmd, args ...string) error {
		branches := GetBranches()
		currentbranch := GetCurrentBranch()
		pat := regexp.MustCompile(`TEC-(\d+)`)
		matches := pat.FindStringSubmatch(currentbranch)

		if pat.Match([]byte(currentbranch)) {
			fmt.Println(matches)
		} else {
			fmt.Println("nao deu match")
			fmt.Println(branches)
		}

		return nil
	},
}

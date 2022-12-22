package github

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
)

// GetCurrentOrg return the current org name by parsing the result of `git remote -v` command
func GetCurrentOrg() string {
	var out bytes.Buffer
	cmd := exec.Command("git", "remote", "-v")
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return ""
	}

	pat := regexp.MustCompile(`git@github.com:(.*)/(.*)\.git`)
	matches := pat.FindStringSubmatch(out.String())

	return matches[1]
}

// GetCurrentRepo return the current repo name by parsing the result of `git remote -v` command
func GetCurrentRepo() string {
	var out bytes.Buffer
	cmd := exec.Command("git", "remote", "-v")
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return ""
	}

	pat := regexp.MustCompile(`git@github.com:(.*)/(.*)\.git`)
	matches := pat.FindStringSubmatch(out.String())

	return matches[2]
}

func GetCurrentBranch() string {
	var out bytes.Buffer
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return ""
	}

	return out.String()
}

func GetBranches() []string {
	var out bytes.Buffer
	cmd := exec.Command("git", "branch")
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return []string{}
	}

	sanitize := strings.Replace(out.String(), "*", "", -1)
	return strings.Split(sanitize, " ")
}

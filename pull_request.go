package github

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type Pr struct {
	Org         string
	Repo        string
	Prefixtoken string
}

// gh api   -H Accept: application/vnd.github+json   /repos/lami-health/website/branches/feature/TEC-1693
// Exist check if branch exist on upstream
func (pr *Pr) Exist() bool {
	var out bytes.Buffer
	url := fmt.Sprintf("/repos/%s/%s/branches/feature/%s", pr.Org, pr.Repo, pr.Prefixtoken)
	cmd := exec.Command("gh", "api", "-H", "Accept: application/vnd.github+json", url)
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return false
	}

	return strings.Contains(out.String(), fmt.Sprintf("feature/%s", pr.Prefixtoken))
}

func (pr *Pr) GetUrl() string {
	var out bytes.Buffer
	cmd := exec.Command("gh", "pr", "view")
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return ""
	}

	pat := regexp.MustCompile(`https?://github.com.*`)
	url := pat.FindString(out.String())

	return url
}

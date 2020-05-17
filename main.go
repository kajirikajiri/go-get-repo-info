package main

import (
	"context"
	"fmt"
	"flag"
	"io/ioutil"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os/user"
	"strings"
	"regexp"
)

var regxNewline = regexp.MustCompile(`\r\n|\r|\n`) //throw panic if fail

func convNewline(str, nlcode string) string {
    return regxNewline.Copy().ReplaceAllString(str, nlcode)
}


func useIoutilReadFile(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
			panic(err)
	}

	return string(bytes)
}

func main() {

	usr, _ := user.Current()
	f := strings.Replace("~/go-get-repo-info-access-token",  "~", usr.HomeDir, 1)
	fmt.Println(f)

	accessToken := useIoutilReadFile(f)
	regexedAccessToken := convNewline(accessToken, "")

	if accessToken == "" {
		panic("access-token is blank. create go-get-repo-info/access-token. and write access-token. not \\n")
	}

	var (
		status = flag.String("status", "open", "open|close|all" )
		organization = flag.String("org", "kajirikajiri", "ex)OnetapInc|kajirikajiri" )
		branch = flag.String("branch", "", "ex}feature/issue-700" )
		base = flag.String("base", "develop", "ex}develop|master|release" )
		sort = flag.String("sort", "created", "created|updated|popularity|long-running" )
		direction = flag.String("direction", "desc", "asc|desc" )
		repo = flag.String("repo", "go-get-repo-info", "repository name" )
	)
	flag.Parse()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: regexedAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opts := &github.PullRequestListOptions{*status, *organization + ":" + *branch, *base, *sort, *direction, github.ListOptions{Page: 1}}

	pulls, _, err := client.PullRequests.List(ctx, *organization, *repo, opts)
	if err != nil {
		fmt.Print(err)
	}

	for _ ,pull := range pulls {
		fmt.Print(*pull.Title)
	}
}

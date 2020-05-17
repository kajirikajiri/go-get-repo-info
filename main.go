package main

import (
	"context"
	"fmt"
	"flag"
	"io/ioutil"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func useIoutilReadFile(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
			panic(err)
	}

	return string(bytes)
}

func main() {

	accessToken := useIoutilReadFile("~/go-get-repo-info-access-token")

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
		&oauth2.Token{AccessToken: accessToken},
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

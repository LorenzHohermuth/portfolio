package github

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofor-little/env"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	owner =  "lorenzhohermuth"
	repo = "portfolio"
	basePath = "/"
)

type response struct {
	content *github.RepositoryContent
	err			error
}

func FetchGHFile(ctx context.Context, path string) (string, error){
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 500)
	defer cancel()

	resChan := make(chan response)

	go func() {
		res, err := fetchGithub(ctx, path)
		resChan <- response{
			content: res,
			err: err,
		}
	}()

	for {
		select {
		case <- ctx.Done():
			return "", fmt.Errorf("fetching from Github took to long")
		case resp := <- resChan:
			text, errRepoCont := resp.content.GetContent()
			if errRepoCont != nil {
				panic(errRepoCont)
			}
			
			return text, resp.err
		}
	}

}

func fetchGithub(ctx context.Context, path string) (*github.RepositoryContent, error) {
	client := github.NewClient(getAuthToken())
	fileContent, _, _, err := client.Repositories.GetContents(ctx, owner, repo, path, nil)
	return fileContent, err
}

func getAuthToken() *http.Client {
	ctx := context.Background()
	pat, err := env.MustGet("PAT")
	if err != nil {
		panic(err)
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: pat},
	)
	return oauth2.NewClient(ctx, ts)
}

package github

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Client interface {
	FetchRepoDescription(ctx context.Context, owner, name string) (string, error)
	FetchFile(ctx context.Context, owner, name, branch, filePath string) (*FileBlob, error)
}

func NewClient(token string) Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	return &clientImpl{client: client}
}

type clientImpl struct {
	client *githubv4.Client
}

// FetchRepoDescription fetches description of repo with owner and name.
func (c *clientImpl) FetchRepoDescription(ctx context.Context, owner, name string) (string, error) {
	var q struct {
		Repository struct {
			Description string
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}

	err := c.client.Query(ctx, &q, variables)
	return q.Repository.Description, err
}

// FetchFile fetches a file content
func (c *clientImpl) FetchFile(ctx context.Context, owner, name, branch, filePath string) (*FileBlob, error) {

	var q struct {
		Repository struct {
			File FileBlob `graphql:"object(expression: $file)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
		"file":  githubv4.String(fmt.Sprintf("%s:%s", branch, filePath)),
	}

	err := c.client.Query(ctx, &q, variables)
	return &q.Repository.File, err
}

type FileBlob struct {
	Blob Blob `graphql:"... on Blob"`
}

type Blob struct {
	Text string
}

type FileContent struct {

	//CommitUrl string
	//Name     string
	//Type     string

	//IsBinary githubv4.Boolean
}

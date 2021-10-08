package internal

import (
	"fmt"
	"github.com/weaveworks/weave-gitops/pkg/gitproviders"
)

type gitProviderClient struct {
	token string
}

func NewGitProviderClient(token string) gitproviders.Client {
	return &gitProviderClient{
		token: token,
	}
}

// GetProvider returns a GitProvider containing either the token stored in the <git provider>_TOKEN env var
// or a token retrieved via the CLI auth flow
func (c *gitProviderClient) GetProvider(repoUrl gitproviders.NormalizedRepoURL, getAccountType gitproviders.AccountTypeGetter) (gitproviders.GitProvider, error) {
	provider, err := gitproviders.New(gitproviders.Config{Provider: repoUrl.Provider(), Token: c.token}, repoUrl.Owner(), getAccountType)
	if err != nil {
		return nil, fmt.Errorf("error creating git provider client: %w", err)
	}

	return provider, nil
}

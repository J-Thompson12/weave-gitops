package internal

import (
	"context"
	"fmt"
	"github.com/weaveworks/weave-gitops/pkg/gitproviders"
	"github.com/weaveworks/weave-gitops/pkg/logger"
	"github.com/weaveworks/weave-gitops/pkg/services/auth"
	"os"
)

type gitProviderClient struct {
	stdout        *os.File
	lookupEnvFunc func(key string) (string, bool)
	log           logger.Logger
}

func NewGitProviderClient(stdout *os.File, lookupEnvFunc func(key string) (string, bool), log logger.Logger) gitproviders.Client {
	return &gitProviderClient{
		stdout:        stdout,
		lookupEnvFunc: lookupEnvFunc,
		log:           log,
	}
}

// GetProvider returns a GitProvider containing either the token stored in the <git provider>_TOKEN env var
// or a token retrieved via the CLI auth flow
func (c *gitProviderClient) GetProvider(repoUrl gitproviders.NormalizedRepoURL, getAccountType gitproviders.AccountTypeGetter) (gitproviders.GitProvider, error) {
	ctx := context.Background()

	tokenVarName, err := getTokenVarName(repoUrl.Provider())
	if err != nil {
		return nil, fmt.Errorf("could not determine git provider token name: %w", err)
	}

	authHandler, err := auth.NewAuthCLIHandler(repoUrl.Provider())
	if err != nil {
		return nil, fmt.Errorf("error initializing cli auth handler: %w", err)
	}

	token, exists := c.lookupEnvFunc(tokenVarName)
	if !exists {
		c.log.Warningf("Setting the %q environment variable to a valid token will allow ongoing use of the CLI without requiring a browser-based auth flow...\n", tokenVarName)

		generatedToken, err := authHandler(ctx, c.stdout)
		if err != nil {
			return nil, fmt.Errorf("could not complete auth flow: %w", err)
		}

		token = generatedToken
	} else if err != nil {
		return nil, fmt.Errorf("could not get access token: %w", err)
	}

	provider, err := gitproviders.New(gitproviders.Config{Provider: repoUrl.Provider(), Token: token}, repoUrl.Owner(), getAccountType)
	if err != nil {
		return nil, fmt.Errorf("error creating git provider client: %w", err)
	}

	return provider, nil
}

func getTokenVarName(providerName gitproviders.GitProviderName) (string, error) {
	switch providerName {
	case gitproviders.GitProviderGitHub:
		return "GITHUB_TOKEN", nil
	case gitproviders.GitProviderGitLab:
		return "GITLAB_TOKEN", nil
	default:
		return "", fmt.Errorf("unknown git provider: %q", providerName)
	}
}

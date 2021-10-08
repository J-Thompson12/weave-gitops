package gitproviders

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . Client
type Client interface {
	GetProvider(repoUrl NormalizedRepoURL, getAccountType AccountTypeGetter) (GitProvider, error)
}

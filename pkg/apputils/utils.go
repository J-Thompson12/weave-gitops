package apputils

import (
	"context"
	"fmt"
	wego "github.com/weaveworks/weave-gitops/api/v1alpha1"
	"github.com/weaveworks/weave-gitops/pkg/flux"
	"github.com/weaveworks/weave-gitops/pkg/git"
	"github.com/weaveworks/weave-gitops/pkg/gitproviders"
	"github.com/weaveworks/weave-gitops/pkg/kube"
	"github.com/weaveworks/weave-gitops/pkg/logger"
	"github.com/weaveworks/weave-gitops/pkg/osys"
	"github.com/weaveworks/weave-gitops/pkg/runner"
	"github.com/weaveworks/weave-gitops/pkg/services/app"
	"github.com/weaveworks/weave-gitops/pkg/services/auth"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . AppFactory

// AppFactory provides helpers for generating various WeGO service objects at runtime.
type AppFactory interface {
	GetKubeService() (kube.Kube, error)
	GetAppService(ctx context.Context) (app.AppService, error)
	GetGitClients(ctx context.Context, gpClient gitproviders.Client, params AppServiceParams) (git.Git, gitproviders.GitProvider, error)
}

type AppServiceParams struct {
	URL              string
	ConfigURL        string
	Namespace        string
	IsHelmRepository bool
	DryRun           bool
}

func NewAppServiceParams(app *wego.Application, dryRun bool) AppServiceParams {
	isHelmRepository := app.Spec.SourceType == wego.SourceTypeHelm

	return AppServiceParams{
		URL:              app.Spec.URL,
		ConfigURL:        app.Spec.ConfigURL,
		Namespace:        app.Namespace,
		IsHelmRepository: isHelmRepository,
		DryRun:           dryRun,
	}
}

type defaultAppFactory struct {
	fluxClient flux.Flux
	log        logger.Logger
	osClient   osys.Osys
}

func NewAppFactory(osClient osys.Osys, cliRunner *runner.CLIRunner, log logger.Logger) AppFactory {
	return &defaultAppFactory{
		fluxClient: flux.New(osClient, cliRunner),
		log:        log,
		osClient:   osClient,
	}
}

func (f *defaultAppFactory) GetAppService(ctx context.Context) (app.AppService, error) {
	kubeClient, err := f.GetKubeService()
	if err != nil {
		return nil, fmt.Errorf("error initializing clients: %w", err)
	}

	return app.New(ctx, f.log, f.fluxClient, kubeClient, f.osClient), nil
}

func (f *defaultAppFactory) GetKubeService() (kube.Kube, error) {
	kubeClient, _, err := kube.NewKubeHTTPClient()
	if err != nil {
		return nil, fmt.Errorf("error creating k8s http client: %w", err)
	}

	return kubeClient, nil
}

func IsClusterReady(log logger.Logger) error {
	kube, _, err := kube.NewKubeHTTPClient()
	if err != nil {
		return fmt.Errorf("error creating k8s http client: %w", err)
	}

	return app.IsClusterReady(log, kube)
}

func (f *defaultAppFactory) GetGitClients(ctx context.Context, gpClient gitproviders.Client, params AppServiceParams) (git.Git, gitproviders.GitProvider, error) {
	isExternalConfig := app.IsExternalConfigUrl(params.ConfigURL)

	var providerUrl string

	switch {
	case !params.IsHelmRepository:
		providerUrl = params.URL
	case isExternalConfig:
		providerUrl = params.ConfigURL
	default:
		return nil, nil, nil
	}

	normalizedUrl, err := gitproviders.NewNormalizedRepoURL(providerUrl)
	if err != nil {
		return nil, nil, fmt.Errorf("error normalizing url: %w", err)
	}

	kube, _, err := kube.NewKubeHTTPClient()
	if err != nil {
		return nil, nil, fmt.Errorf("error creating k8s http client: %w", err)
	}

	targetName, err := kube.GetClusterName(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting target name: %w", err)
	}

	authSvc, err := f.getAuthService(normalizedUrl, gpClient, params.DryRun)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating auth service: %w", err)
	}

	var appClient, configClient git.Git

	if !params.IsHelmRepository {
		// We need to do this even if we have an external config to set up the deploy key for the app repo
		appRepoClient, appRepoErr := authSvc.CreateGitClient(ctx, normalizedUrl, targetName, params.Namespace)
		if appRepoErr != nil {
			return nil, nil, appRepoErr
		}

		appClient = appRepoClient
	}

	if isExternalConfig {
		normalizedConfigUrl, err := gitproviders.NewNormalizedRepoURL(params.ConfigURL)
		if err != nil {
			return nil, nil, fmt.Errorf("error normalizing url: %w", err)
		}

		configRepoClient, configRepoErr := authSvc.CreateGitClient(ctx, normalizedConfigUrl, targetName, params.Namespace)
		if configRepoErr != nil {
			return nil, nil, configRepoErr
		}

		configClient = configRepoClient
	} else {
		configClient = appClient
	}

	return configClient, authSvc.GetGitProvider(), nil
}

func (f *defaultAppFactory) getAuthService(normalizedUrl gitproviders.NormalizedRepoURL, gpClient gitproviders.Client, dryRun bool) (auth.AuthService, error) {
	var (
		gitProvider gitproviders.GitProvider
		err         error
	)

	if dryRun {
		if gitProvider, err = gitproviders.NewDryRun(); err != nil {
			return nil, fmt.Errorf("error creating git provider client: %w", err)
		}
	} else {
		if gitProvider, err = gpClient.GetProvider(normalizedUrl, gitproviders.GetAccountType); err != nil {
			return nil, fmt.Errorf("error obtaining git provider token: %w", err)
		}
	}

	_, rawClient, err := kube.NewKubeHTTPClient()
	if err != nil {
		return nil, fmt.Errorf("error creating k8s http client: %w", err)
	}

	return auth.NewAuthService(f.fluxClient, rawClient, gitProvider, f.log)
}

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
	"k8s.io/apimachinery/pkg/types"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . AppFactory

// AppFactory provides helpers for generating various WeGO service objects at runtime.
type AppFactory interface {
	GetKubeService() (kube.Kube, error)
	GetAppService(ctx context.Context, providerClient gitproviders.Client, name, namespace string) (app.AppService, error)
	GetAppServiceForAdd(ctx context.Context, providerClient gitproviders.Client, params AppServiceParams) (app.AppService, error)
}

type AppServiceParams struct {
	URL              string
	ConfigURL        string
	Namespace        string
	IsHelmRepository bool
	DryRun           bool
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

func (f *defaultAppFactory) GetAppService(ctx context.Context, providerClient gitproviders.Client, appName, namespace string) (app.AppService, error) {
	kubeClient, err := f.GetKubeService()
	if err != nil {
		return nil, fmt.Errorf("error initializing clients: %w", err)
	}

	configClient, gitProvider, err := f.getGitClientsForApp(ctx, providerClient, appName, namespace, false)
	if err != nil {
		return nil, fmt.Errorf("error getting git clients: %w", err)
	}

	return app.New(ctx, f.log, configClient, gitProvider, f.fluxClient, kubeClient, f.osClient), nil
}

func (f *defaultAppFactory) GetAppServiceForAdd(ctx context.Context, providerClient gitproviders.Client, params AppServiceParams) (app.AppService, error) {
	kubeClient, err := f.GetKubeService()
	if err != nil {
		return nil, fmt.Errorf("error initializing clients: %w", err)
	}

	configClient, gitProvider, err := f.getGitClients(ctx, providerClient, params.URL, params.ConfigURL, params.Namespace, params.IsHelmRepository, params.DryRun)
	if err != nil {
		return nil, fmt.Errorf("error getting git clients: %w", err)
	}

	return app.New(ctx, f.log, configClient, gitProvider, f.fluxClient, kubeClient, f.osClient), nil
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

func (f *defaultAppFactory) getGitClientsForApp(ctx context.Context, gpClient gitproviders.Client, appName string, namespace string, dryRun bool) (git.Git, gitproviders.GitProvider, error) {
	kube, _, err := kube.NewKubeHTTPClient()
	if err != nil {
		return nil, nil, fmt.Errorf("error creating k8s http client: %w", err)
	}

	app, err := kube.GetApplication(ctx, types.NamespacedName{Namespace: namespace, Name: appName})
	if err != nil {
		return nil, nil, fmt.Errorf("could not retrieve application %q: %w", appName, err)
	}

	isHelmRepository := app.Spec.SourceType == wego.SourceTypeHelm

	return f.getGitClients(ctx, gpClient, app.Spec.URL, app.Spec.ConfigURL, namespace, isHelmRepository, dryRun)
}

func (f *defaultAppFactory) getGitClients(ctx context.Context, gpClient gitproviders.Client, url, configUrl, namespace string, isHelmRepository bool, dryRun bool) (git.Git, gitproviders.GitProvider, error) {
	isExternalConfig := app.IsExternalConfigUrl(configUrl)

	var providerUrl string

	switch {
	case !isHelmRepository:
		providerUrl = url
	case isExternalConfig:
		providerUrl = configUrl
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

	authsvc, err := f.getAuthService(normalizedUrl, gpClient, dryRun)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating auth service: %w", err)
	}

	var appClient, configClient git.Git

	if !isHelmRepository {
		// We need to do this even if we have an external config to set up the deploy key for the app repo
		appRepoClient, appRepoErr := authsvc.CreateGitClient(ctx, normalizedUrl, targetName, namespace)
		if appRepoErr != nil {
			return nil, nil, appRepoErr
		}

		appClient = appRepoClient
	}

	if isExternalConfig {
		normalizedConfigUrl, err := gitproviders.NewNormalizedRepoURL(configUrl)
		if err != nil {
			return nil, nil, fmt.Errorf("error normalizing url: %w", err)
		}

		configRepoClient, configRepoErr := authsvc.CreateGitClient(ctx, normalizedConfigUrl, targetName, namespace)
		if configRepoErr != nil {
			return nil, nil, configRepoErr
		}

		configClient = configRepoClient
	} else {
		configClient = appClient
	}

	return configClient, authsvc.GetGitProvider(), nil
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

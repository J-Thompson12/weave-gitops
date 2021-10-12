package add

// Provides support for adding an application to gitops managment.

import (
	"context"
	"fmt"
	"github.com/weaveworks/weave-gitops/cmd/internal"
	"github.com/weaveworks/weave-gitops/pkg/logger"
	"github.com/weaveworks/weave-gitops/pkg/osys"
	"github.com/weaveworks/weave-gitops/pkg/runner"
	"os"
	"path/filepath"
	"strings"

	"github.com/lithammer/dedent"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/weaveworks/weave-gitops/cmd/gitops/version"
	"github.com/weaveworks/weave-gitops/pkg/apputils"
	"github.com/weaveworks/weave-gitops/pkg/git"
	"github.com/weaveworks/weave-gitops/pkg/git/wrapper"
	"github.com/weaveworks/weave-gitops/pkg/services/app"
)

const (
	SSHAuthSock = "SSH_AUTH_SOCK"
)

var params app.AddParams

var Cmd = &cobra.Command{
	Use:   "add [--name <name>] [--url <url>] [--branch <branch>] [--path <path within repository>] <repository directory>",
	Short: "Add a workload repository to a gitops cluster",
	Long: strings.TrimSpace(dedent.Dedent(`
        Associates an additional application in a git repository with a gitops cluster so that its contents may be managed via GitOps
    `)),
	Example: `
  # Add application to gitops control from local git repository
  gitops app add .

  # Add podinfo application to gitops control from github repository
  gitops app add --url git@github.com:myorg/podinfo

  # Get status of podinfo application
  gitops app status podinfo
`,
	RunE:          runCmd,
	SilenceUsage:  true,
	SilenceErrors: true,
	PostRun: func(cmd *cobra.Command, args []string) {
		version.CheckVersion(version.CheckpointParamsWithFlags(version.CheckpointParams(), cmd))
	},
}

func init() {
	Cmd.Flags().StringVar(&params.Name, "name", "", "Name of application")
	Cmd.Flags().StringVar(&params.Url, "url", "", "URL of remote repository")
	Cmd.Flags().StringVar(&params.Path, "path", app.DefaultPath, "Path of files within git repository")
	Cmd.Flags().StringVar(&params.Branch, "branch", "", "Branch to watch within git repository")
	Cmd.Flags().StringVar(&params.DeploymentType, "deployment-type", app.DefaultDeploymentType, "deployment type [kustomize, helm]")
	Cmd.Flags().StringVar(&params.Chart, "chart", "", "Specify chart for helm source")
	Cmd.Flags().StringVar(&params.AppConfigUrl, "app-config-url", "", "URL of external repository (if any) which will hold automation manifests; NONE to store only in the cluster")
	Cmd.Flags().StringVar(&params.HelmReleaseTargetNamespace, "helm-release-target-namespace", "", "Namespace in which to deploy a helm chart; defaults to the gitops installation namespace")
	Cmd.Flags().BoolVar(&params.DryRun, "dry-run", false, "If set, 'gitops app add' will not make any changes to the system; it will just display the actions that would have been taken")
	Cmd.Flags().BoolVar(&params.AutoMerge, "auto-merge", false, "If set, 'gitops app add' will merge automatically into the set --branch")
}

func ensureUrlIsValid() error {
	if params.Url == "" {
		// Find the url using an unauthenticated git client. We just need to read the URL.
		// params.Dir must be defined here because we already checked for it above.
		repoUrlString, repoUrlErr := git.New(nil, wrapper.NewGoGit()).GetRemoteUrl(params.Dir, "origin")
		if repoUrlErr != nil {
			return fmt.Errorf("could not get remote url for directory %s: %w", params.Dir, repoUrlErr)
		}

		params.Url = repoUrlString
	}

	return nil
}

func runCmd(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	params.Namespace, _ = cmd.Parent().Flags().GetString("namespace")

	if params.Url != "" && len(args) > 0 {
		return fmt.Errorf("you should choose either --url or the app directory")
	}

	if len(args) > 0 {
		path, err := filepath.Abs(args[0])
		if err != nil {
			return errors.Wrap(err, "failed to get absolute path for the repo directory")
		}

		params.Dir = path
	}

	if urlErr := ensureUrlIsValid(); urlErr != nil {
		return urlErr
	}

	log := logger.NewCLILogger(os.Stdout)
	if readyErr := apputils.IsClusterReady(log); readyErr != nil {
		return readyErr
	}

	factory := apputils.NewAppFactory(osys.New(), &runner.CLIRunner{}, log)

	providerClient := internal.NewGitProviderClient(os.Stdout, os.LookupEnv, log)
	appService, appError := factory.GetAppService(ctx)
	if appError != nil {
		return fmt.Errorf("failed to create app service: %w", appError)
	}

	gitClient, gitProvider, gitErr := factory.GetGitClients(ctx, providerClient, apputils.AppServiceParams{
		URL:              params.Url,
		ConfigURL:        params.AppConfigUrl,
		Namespace:        params.Namespace,
		IsHelmRepository: params.IsHelmRepository(),
		DryRun:           params.DryRun,
	})
	if gitErr != nil {
		return fmt.Errorf("failed to get git clients: %w", gitErr)
	}

	if err := appService.Add(gitClient, gitProvider, params); err != nil {
		return errors.Wrapf(err, "failed to add the app %s", params.Name)
	}

	return nil
}

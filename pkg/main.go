package main

import (
	"context"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/lwj5/ephemeral-launcher/pkg/launcher"
	"github.com/mittwald/go-helm-client"
	"helm.sh/helm/v3/pkg/repo"
	"k8s.io/client-go/rest"
)

const repoName string = "default"

func main() {
	var c launcher.Configuration
	err := envconfig.Process("launcher", &c)
	if err != nil {
		panic(err.Error())
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	options := helmclient.RestConfClientOptions{
		Options: &helmclient.Options{
			Namespace: c.Namespace,
		},
		RestConfig: config,
	}
	client, err := helmclient.NewClientFromRestConf(&options)
	if err != nil {
		panic(err.Error())
	}

	repoEntry := repo.Entry{
		Name: repoName,
		URL:  c.RepoURL,
	}

	err = client.AddOrUpdateChartRepo(repoEntry)
	if err != nil {
		panic(err.Error())
	}

	chartSpec := helmclient.ChartSpec{
		ReleaseName: c.ChartReleaseName,
		ChartName:   repoName + "/" + c.ChartName,
		Namespace:   c.Namespace,
		Version:     c.ChartVersion,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err = client.InstallOrUpgradeChart(ctx, &chartSpec)
	if err != nil {
		panic(err.Error())
	}

}

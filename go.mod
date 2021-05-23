module github.com/lwj5/ephemeral-launcher

go 1.16

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)

require (
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/mittwald/go-helm-client v0.4.2
	helm.sh/helm/v3 v3.5.1
	k8s.io/client-go v0.20.1
)

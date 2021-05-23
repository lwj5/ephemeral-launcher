<!-- PROJECT LOGO -->
<br />
<p align="center">

  <h1 align="center">Ephemeral Launcher</h3>

  <p align="center">
    Deploy your Helm charts directly from a k8s cluster.
    <br />
    <a href="#getting-started"><strong>Get Started »</strong></a>
    <br />
    <br />
    <a href="https://github.com/lwj5/ephemeral-launcher/issues">Report Bug</a>
    ·
    <a href="https://github.com/lwj5/ephemeral-launcher/pulls">Request Feature</a>
  </p>
</p>

<!-- ABOUT THE PROJECT -->
## About The Project

**Disclaimer** Please note that this project is still in beta and many variables are not yet exposed. If you wish to expedite certain processes, feel free to contribute. See [Contributing](#contributing).

Deploying and managing Helm charts shouldn't be difficult. We were trying to find a way to automate the process of installing, upgrading, and removing charts using jobs/pods but there were none available. Hence, here's Ephemeral Launcher, your astromech pod that does all of these.

Key features:
* Automatically downloads the correct chart from the given repo
* Automatically determines an upgrade or install
* Set everything you need via environment variables :smile:

<!-- GETTING STARTED -->
## Getting Started

To install or upgrade a chart simply apply the following YAML file to your k8s cluster replacing `<namespace>`, `<repo-url>`, `<chart-release-name>`, `<chart-name>` and `<chart-name>` with the appropriate values.

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: launcher
  namespace: <namespace>
spec:
  template:
    spec:
      containers:
        - name: launcher
          image: lwj5/ephemeral-launcher
          env:
            - name: LAUNCHER_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
            - name: LAUNCHER_REPO_URL
              value: <repo-url>
            - name: LAUNCHER_CHART_RELEASE_NAME
              value: <chart-release-name>
            - name: LAUNCHER_CHART_NAME
              value: <chart-name>
            - name: LAUNCHER_CHART_VERSION
              value: "<version>"
      restartPolicy: Never
  backoffLimit: 3

```

<!-- USAGE EXAMPLES -->
## Usage

`LAUNCHER_REPO_URL`: the URL of the Helm chart repository

`LAUNCHER_CHART_RELEASE_NAME`: the name to be/was (install/upgrade) given to the installation

`LAUNCHER_CHART_NAME`: the name of Helm chart in the repository (Do not append repository name)

`LAUNCHER_CHART_VERSION`: the version of the chart to install/upgrade to.

<!-- ROADMAP -->
## Roadmap

- Add options to remove the chart
- Add options to export more variables
- Add webhook upon completion

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch
3. Commit your Changes
4. Push to the Branch
5. Open a Pull Request


<!-- LICENSE -->
## License

Distributed under the GPLv3 License. See `LICENSE` for more information.


<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [Go Helm Client](https://github.com/mittwald/go-helm-client)
* [Golang Env](https://github.com/kelseyhightower/envconfig)
* [Best README Template](https://github.com/othneildrew/Best-README-Template)

package launcher

// Configuration to store env variables
type Configuration struct {
	Namespace        string `required:"true" split_words:"true"`
	RepoURL          string `required:"true" envconfig:"repo_url"`
	ChartReleaseName string `required:"true" split_words:"true"`
	ChartName        string `required:"true" split_words:"true"`
	ChartVersion     string `required:"true" split_words:"true"`
}

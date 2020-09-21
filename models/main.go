package models

type K8sContext struct {
	APIVersion string `yaml:"apiVersion"`
	Contexts   []struct {
		Context struct {
			Cluster string `yaml:"cluster"`
			User    string `yaml:"user"`
		} `yaml:"context"`
		Name string `yaml:"name"`
	} `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
}

type K8sCluster struct {
	Name     string
	Cluster  string
	Location string
	Project  string
	Cloud    string
}

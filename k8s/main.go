package k8s

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/BESTSELLER/gke-sacle2/models"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

func Getk8sContext(clusterName string) models.K8sCluster {
	content, err := ioutil.ReadFile(getConfigFile())
	if err != nil {
		fmt.Println(err)
	}

	var yamlData models.K8sContext

	err = yaml.Unmarshal([]byte(content), &yamlData)
	if err != nil {
		fmt.Println(err)
	}

	var cluster models.K8sCluster
	contexts := yamlData.Contexts

	for _, v := range contexts {
		if clusterName == v.Name {
			c := strings.Split(v.Context.Cluster, "_")
			cluster.Cloud = c[0]
			cluster.Project = c[1]
			cluster.Location = c[2]
			cluster.Name = c[3]
		}
	}

	if cluster.Name == "" {
		fmt.Printf("Was unabled to cluster %s in kube/config \n", clusterName)
		os.Exit(1)
	}

	return cluster
}

func getConfigFile() (configFile string) {
	var k8sConfigFile string
	var err error
	os := runtime.GOOS

	switch os {
	case "darwin":
		k8sConfigFile, err = homedir.Expand("~/.kube/config")

	}
	if err != nil {
		fmt.Println(err)
	}

	return k8sConfigFile
}

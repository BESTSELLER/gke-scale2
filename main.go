package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BESTSELLER/gke-sacle2/gke"
	"github.com/BESTSELLER/gke-sacle2/k8s"
)

func main() {
	k8scontext := flag.String("context", "", "kubernetes context name (Required) - kubectl config get-contexts")
	nodeCount := flag.Int("nodecount", 1, "The number of node you wish to scale to")

	flag.Parse()
	if *k8scontext == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	k8scluster := k8s.Getk8sContext(*k8scontext)

	fmt.Println(k8scluster.Cluster)
	project := k8scluster.Project
	location := k8scluster.Location
	cluster := k8scluster.Name

	clusterID := "projects/" + project + "/locations/" + location + "/clusters/" + cluster

	nodePools, err := gke.ListNodePools(clusterID)
	if err != nil {
		fmt.Println(err)
	}

	for _, np := range nodePools {
		fmt.Println("now scaling node pool:", np)
		gke.ScaleNodePool(int32(*nodeCount), clusterID, np)
	}

}

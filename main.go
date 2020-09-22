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

	// fmt.Println(k8scluster.Cluster)
	project := k8scluster.Project
	location := k8scluster.Location
	cluster := k8scluster.Name

	clusterID := "projects/" + project + "/locations/" + location + "/clusters/" + cluster

	nodePools, err := gke.ListNodePools(clusterID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("You are about to scale the GKE cluster: %s\n", cluster)
	for _, np := range nodePools {
		fmt.Printf("The node pool %s will scale to %v node/nodes per zone.\n", np, *nodeCount)
	}

	fmt.Printf("Type 'yes' to continue ")
	var confirmation string
	fmt.Scanln(&confirmation)

	if confirmation == "yes" {
		for _, np := range nodePools {
			fmt.Println("Now scaling node pool:", np)
			gke.ScaleNodePool(int32(*nodeCount), clusterID, np)
		}
	} else {
		fmt.Println("Scaling has been canceled")
		os.Exit(1)
	}

}

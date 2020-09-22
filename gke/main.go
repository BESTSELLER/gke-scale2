package gke

import (
	"context"
	"fmt"

	container "cloud.google.com/go/container/apiv1"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

func ScaleNodePool(nodeCount int32, clusterID string, nodePool string) {

	ctx := context.Background()
	c, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		fmt.Println(err)
	}

	req := &containerpb.SetNodePoolSizeRequest{
		NodeCount: nodeCount,
		Name:      clusterID + "/nodePools/" + nodePool,
	}
	resp, err := c.SetNodePoolSize(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func ListNodePools(clusterID string) (nodePools []string, err error) {
	ctx := context.Background()
	// c, err := google.DefaultClient(ctx, container.ClusterManagerClient)
	c, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return nil, err
	}

	// fmt.Println(clusterID)
	req := &containerpb.ListNodePoolsRequest{
		Parent: clusterID,
	}
	resp, err := c.ListNodePools(ctx, req)
	if err != nil {
		return nil, err
	}

	var nps []string

	for _, v := range resp.NodePools {
		nps = append(nps, v.Name)
	}
	return nps, nil
}

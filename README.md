# gke-scale2

Allow you to quickly scale the node pool within GKE.

## Usage.
Get the name of kubernetes context by using the following command. 
`kubectx config get-contexts` or `kubectx`

Now that you have the context name for the cluster you can scale the GKE deployment to the desired size.

`./gke-scale2 -context gke_project-name_here_location-here_cluster-name-here`

The below will scale the nodepool count to 2 in each zone

`./gke-scale2 -context gke_project-name_here_location-here_cluster-name-here -nodecount 2`
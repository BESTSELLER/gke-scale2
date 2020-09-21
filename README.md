# gke-scale2

Allow you to quickly scale the node pool within GKE.

## Usage.
Get the name of kubernetes context the below command 
`kubectx config get-contexts` or `kubectx`

Now that you have the context name for the cluster you can scale the GKE deployment to the desired size
`./gke-scale2 -context gke_project-name_here_location-here_cluster-name-here`

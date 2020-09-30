# gke-scale2

Allow you to quickly scale the node pool within GKE.

## Prerequirement
The following application needs to be installed before execution of gke-scale2.
* kubectl
* gcloud

## Installation

* homebrew
  
  If you work on MacOS, you can install gke-scale2 via [brew](https://brew.sh/)
  ```
  brew tap BESTSELLER/homebrew-tap
  brew install gke-scale2
  ```
  
* scoop
  
  If you work on Window platform, you can install gke-scale2 via [scoop](https://scoop.sh/)
  ```
  scoop bucket add bestseller https://github.com/BESTSELLER/scoop-bucket
  scoop install gke-scale2
  ```

* binary
  
  You can download the latest version from the [release page](https://github.com/BESTSELLER/gke-scale2/releases/latest)

* soure code
  
  You can install from source code
  ```
  go get -u -v github.com/BESTSELLER/gke-scale2
  
  go install -a -v github.com/BESTSELLER/gke-scale2
  ```

## Usage.
Make sure you are logged into gcp via the below command

` gcloud auth application-default login`

Or the environment variable GOOGLE_APPLICATION_CREDENTIALS is set

`export GOOGLE_APPLICATION_CREDENTIALS="/path/to/key.json"`

<br/>

Get the name of kubernetes context by using the following command. 
`kubectl config get-contexts` or `kubectx`

Now that you have the context name for the cluster you can scale the GKE deployment to the desired size.

`gke-scale2 -context gke_project-name_here_location-here_cluster-name-here`

The below will scale the nodepool count to 2 in each zone

`gke-scale2 -context gke_project-name_here_location-here_cluster-name-here -nodecount 2`

#  	Tracking syscall and Function Latency in your k8s Cluster with eBPF


## Build the docker images
build the app image `docker build . -t repo/example-tool-name` 

build the init image `docker build -f init/Dockerfile  -t repo/bcc-linux-headers`

push your images ex: `docker push repo/example-tool-name`

### Create the daemonset 
If you are using your own container repo and images make sure to update the containers in the deployment file.

`kubectl apply -f deployment.yaml`


## questions or improvements 
email: matt@containiq.com 

## Thanks 
* [BPF Compiler Collection (BCC)](https://github.com/iovisor/bcc)
* [kubectl-trace](https://github.com/iovisor/kubectl-trace)


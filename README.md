# Brgo-CD
To learn more about the inner workings of [ArgoCD](https://argoproj.github.io/cd) why not try to re-implement a really basic version of it?
This is what Brgo-CD is about.


# How to use

```
task init
task build_base_ubuntu
task build_base_git
task build_base_kubectl
task build_gitserver
task build_applyserver
task gitserver
task applyserver
```

1. This will clone the sample manifests repo
1. Render the manifests
1. Apply them to the cluster specified as context in the kubeconfig file

Remarks:

You need to have your Kubeconfig at `$(pwd)/kubeconfig.yaml`
 

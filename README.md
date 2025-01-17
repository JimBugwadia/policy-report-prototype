# policy-report-prototype
Building a prototype of Policy Report Generator. It aims to run a CIS benchmark check with a tool called [kube-bench](https://github.com/aquasecurity/kube-bench) and produce a policy report based on the Custom Resource Definition accordingly.

## Running

**Prerequisites**: 
* Since the policy-report-prototype uses `apps/v1` deployments, the Kubernetes cluster version should be greater than 1.9.
* To run the Kubernetes cluster locally, tools like [kind](https://kind.sigs.k8s.io/) or [minikube](https://minikube.sigs.k8s.io/docs/start/) can be used. In our case, we will be going with [kind](https://kind.sigs.k8s.io/). You can follow the links if kind or minikube aren't installed on your local machine.
* We will also need `kube-bench` binary with its `cfg`. 

This can be accomplished by: 

Step 1: `git clone https://github.com/mritunjaysharma394/policy-report-prototype.git`

Step 2: `cd policy-report-prototype`

Step 3: ```docker run --rm -v `pwd`:/host aquasec/kube-bench:latest install```

### Next Steps

```sh
# 1. create a local Kubernetes cluster
kind create cluster
    OR
minikube start

# 2. create a CustomResourceDefinition
kubectl create -f crd/wgpolicyk8s.io_policyreports.yaml

# 3. Build
go build

# 4. Create policy report using
./policyreport -policyName=cis-sample-policy-report -namespace=default

# 5. check policyreports created through the custom resource
kubectl get policyreports
```
**Notes**: 
* Flags `-policyName` and `-namespace`are user configurable and can be changed by changing the variable on the right hand side. 
* In order to generate policy report in the form of YAML, step 5 can be written as `kubectl get policyreports -o yaml > res.yaml` which will generate it as `res.yaml` in this case.

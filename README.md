# pid2Pod

Display (Namespace, Pod, Container, Primary PID) from a host PID, fails if the target process is running on host

## Pre-requisites

User MUST be able to run `crictl` on the host.

## Install

```shell
curl -Lo ./pod2pid https://github.com/k8s-school/pid2pod/releases/download/v0.0.1/pid2pod-linux-amd64 
chmod +x ./pod2pid
mv ./pod2pid /some-dir-in-your-PATH/pod2pid
```

## Examples

```shell
./pod2pid 1525
NAMESPACE     POD                 CONTAINER     PRIMARY PID
kube-system   calico-node-6kt29   calico-node   1284
```

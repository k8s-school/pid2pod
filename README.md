# pid2Pod

Display (Namespace, Pod, Container, Primary PID) from a host PID, fails if the target process is running on host

## Pre-requisites

User MUST be able to run `crictl` on the host.

## Install

```shell
curl -Lo ./pid2pod https://github.com/k8s-school/pid2pod/releases/download/v0.0.1/pid2pod-linux-amd64 
chmod +x ./pid2pod
mv ./pid2pod /some-dir-in-your-PATH/pid2pod
```

## Examples

```shell
./pid2pod 1525
NAMESPACE     POD                 CONTAINER     PRIMARY PID
kube-system   calico-node-6kt29   calico-node   1284
```

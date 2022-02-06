# pid2Pod

Display (Namespace, Pod, Container, Primary PID) from a host PID, fails if the target process is running on host

## Pre-requisites

User MUST be able to run `crictl` on the host.

## Examples

```shell
./pod2pid 1525
NAMESPACE     POD                 CONTAINER     PRIMARY PID
kube-system   calico-node-6kt29   calico-node   1284
```

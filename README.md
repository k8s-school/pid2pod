# ✨ pid2pod

Retrieve and display Pod and Container from a host PID, fails if the target process is not running inside a container

## 🏁 Pre-requisites

User MUST be able to run `crictl` on the host.

## 💻 Installation

```shell
curl -Lo ./pid2pod https://github.com/k8s-school/pid2pod/releases/download/v0.0.1/pid2pod-linux-amd64
chmod +x ./pid2pod
mv ./pid2pod /some-dir-in-your-PATH/pid2pod
```

## 👁️ Examples

```shell
./pid2pod 1525
NAMESPACE     POD                 CONTAINER     PRIMARY PID
kube-system   calico-node-6kt29   calico-node   1284

./pid2pod 2028
NAMESPACE            POD                                       CONTAINER                PRIMARY PID
local-path-storage   local-path-provisioner-547f784dff-rl446   local-path-provisioner   2028

./pid2pod $(pgrep kubelet)
process 182 is running on host

./pid2pod $RANDOM
error: no process with pid 18489
```
## TODO

Retrieve pod related to `pause` process, relationship between those two is tracked here:
```shell
crictl inspectp  --output go-template --template "{{.info.pid}}" 20bdea27e88d3
```

# pid2Pod

Returns (Namespace, Pod, Container) from a host PID, fails if the target process is running on host

## Pre-requisites

Use must be able to run `crictl` on the host.

## Examples

```shell
pod2pid -p 1
```

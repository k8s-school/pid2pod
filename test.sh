#!/bin/bash

set -euxo pipefail

go build .
docker cp pid2pod kind-control-plane:/
docker exec -it kind-control-plane /pid2pod 2

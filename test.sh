#!/bin/bash

set -euo pipefail

go build .
docker cp pid2pod kind-control-plane:/

for program in coredns kube-proxy bird6 kubelet RANDOM1 RANDOM2 RANDOM3; do
  PID=$(docker exec -- kind-control-plane sh -c "pgrep $program || echo '$RANDOM'")
  echo "Looking for $PID ($program)"
  docker exec kind-control-plane /pid2pod "$PID" || echo "error returned"
  echo
done


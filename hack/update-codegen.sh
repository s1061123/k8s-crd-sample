#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
PKG_NAME="github.com/s1061123/k8s-crd-sample/pkg"
GROUP_NAME="k8s.cni.cncf.io"

bash vendor/k8s.io/code-generator/generate-groups.sh all \
  ${PKG_NAME}/client ${PKG_NAME}/apis \ ${GROUP_NAME}:v1 \
  --go-header-file ${SCRIPT_ROOT}/hack/custom-boilerplate.go.txt

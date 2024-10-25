#!/bin/bash

# 获取最新的 KubeVirt 版本号
export KUBEVIRT_VERSION=$(curl -s https://storage.googleapis.com/kubevirt-prow/release/kubevirt/kubevirt/stable.txt)
echo "Latest KubeVirt version: $KUBEVIRT_VERSION"

# 获取最新的 CDI 版本号
export CDI_VERSION=$(basename $(curl -s -w %{redirect_url} https://github.com/kubevirt/containerized-data-importer/releases/latest))
echo "Latest CDI version: $CDI_VERSION"

# 1. 安装 KubeVirt CRD 和 Operator
echo "Installing KubeVirt CRDs and Operator..."
kubectl create namespace kubevirt
kubectl apply -f https://github.com/kubevirt/kubevirt/releases/download/${KUBEVIRT_VERSION}/kubevirt-operator.yaml

# 部署 KubeVirt CRDs
kubectl apply -f https://github.com/kubevirt/kubevirt/releases/download/${KUBEVIRT_VERSION}/kubevirt-cr.yaml

# 2. 安装 CDI (Containerized Data Importer)
echo "Installing CDI..."
kubectl create namespace cdi
kubectl apply -f https://github.com/kubevirt/containerized-data-importer/releases/download/${CDI_VERSION}/cdi-operator.yaml

# 部署 CDI CRDs
kubectl apply -f https://github.com/kubevirt/containerized-data-importer/releases/download/${CDI_VERSION}/cdi-cr.yaml

# 3. 等待 KubeVirt 和 CDI 安装完成
echo "Waiting for KubeVirt and CDI components to be ready..."
kubectl wait --for=condition=Available -n kubevirt kv kubevirt --timeout=10m
kubectl wait --for=condition=Available -n cdi cdi cdi --timeout=10m

# 检查安装状态
echo "Verifying KubeVirt installation..."
kubectl get all -n kubevirt

echo "Verifying CDI installation..."
kubectl get all -n cdi

# 4. 安装 virtctl 客户端工具 (可选)
echo "Installing virtctl client tool..."
curl -LO https://github.com/kubevirt/kubevirt/releases/download/${KUBEVIRT_VERSION}/virtctl-${KUBEVIRT_VERSION}-linux-amd64
chmod +x virtctl-${KUBEVIRT_VERSION}-linux-amd64
sudo mv virtctl-${KUBEVIRT_VERSION}-linux-amd64 /usr/local/bin/virtctl

echo "KubeVirt, CDI, and virtctl installation completed successfully!"
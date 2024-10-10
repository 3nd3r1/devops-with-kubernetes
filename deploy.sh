export SOPS_AGE_KEY_FILE=./key.txt

# Deploy prometheus
kubectl create namespace prometheus --dry-run=client -o yaml | kubectl apply -f -
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add stable https://charts.helm.sh/stable
helm install prometheus-community/kube-prometheus-stack --generate-name --namespace prometheus

# Deploy argo-rollouts
kubectl create namespace argo-rollouts --dry-run=client -o yaml | kubectl apply -f -
kubectl apply -n argo-rollouts -f https://github.com/argoproj/argo-rollouts/releases/latest/download/install.yaml

# Deploy argocd
kubectl create namespace argocd --dry-run=client -o yaml | kubectl apply -f -
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Deploy ping-pong and log-output
#kubectl create namespace applications --dry-run=client -o yaml | kubectl apply -f -
#kubectl config set-context --current --namespace=applications
#sops --decrypt ./ping-pong/manifests/secrets/secret.enc.yaml | kubectl apply -f -
#kubectl apply -k ./ping-pong/
#kubectl apply -k ./log-output/
#kubectl wait --for=condition=ready pod -l project=ping-pong --timeout=300s
#kubectl wait --for=condition=ready pod -l project=log-output --timeout=300s

# Setup todo-project
kubectl create namespace production --dry-run=client -o yaml | kubectl apply -f -
kubectl create namespace staging --dry-run=client -o yaml | kubectl apply -f -
sops --decrypt ./todo-project/manifests/secrets/secret.enc.yaml | kubectl apply -n production -f -
sops --decrypt ./todo-project/manifests/secrets/secret.enc.yaml | kubectl apply -n staging -f -
helm upgrade --install --set commonLabels.project=todo-project --set auth.enabled=false prod-todo-project-nats oci://registry-1.docker.io/bitnamicharts/nats --namespace=production
helm upgrade --install --set commonLabels.project=todo-project --set auth.enabled=false staging-todo-project-nats oci://registry-1.docker.io/bitnamicharts/nats --namespace=staging
kubectl apply -n argocd -f ./todo-project/manifests/application.yaml

echo "Waiting for argocd server:"
kubectl -n argocd wait --for=condition=ready pod -l app.kubernetes.io/name=argocd-server --timeout=300s
echo "Argocd admin password:"
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

export SOPS_AGE_KEY_FILE=./key.txt

kubectl create namespace applications
sops --decrypt ./ping-pong/manifests/secrets/secret.enc.yaml | kubectl apply -f -
kubectl apply -f ./ping-pong/manifests/deployment.yaml
kubectl apply -f ./ping-pong/manifests/service.yaml
kubectl apply -f ./ping-pong/manifests/statefulset.yaml
#kubectl apply -f ./log-output/manifests/

#kubectl create namespace project
#sops --decrypt ./todo-project/manifests/secrets/secret.enc.yaml | kubectl apply -f -
#kubectl apply -f ./pvs/todo-project/
#kubectl apply -f ./todo-project/manifests/

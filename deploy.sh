kubectl create namespace applications
kubectl apply -f ./ping-pong/manifests/
kubectl apply -f ./log-output/manifests/

#kubectl create namespace project
#kubectl apply -f ./pvs/todo-project/
#kubectl apply -f ./todo-project/manifests/

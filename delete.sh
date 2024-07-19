kubectl delete service log-output-svc
kubectl delete service ping-pong-svc
kubectl delete service todo-project-backend-svc
kubectl delete service todo-project-frontend-svc
kubectl delete service todo-project-imagenator-svc

kubectl delete ingress log-output-ingress
kubectl delete ingress todo-project-ingress

kubectl delete deployment log-output-dep
kubectl delete deployment ping-pong-dep
kubectl delete deployment todo-project-backend-dep
kubectl delete deployment todo-project-frontend-dep
kubectl delete deployment todo-project-imagenator-dep

kubectl delete pvc todo-project-data-pvc
kubectl delete pv todo-project-data-pv

docker build -t 3nd3r1/log-output:latest -f ./log-output/build/Dockerfile ./log-output
docker build -t 3nd3r1/ping-pong:latest -f ./ping-pong/build/Dockerfile ./ping-pong
#docker build -t 3nd3r1/todo-project-backend:latest -f ./todo-project/backend/build/Dockerfile ./todo-project/backend
#docker build -t 3nd3r1/todo-project-frontend:latest -f ./todo-project/frontend/build/Dockerfile ./todo-project/frontend
#docker build -t 3nd3r1/todo-project-imagenator:latest -f ./todo-project/imagenator/build/Dockerfile ./todo-project/imagenator

docker push 3nd3r1/log-output:latest
docker push 3nd3r1/ping-pong:latest
#docker push 3nd3r1/todo-project-backend:latest
#docker push 3nd3r1/todo-project-frontend:latest
#docker push 3nd3r1/todo-project-imagenator:latest

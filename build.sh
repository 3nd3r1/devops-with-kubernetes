docker build -t 3nd3r1/log-output-reader:latest -f ./log-output/build/reader/Dockerfile ./log-output
docker build -t 3nd3r1/log-output-writer:latest -f ./log-output/build/writer/Dockerfile ./log-output
docker build -t 3nd3r1/ping-pong:latest -f ./ping-pong/build/Dockerfile ./ping-pong
#docker build -t 3nd3r1/todo-project-backend:latest -f ./todo-project/backend/build/Dockerfile ./todo-project/backend

docker push 3nd3r1/log-output-reader:latest
docker push 3nd3r1/log-output-writer:latest
docker push 3nd3r1/ping-pong:latest
#docker push 3nd3r1/todo-project-backend:latest

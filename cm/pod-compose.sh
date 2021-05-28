#! /bin/bash

podman rm -f go-api go-index
podman pod rm gofiles
podman pod create gofiles --publish 3000:80

# podman run --detach \
#   --pod gofile \
#   -v "$PWD/nginx.conf:/etc/nginx/nginx.conf:z" \
#   -v "$PWD/results:/usr/share/nginx/html/results:z" \ 
#   -p 8989:80 \
#   --name nginx
#   nginx

# podman run --detach \
#   --pod gofile \
#   -v "$PWD/nginx.conf:/etc/nginx/nginx.conf:z" \
#   -v "$PWD/results:/usr/share/nginx/html/results:z" \ 
#   -p 8989:80 \
#   --name go-index
#   go-index
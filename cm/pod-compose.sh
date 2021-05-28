#! /bin/bash

podman rm -f go-api go-index
podman pod rm gofiles
podman pod create --name=gofiles --publish 3000:80

podman run --detach \
  --pod gofiles \
  -v "$PWD/nginx/nginx.conf:/etc/nginx/nginx.conf:z" \
  -v "$PWD/results:/usr/share/nginx/html/results:z" \
  --name go-index \
  nginx:1.16


# podman run --detach \
#   --pod gofile \
#   -v "$PWD/nginx.conf:/etc/nginx/nginx.conf:z" \
#   -v "$PWD/results:/usr/share/nginx/html/results:z" \ 
#   -p 8989:80 \
#   --name go-index
#   go-index


podman run --detach \
  --pod gofiles \
  --name go-api \
  -v "$PWD/results:/app/results:z" \
  localhost/go-api
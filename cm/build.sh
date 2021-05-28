#! /bin/bash

podman build --tag go-api --file ../api/backend.containerfile
podman build --tag go-index --file ../nginx/nginx.containerfile
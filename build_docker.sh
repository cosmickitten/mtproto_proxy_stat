#!/bin/sh
docker build -t $DOCKER_REPO . -
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push $DOCKER_REPO
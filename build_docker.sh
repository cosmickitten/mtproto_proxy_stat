#!/bin/sh
docker build --rm -t "$DOCKER_USERNAME"/"$DOCKER_REPO"  . -
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push "$DOCKER_USERNAME"/"$DOCKER_REPO"
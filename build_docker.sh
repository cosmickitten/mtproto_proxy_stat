#!/bin/sh
docker build --rm .
echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
docker push $DREPO
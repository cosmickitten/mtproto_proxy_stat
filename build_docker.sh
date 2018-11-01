#!/bin/sh
docker build --rm .
echo "jkug9uh4y" | docker login -u "cosmickitten" --password-stdin
docker push cosmickitten/tgproxyservice
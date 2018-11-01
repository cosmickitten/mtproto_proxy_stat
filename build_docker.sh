#!/bin/sh
docker build --rm -t cosmickitten/tgproxyservice  . -
echo "jkug9uh4y" | docker login -u "cosmickitten" --password-stdin
docker push cosmickitten/tgproxyservice
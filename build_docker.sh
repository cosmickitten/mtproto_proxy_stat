#!/bin/sh
docker build --rm -t cosmickitten/tgproxyservice  . -
echo "jkug9uh4y" | docker login -username cosmickitten --password-stdin
docker push cosmickitten/tgproxyservice
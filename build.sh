#!/bin/sh

docker build --build-arg http_proxy=$http_proxy --build-arg https_proxy=$https_proxy -t brisk-cli . && \
docker create --name brisk-cli brisk-cli && \
docker cp brisk-cli:/root/brisk . && \
docker rm -f brisk-cli
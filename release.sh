#!/bin/sh

docker build --build-arg http_proxy=$http_proxy --build-arg https_proxy=$https_proxy -t brisk-cli . -f Dockerfile.release && \
docker create --name brisk-cli brisk-cli && \
docker cp brisk-cli:/root/brisk . && \
docker cp brisk-cli:/root/brisk-darwin . && \
docker cp brisk-cli:/root/brisk-armhf . && \
docker cp brisk-cli:/root/brisk-arm64 . && \
docker cp brisk-cli:/root/brisk.exe . && \
docker rm -f brisk-cli
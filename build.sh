#!/bin/sh

docker build -t brisk-cli . && \
docker create --name brisk-cli brisk-cli && \
docker cp openbrisk_cli:/root/brisk-cli . && \
docker rm -f brisk-cli
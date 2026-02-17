#!/usr/bin/env bash
set -e
ARGS=""
# build arg image2 is set by Skaffold to be the image built for app2
if [[ "${PUSH_IMAGE}" == "true" ]]; then
ARGS="--push"
else
ARGS="--load"
fi
docker buildx build $ARGS -t "$IMAGE" --build-arg image2 .

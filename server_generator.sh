#!/bin/bash

docker run --rm -v "${PWD}/schema:/local" \
    openapitools/openapi-generator-cli generate \
    -i ./local/openapi.yaml \
    -g go-server \
    -o /local/generated
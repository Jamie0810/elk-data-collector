#!/usr/bin/env bash
. ./build/docker.env

VERSION="DEV_2020062502"
IMAGE_NAME="${TEAM_NAME}-${SERVICE_NAME}-lambda"
docker build . -f ./build/Dockerfile -t "${IMAGE_NAME}:${VERSION}" -t "${IMAGE_NAME}:latest"

# aws ecr get-login-password --region "${REGION}" | docker login --username AWS --password-stdin "${ACCOUNT_ID}".dkr.ecr.us-east-1.amazonaws.com    
# docker tag  "${IMAGE_NAME}":latest "${ACCOUNT_ID}".dkr.ecr.us-east-1.amazonaws.com/"${IMAGE_NAME}":latest
# docker push "${ACCOUNT_ID}".dkr.ecr.us-east-1.amazonaws.com/hello-world:latest        

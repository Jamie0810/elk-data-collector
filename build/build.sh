#!/usr/bin/env bash
. ./build/docker.env

CURRENT_TIME=$(date +"%Y%m%d%H%M")
VERSION="${STAGE}_${CURRENT_TIME}"
IMAGE_NAME="${TEAM_NAME}-${STAGE}-${BUSINESS_DOMAIN}-${SERVICE_NAME}"
docker build . -f ./build/Dockerfile -t "${IMAGE_NAME}:${VERSION}" -t "${IMAGE_NAME}:latest"

aws ecr get-login-password --region "${REGION}" | docker login --username AWS --password-stdin "${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com"
docker tag  "${IMAGE_NAME}:latest" "${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com/${IMAGE_NAME}:latest"

docker push "${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com/${IMAGE_NAME}:${VERSION}"
docker push "${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com/${IMAGE_NAME}:latest"

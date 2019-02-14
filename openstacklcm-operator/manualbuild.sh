#!/bin/bash -x
export COMPONENT=openstacklcm-operator
export VERSION=0.0.1
export DHUBREPO="kubekit99/$COMPONENT-dev"
export DOCKER_NAMESPACE="kubekit99"
export DOCKER_USERNAME="kubekitdevops"
export DOCKER_PASSWORD=$KUBEKITDEVOPSPWD

#echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
operator-sdk build $DHUBREPO:v$VERSION

docker tag $DHUBREPO:v$VERSION $DHUBREPO:latest
#docker push $DHUBREPO

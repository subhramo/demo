#!/bin/bash -x

aws ecr get-login-password --region ap-southeast-2 | docker login --username AWS --password-stdin 825030697311.dkr.ecr.ap-southeast-2.amazonaws.com

docker tag demo:latest 825030697311.dkr.ecr.ap-southeast-2.amazonaws.com/demo:latest

docker push 825030697311.dkr.ecr.ap-southeast-2.amazonaws.com/demo:latest

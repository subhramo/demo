#!/bin/bash -x

aws ecr get-login-password --region ap-southeast-2 | docker login --username AWS --password-stdin xxxxxxxxx.dkr.ecr.ap-southeast-2.amazonaws.com

docker tag demo:latest xxxxxxxxx.dkr.ecr.ap-southeast-2.amazonaws.com/demo:latest

docker push xxxxxxxxxxx.dkr.ecr.ap-southeast-2.amazonaws.com/demo:latest

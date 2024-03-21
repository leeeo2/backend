#!/bin/bash

docker rm -f backend
docker run \
	-d -p 8888:8888 \
     -v ./etc/backend.yml:/etc/backend/backend.yml \
	--name backend \
	--restart always \
	$1

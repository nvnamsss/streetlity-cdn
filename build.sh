#/bin/bash

docker image rm -f driver_container
docker build . -t driver_container


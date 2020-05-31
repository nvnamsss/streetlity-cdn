#/bin/bash
name=$1

echo "STOP RUNNING API CONTAINER"
docker stop -t 30 ${name}_driver_container 
docker rm -f ${name}_driver_container

echo "DONE STOPPING"

docker run --name ${name}_driver_container -d\
            --network common-net \
            --restart always \
            -p 9003:9003 \
            driver_container

docker cp config.json ${name}_api_container:/server/config/config.json    

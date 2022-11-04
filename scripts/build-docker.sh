#!/bin/bash
cd ..

docker build -t cosmonaut-cat/backend_event_handler:latest  -f cmd/event_handler/Dockerfile .   
docker build -t cosmonaut-cat/backend_front_api:latest  -f cmd/front_api/Dockerfile .   
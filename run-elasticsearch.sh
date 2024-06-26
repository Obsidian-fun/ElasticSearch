#! /bin/bash

docker rm -rf elasticsearch

docker run -d --name elasticsearch -p 9200:9200 -e discovery.type=single-node \
    -v elasticsearch:/usr/share/elasticsearch/data \
    docker.elastic.co/elasticsearch/elasticsearch:7.5.2

docker ps

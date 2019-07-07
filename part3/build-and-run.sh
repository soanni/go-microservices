#!/bin/bash
docker-compose down
cd consignment-cli
make
cd ../consignment-service
make
cd ../vessel-service
make
cd ..
docker-compose up -d consignment-service vessel-service datastore
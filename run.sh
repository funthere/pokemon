#!/bin/bash

replicas=10
base_port=8090

cat << EOF > docker-compose.generated.yml
version: '3.8'

services:
EOF

for i in $(seq 1 $replicas); do
  port=$((base_port + i - 1))
  type="type_$i"
  cat << EOF >> docker-compose.generated.yml
  service-a-$i:
    image: pokemon_service-a
    ports:
      - "$port:8081"
    expose:
      - "8081"
    environment:
      - SENSOR_TYPE=$type
EOF
done

docker-compose -f docker-compose.generated.yml up

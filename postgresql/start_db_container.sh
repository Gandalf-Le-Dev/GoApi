#!/bin/bash

container_name="postgres_db"

if ! docker ps -q -f name="$container_name" | grep -q . ; then
    echo "Starting Docker Compose services..."
    docker-compose up -d
    echo "Services started."
else
    echo "$container_name is already running."
fi

read -p "Press [Enter] key to continue..."

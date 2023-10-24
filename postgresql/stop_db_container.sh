#!/bin/bash

container_name="postgres_db"

if ! docker ps -q -f name="$container_name" | grep -q . ; then
    echo "$container_name is not running."
else
    echo "Stopping Docker Compose services..."
    docker-compose down
    echo "Services stopped."
fi

read -p "Press [Enter] key to continue..."

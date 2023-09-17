#!/bin/bash

#Start the Docker service
sudo systemctl start docker

# Run Docker Compose commands
cd /
cd vagrant/docker
docker-compose up

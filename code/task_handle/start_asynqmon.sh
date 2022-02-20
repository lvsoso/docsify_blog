#! /bin/bash

docker run --rm -d\
    --name asynqmon \
    --network host \
    hibiken/asynqmon
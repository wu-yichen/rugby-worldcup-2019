#!/bin/bash

set -x

error_message=
function setup() {

    error_message=$(fly -t ps set-pipeline -n -c pipeline.yml -p ps --load-vars-from .credential.yml 2>&1)
}

setup
if [[ $? -gt 0 ]]; then
    if [[ $error_message =~ "is the targeted Concourse running" ]]; then
        echo "concourse is not started yet, starting now..."
        docker-compose up
        fly -t ps set-pipeline -c pipeline.yml -p ps --load-vars-from .credential.yml
    fi
fi
#!/bin/bash



function ___main {

    $(which docker) build --tag ${DOCKER_IMAGE_NAME} --file $(pwd)/Dockerfile .
    if [ $? != 0 ] ; then
        return 1
    fi

    return 0
}



___main "${@}"

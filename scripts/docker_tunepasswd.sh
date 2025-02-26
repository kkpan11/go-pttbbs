#!/bin/bash
# Aligning the .PASSWDS with the expected MAX_USER from c-pttbbs
#
# https://github.com/ptt/pttbbs/blob/master/util/tunepasswd.c

if [ "$#" != "2" ]; then
    echo "usage: docker_initbbs.sh [BBSHOME] [GOPTTBBS_IMAGE]"
    exit 255
fi

BBSHOME=$1
GOPTTBBS_IMAGE=$2
echo "BBSHOME: ${BBSHOME} GOPTTBBS_IMAGE: ${GOPTTBBS_IMAGE}"

project=go-pttbbs-tunepasswd

docker container stop ${project}
docker container rm ${project}
docker run --rm --name ${project} -v ${BBSHOME}:/home/bbs ${GOPTTBBS_IMAGE} /opt/bbs/bin/tunepasswd

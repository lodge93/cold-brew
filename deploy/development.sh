#!/bin/bash

set -e

DEPLOY_DIR="/home/pi/cold-brew"
SKIP_FRONTEND_ARTIFACTS=false
REMOTE_HOST="cold-brew.dev"

while getopts 'sd:h:' flag; do
    case "${flag}" in
        s) SKIP_FRONTEND_ARTIFACTS=true ;;
        h) REMOTE_HOST="${OPTARG}" ;;
    esac
done

echo "Building local project"
GOOS=linux GOARCH=arm go build -o build/cold-brew
if [ "$SKIP_FRONTEND_ARTIFACTS" = false ]
then
    npm run build
    tar -zcvf build/frontend.tar.gz assets/dist
fi

echo "Creating remote deploy directory"
ssh pi@${REMOTE_HOST} "
    mkdir -p ${DEPLOY_DIR}
"

echo "Adding remote service configuration"
scp deploy/cold-brew.service pi@${REMOTE_HOST}:${DEPLOY_DIR}
ssh pi@${REMOTE_HOST} "
    sudo mv /${DEPLOY_DIR}/cold-brew.service /etc/systemd/system/cold-brew.service
    sudo systemctl daemon-reload
"

echo "Adding remote configuration file"
scp deploy/config.yml pi@${REMOTE_HOST}:${DEPLOY_DIR}

echo "Stopping remote application"
ssh pi@${REMOTE_HOST} "
    sudo systemctl stop cold-brew
"

echo "Copying built artifacts to remote host"
scp build/cold-brew pi@${REMOTE_HOST}:${DEPLOY_DIR}
if [ "$SKIP_FRONTEND_ARTIFACTS" = false ]
then
    scp build/frontend.tar.gz pi@${REMOTE_HOST}:${DEPLOY_DIR}
    ssh pi@${REMOTE_HOST} "
        cd ${DEPLOY_DIR}
        tar -xvzf ${DEPLOY_DIR}/frontend.tar.gz
        rm ${DEPLOY_DIR}/frontend.tar.gz
    "
fi

echo "Starting remote application"
ssh pi@${REMOTE_HOST} "
    sudo systemctl start cold-brew
"
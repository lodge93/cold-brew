#!/bin/bash
# Stick this in /usr/local/bin/updater and run "updater 28" to install build 28.

set -e

VERSION=$1

mkdir -p /tmp/${VERSION}/frontend
mkdir -p /tmp/${VERSION}/backend

curl -Lo /tmp/${VERSION}/frontend/frontend-${VERSION}.zip https://storage.googleapis.com/project-build-storage/cold-brew/snapshots/${VERSION}/frontend/frontend-${VERSION}.zip
curl -Lo /tmp/${VERSION}/backend/cold-brew-server_0.0.1-${VERSION}_armhf.deb https://storage.googleapis.com/project-build-storage/cold-brew/snapshots/${VERSION}/backend/cold-brew-server_0.0.1-${VERSION}_armhf.deb

sudo dpkg -i /tmp/${VERSION}/backend/cold-brew-server_0.0.1-${VERSION}_armhf.deb
sudo rm -rf /root/assets
sudo unzip /tmp/${VERSION}/frontend/frontend-${VERSION}.zip -d /root

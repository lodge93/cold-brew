# cold-brew
[![Build Status](https://travis-ci.org/lodge93/cold-brew.svg?branch=master)](https://travis-ci.org/lodge93/cold-brew)
[![Go Report Card](https://goreportcard.com/badge/github.com/lodge93/cold-brew)](https://goreportcard.com/report/github.com/lodge93/cold-brew)
[![GoDoc](https://godoc.org/github.com/lodge93/cold-brew?status.svg)](https://godoc.org/github.com/lodge93/cold-brew)

Cold Brew is a project to control a kyoto cold brew tower via software.

## Project Status
This was originally a successful hackweek project at my current place of
employment. I am currently in the process of refactoring code and gathering a
group of resources (documentation, release pipelines, wiring diagrams, etc) so
that others can build their own electronic cold brew towers.

## Frontend Development

``` bash
# install dependencies
npm install

# add your raspberry pi to /etc/hosts (this can also be changed in the webpack config)
sudo bash -c "echo '<raspberry pi IP address> cold-brew.dev' >> /etc/hosts"

# serve with hot reload at localhost:3000
npm run dev
```

## Backend Development

```bash
# ensure dependencies are properly vendored
dep ensure

# run server
go run main.go
```

## Deploy Development Build

```bash
# deploy using cold-brew.dev in /etc/hosts
./deploy/development.sh

# deploy skipping the frontend assets
./deploy/development.sh -s

# deploy setting custom remote host
./deploy/development.sh -h <raspberry pi IP address>
```
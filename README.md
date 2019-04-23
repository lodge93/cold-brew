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

This project has preliminary support for wrapping the existing UI into an
electron application.

```bash
# install electron globally
npm install -g electron

# start the development server
npm run dev

# launch the electron application
electron .
```

## Backend Development

```bash
# run server
go run main.go
```

The mocks for unit testing were generated using 
[mock](https://github.com/golang/mock):
```bash
# generate mocks
mockgen -destination dripper/mock_dripper/mock_dripper.go github.com/lodge93/cold-brew/dripper MotorController
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

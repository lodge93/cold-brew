# cold-brew
[![Build Status](https://cloud.drone.io/api/badges/betterengineering/cold-brew/status.svg)](https://cloud.drone.io/betterengineering/cold-brew)
[![Go Report Card](https://goreportcard.com/badge/github.com/betterengineering/cold-brew)](https://goreportcard.com/report/github.com/betterengineering/cold-brew)
[![GoDoc](https://godoc.org/github.com/betterengineering/cold-brew?status.svg)](https://godoc.org/github.com/betterengineering/cold-brew)

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

# serve with hot reload at localhost:3000
npm run dev
```

## Backend Development

```bash
# run server
make run
```

The mocks for unit testing were generated using 
[mock](https://github.com/golang/mock):
```bash
# generate mocks
make mocks
```

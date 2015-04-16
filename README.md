# go-uptime

[![Build Status](https://travis-ci.org/maxcnunes/go-uptime-api.svg?branch=master)](https://travis-ci.org/maxcnunes/go-uptime-api)

Simple monitor server to check uptime of any target reachable through HTTP.

The Go Uptime is composed of an API and [APP](https://github.com/maxcnunes/go-uptime-app) separated in different projects.

## API

This current project is responsible for manipulating the targets data and polling all targets' URL to check if each one is up or down. Also the monitor will listen to all Docker events and capture the URL from all containers that has the `VIRTUAL_HOST` environment variable.
Concerned in a better user experience the monitor uses web socket to notify the connected clients when a target has been created or updated.


# Developing

The simplest way is using Dockito vagrant box and docker-compose to provide a configured environment for you.

Setup the [Dockito vagrant box](https://github.com/dockito/devbox#dockito-vagrant-box) then inside the VM execute the command below to the docker-compose start the container:

```bash
docker-compose run local
```

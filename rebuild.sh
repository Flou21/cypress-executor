#!/bin/bash


# build own firefox image
docker build -t registry.flou21.de/firefox-cypress-browser:firefox86 -f firefox86/base/Dockerfile .

# push own firefox image
docker push registry.flou21.de/firefox-cypress-browser:firefox86 -f firefox-cypress-86

# build final test images
docker build -t registry.flou21.de/cypress-runner:edge88 -f edge88/Dockerfile .
docker build -t registry.flou21.de/cypress-runner:chrome91 -f chrome91/Dockerfile .
docker build -t registry.flou21.de/cypress-runner:chrome87 -f chrome87/Dockerfile .
docker build -t registry.flou21.de/cypress-runner:firefox86 -f firefox86/Dockerfile^.

# push all images
docker push registry.flou21.de/cypress-runner:edge88
docker push registry.flou21.de/cypress-runner:chrome91
docker push registry.flou21.de/cypress-runner:chrome87
docker push registry.flou21.de/cypress-runner:firefox86
# cypress executor

little go program that clones a repository and executes cypress tests in different browsers.


## usage

you can simply execute the rebuild script to build the images and push them to a registry.
please change the container image repository url

then start one with the following command e.g.

````
docker run registry.flou21.de/cypress-runner:chrome91 /app --browser chrome --repository https://github.com/Coflnet/hypixel-react --branch master
````


### mail

please change the mail stuff in main.go
you have to do it in the source code, because i don't want to do it otherwise now

### parameters

#### --browser

the browser, normally chrome, firefox, edge, etc.

#### --repository

the repository to clone

#### --branch

the branch^^
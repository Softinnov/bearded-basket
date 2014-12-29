bearded-basket
==============

![architecture](https://cloud.githubusercontent.com/assets/2647865/4824081/51a92f8a-5f56-11e4-8eb0-f0b978fba039.png)

PREREQUISITE
============

### Production

 - docker 1.2 (add modification on /etc/hosts)

### Developement

 - git
 - golang >= 1.3 (+godep)
 - docker >= 1.3 (add modification on /etc/hosts)

DEPLOYMENT
==========

The ssh keys have to be already installed in the server and the database initialised.

Build the docker image:
```bash
$ docker build -t softinnov/docker_dev .
```

Run the docker image:
```bash
$ docker run -it --rm -v `pwd`:/gopath/src/github.com/softinnov/bearded-basket --privileged -v [path].ssh:/root/.ssh softinnov/docker_dev
```

Once in the container:
```bash
$ service docker start
$ cd platforms/prod/deploy && go get ./... && go build && cd ..
$ ./deploy/deploy master --dir="scripts" --ip=<ip master>
$ ./deploy/deploy slave --dir="scripts" --ip=<ip slave> --master=<ip master>
```


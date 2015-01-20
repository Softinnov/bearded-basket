bearded-basket
==============

![architecture](https://docs.google.com/drawings/d/17gI1qThbS60gCfTtME9xIFRNDxauJp4hurALpPH_I8U/pub?w=678&h=488)

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


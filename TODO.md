
Add to README for docker

Build the docker image:
```bash
$ docker build -t softinnov/docker_dev .
```

Run the docker image
```bash
$ docker run -it --rm -v `pwd`:/gopath/src/github.com/softinnov/bearded-basket --privileged -v /root/.ssh:/root/.ssh softinnov/docker_dev
```

Once in the container:
```bash
$ service docker start
$ ./deploy/deploy master --dir="scripts" --ip=172.17.0.89
```

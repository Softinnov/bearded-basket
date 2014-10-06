
INSTALL
======

```sh
$ docker build -t softinnov/dev-server .
```

```sh
$ docker run --rm -v $(pwd) softinnov/dev-server go build -v
```

```sh
$ docker run -d --name dev-back -v $(pwd) softinnov/dev-server server
```

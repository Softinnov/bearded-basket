INSTALL
=======

Just build the container:
```sh
$ docker build softinnov/db_test .
```

Then go in parent folder and launch tests like this:
```sh
$ ./3.runserver.sh -t godep go test ./...
```

Extra:
If you want to debug the database, first lauch the db_test:
```sh
$ ./1.rundb.sh -t
```

Then inspect the database user the `docker exec` command:
```sh
$ docker exec -it db_test mysql prod_test
```

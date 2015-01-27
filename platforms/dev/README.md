INSTALL
======

You need Godep installed:
```bash
go get -u github.com/tools/godep
```

Before launching the script, put the data tables inside this directory (.sql & .txt).

Create images with:
```sh
$ ./buildall
```

Then launch every containers like this:
```sh
$ ./0.runconsul.sh
$ ./1.rundb.sh
$ ./2.runchey.sh <path to esc-pdv> $(pwd)/logs
$ ./3.runserver.sh $(pwd)/../server $(pwd)/logs
$ ./4.runclient.sh $(pwd)/../client $(pwd)/logs
```

Finally, open your browser at localhost:8000 (or boot2docker IP for MacOSX)

TESTS
=====

Create images with:
```sh
$ ./buildDB.sh -t
$ ./buildback.sh -t
```

Then launch the test container:
```sh
$ ./3.runserver.sh -t godep go test ./...
```
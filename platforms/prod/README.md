
DEPLOYEMENT
===========

### Debian/Ubuntu

You just need Go on you host and simply build the deploy project:
```bash
$ cd deploy && go build && cd ..
```

There is two configuration servers. The first is the `master` mysql server:
```bash
./deploy/deploy master --init --dir="scripts" --ip=<ip master> --key=<ssh rsa pub>
```

And the other for the `slave` server:
```bash
$ ./deploy/deploy slave --init --dir="scripts" --ip=<ip slave> --key=<ssh rsa pub> --master=<ip master>
```

#### Update

In order to just update the `master` and `slave` servers with the last builds, simply run the commands without the `--init` flag.

BUILD DEPLOYEMENT
=================

In order to build all images for the deployement, simply run `build.sh`.

DEPLOYEMENT
===========

### Debian/Ubuntu

First of all, add your ssh key to be identified without password check by running:

```sh
$ ./scripts/init.sh 0.0.0.0 <id_rsa.pub>
```

Then, run the following command (replace the IP with the proper one):
```sh
$ ./scripts/launch.sh 0.0.0.0 install.sh
```

This command will install all necessary packages for the server configuration.

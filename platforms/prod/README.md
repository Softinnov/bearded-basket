BUILD DEPLOYEMENT
=================

In order to build all images for the deployement, simply run `build.sh`.

DEPLOYEMENT
===========

### Debian/Ubuntu

First of all, add your ssh key to be identified without password check by running:
(Replace the ssh key file by yours)

```sh
$ ./init.sh 0.0.0.0 ~/.ssh/id_rsa.pub
```

Then, run the following command (replace the IP with the proper one):
```sh
$ ./run.sh 0.0.0.0 install.sh
```

This command will install all necessary packages for the server configuration.

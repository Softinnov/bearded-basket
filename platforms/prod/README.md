DEPLOYEMENT
===========

First of all, add you ssh key to be identified without password check by running:
(Replace you ssh key file by yours)

```sh
$ ./init.sh ~/.ssh/id_rsa.pub
```

Then, run the following command:
```sh
$ ./run.sh install.sh
```

This command will install all necessary packages for the server configuration.

INSTALL
=======

Run the container data only:
```bash
$ docker run -d -v /var/lib/mysql --name dbdata busybox echo data-only
```

Database initialization (ONLY if not initalized before):
```bash
$ docker run --rm --volumes-from dbdata -p 3306:3306 softinnov/db
```

Then, you need to stop the last container (db):
```bash
$ docker stop db
$ docker rm db
```

Create database `prod`:
```bash
$ docker run --rm --volumes-from dbdata -P softinnov/db bash -c "/create_db.sh prod"
```

Then, look for admin password in logs.

Populate database (having them inside `$(pwd)`):
```bash
$ docker run --rm -v $(pwd):/data --volumes-from dbdata softinnov/db /bin/bash -c "/import_sql.sh admin [password] prod pdv [+other tables]"
```

Finally launch the container mysql:
```bash
$ docker run -d --volumes-from dbdata --name db ghigt/db
```

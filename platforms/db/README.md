INSTALL
=======

Run the container data only:
```bash
$ docker run -d -v /var/lib/mysql --name dbdata busybox echo data-only
```

Database initialization (ONLY if not initalized before):
First store the username + password in your local env (`$DBUSER` & `$DBPASS`), then launch:
```bash
$ docker run --rm --volumes-from dbdata -e MYSQL_USER=$DBUSER -e MYSQL_PASS=$DBPASS softinnov/db
```

Create database `prod`:
```bash
$ docker run --rm --volumes-from dbdata softinnov/db bash -c "/create_db.sh prod"
```

Then, look for admin password in logs.

Populate database (having them inside `$(pwd)`):
```bash
$ docker run --rm -v $(pwd):/data --volumes-from dbdata softinnov/db /bin/bash -c "/import_sql.sh $DBUSER $DBPASS prod pdv [+other tables]"
```

Finally launch the container mysql:
```bash
$ docker run -d --volumes-from dbdata --name db softinnov/db
```
